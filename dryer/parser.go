package dryer

import (
	"fmt"
	"sort"
	"sync"

	"github.com/splode/dryer/pkg/strings"
)

func Compare(cfg *Config) {
	pathMatrix := strings.UniqueMatrix(cfg.Paths...)

	var wg sync.WaitGroup
	wg.Add(len(pathMatrix))
	res := make([][][]string, 0)
	for _, matrix := range pathMatrix {
		m := matrix
		go func(wg *sync.WaitGroup) {
			clones, err := parse(m[0], m[1], cfg.TokenMin)
			if err != nil {
				fmt.Println(err)
			}
			res = append(res, clones)
			wg.Done()
		}(&wg)
	}
	wg.Wait()

	for _, r := range res {
		if len(r) > 1 {
			print(r)
		}
	}
}

func parse(s, p string, min int) ([][]string, error) {
	srcFile, err := open(s)
	if err != nil {
		return nil, err
	}
	srcTokens := tokenize(srcFile.src, srcFile.absolutePath)

	patFile, err := open(p)
	if err != nil {
		return nil, err
	}
	patTokens := tokenize(patFile.src, patFile.absolutePath)

	res := search(tokenSliceToStringer(srcTokens), tokenSliceToStringer(patTokens), min)

	keys := sortedKeys(res)
	cloneData := make([][]string, 0)
	for _, k := range keys {
		c := getTokenClones(res[k])
		d := cloneTableData(c)
		cloneData = append(cloneData, d...)
	}

	return cloneData, nil
}

func cloneTableData(clones [][]token) [][]string {
	td := make([][]string, len(clones))
	srcB := clones[0][0]
	srcE := clones[0][1]
	patB := clones[1][0]
	patE := clones[1][1]
	td = append(td,
		[]string{srcB.Filename, fmt.Sprintf("%d:%d", srcB.Line, srcB.Column), fmt.Sprintf("%d:%d", srcE.Line, srcE.Column)},
		[]string{patB.Filename, fmt.Sprintf("%d:%d", patB.Line, patB.Column), fmt.Sprintf("%d:%d", patE.Line, patE.Column)},
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
