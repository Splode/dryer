package dryer

import "sort"

type stringer interface {
	string() string
}

func search(src, patterns []stringer, min int) map[int][][]stringer {
	i := indices(src, patterns, min)
	matches := make(map[int][][]stringer, len(i))

	var count int
	keys := sortedIntMap(i)
	for _, k := range keys {
		v := i[k]
		var s []stringer
		if v[1][0]+1 >= len(src) {
			s = src[v[0][0]:]
		} else {
			s = src[v[0][0] : v[1][0]+1]
		}

		var p []stringer
		if v[1][1]+1 >= len(patterns) {
			p = patterns[v[0][1]:]
		} else {
			p = patterns[v[0][1] : v[1][1]+1]
		}

		matches[count] = append(matches[count], s, p)
		count++
	}
	return matches
}

func indices(src, patterns []stringer, min int) map[int][][]int {
	m, n := len(src), len(patterns)
	matches := make(map[int][][]int)

	if m < min || n < min || min < 1 {
		return matches
	}

	lcs := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		lcs[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for ii := 0; ii <= n; ii++ {
			if i == 0 || ii == 0 {
				lcs[i][ii] = 0
			} else if src[i-1].string() == patterns[ii-1].string() {
				lcs[i][ii] = lcs[i-1][ii-1] + 1
			} else {
				lcs[i][ii] = 0
			}
		}
	}

	ind := make(map[int][][]int)
	for i := len(lcs) - 1; i > 0; i-- {
		for ii := len(lcs[i]) - 1; ii > 0; ii-- {
			if lcs[i][ii] >= min {
				c := lcs[i][ii]
				if _, ok := ind[i-c]; !ok {
					ind[i-c] = append(ind[i-c], []int{i - c, ii - c})
					ind[i-c] = append(ind[i-c], []int{i - 1, ii - 1})
				}
			}
		}
	}

	return ind
}

func sortedIntMap(m map[int][][]int) []int {
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
