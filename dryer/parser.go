package dryer

import (
	"fmt"
	"sort"
	"sync"

	"github.com/splode/dryer/pkg/strings"
)

func Compare(cfg *Config) (e error) {
	if cfg.Pattern != "" {
		if cfg.Recurse {
			walkPattern(cfg)
		} else {
			if err := matchPattern(cfg.Dir, cfg); err != nil {
				return err
			}
		}
	}
	pathMatrix := strings.UniqueMatrix(cfg.Paths...)

	var wg sync.WaitGroup
	wg.Add(len(pathMatrix))
	comparisons := make([]*comparison, 0)
	for _, matrix := range pathMatrix {
		m := matrix
		go func(wg *sync.WaitGroup) {
			comp, err := parse(m[0], m[1], cfg)
			if err != nil {
				e = err
			}
			comparisons = append(comparisons, comp)
			wg.Done()
		}(&wg)
	}
	wg.Wait()

	fmt.Println()
	for _, c := range comparisons {
		if c.count() > 0 {
			print(c.tableData)
			// fmt.Printf("Total clones: %d of %d and %d\n", c.count(), c.sources[0].count(), c.sources[1].count())
		}
	}

	return e
}

// comparison represents the results of a search between 2 sources.
type comparison struct {
	sources   []source             // sources used in the comparison
	clones    []clone              // clones represent a collection of matching tokens in a comparison.
	tableData [][]string           // tableData are a collection of clone data serialized for reporting.
	result    map[int][][]stringer // result is the collection of matched tokens
}

// count returns the total number of clones in the comparison.
func (c *comparison) count() int {
	return len(c.result)
}

type source struct {
	path         string
	absolutePath string
	tokens       []token
}

func (s *source) count() int {
	return len(s.tokens)
}

type clone struct {
	tokens    [][]token
	fractions []fraction
}

type fraction struct {
	numerator int
	divisor   int
}

func (f *fraction) string() string {
	return fmt.Sprintf("%.2f%%", (float32(f.numerator)/float32(f.divisor))*100)
}

// parse parses the results of a search for matching tokens between 2 sources, returning a comparison representing the
// results.
func parse(s, p string, cfg *Config) (*comparison, error) {
	srcFile, err := open(s)
	if err != nil {
		return nil, err
	}
	sp := srcFile.path
	if cfg.Absolute {
		sp = srcFile.absolutePath
	}
	srcTokens := tokenize(srcFile.src, sp)

	patFile, err := open(p)
	if err != nil {
		return nil, err
	}
	pp := patFile.path
	if cfg.Absolute {
		pp = patFile.absolutePath
	}
	patTokens := tokenize(patFile.src, pp)

	sources := []source{
		{path: srcFile.path, absolutePath: srcFile.absolutePath, tokens: srcTokens},
		{path: patFile.path, absolutePath: patFile.absolutePath, tokens: patTokens},
	}

	result := search(tokenSliceToStringer(srcTokens), tokenSliceToStringer(patTokens), cfg.TokenMin)

	keys := sortedKeys(result)
	clones := make([]clone, 0)
	tableData := make([][]string, 0)
	for _, k := range keys {
		t := getTokenClones(result[k])
		clone := clone{
			tokens: t,
			fractions: []fraction{
				{numerator: len(result[k][0]), divisor: sources[0].count()},
				{numerator: len(result[k][1]), divisor: sources[1].count()},
			},
		}
		clones = append(clones, clone)

		d := cloneTableData(clone)
		tableData = append(tableData, d...)
	}

	return &comparison{sources: sources, clones: clones, tableData: tableData, result: result}, nil
}

// TODO: this has more to do with printing/reporting; should be in a different file
func cloneTableData(c clone) [][]string {
	td := make([][]string, 0)
	srcB := c.tokens[0][0]
	srcE := c.tokens[0][1]
	patB := c.tokens[1][0]
	patE := c.tokens[1][1]
	td = append(td,
		[]string{srcB.filename, srcB.lineColumn(), srcE.lineColumn(), c.fractions[0].string()},
		[]string{patB.filename, patB.lineColumn(), patE.lineColumn(), c.fractions[1].string()},
		[]string{"\t"},
	)
	return td
}

func getTokenClones(strs [][]stringer) [][]token {
	srcBeg := strs[0][0].(token)
	srcEnd := strs[0][len(strs[0])-1].(token)

	patBeg := strs[1][0].(token)
	patEnd := strs[1][len(strs[0])-1].(token)

	clones := make([][]token, 0)
	clones = append(clones, []token{srcBeg, srcEnd}, []token{patBeg, patEnd})
	return clones
}

func sortedKeys(m map[int][][]stringer) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(a, b int) bool {
		return keys[a] > keys[b]
	})
	return keys
}
