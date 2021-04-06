package main

type Stringer interface {
	String() string
}

func search(src, patterns []Stringer, min int) map[int][][]Stringer {
	i := indices(src, patterns, min)
	matches := make(map[int][][]Stringer, len(i))

	var count int
	for _, v := range i {
		var s []Stringer
		if v[1][0]+1 >= len(src) {
			s = src[v[0][0]:]
		} else {
			s = src[v[0][0] : v[1][0]+1]
		}

		var p []Stringer
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

func indices(src, patterns []Stringer, min int) map[int][][]int {
	m, n := len(src), len(patterns)
	matches := make(map[int][][]int)

	if m < min || n < min || min < 1 {
		return matches
	}

	lcs := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		lcs[i] = make([]int, n+1)
	}

	// index := make([][]int, 0)
	for i := 0; i <= m; i++ {
		for ii := 0; ii <= n; ii++ {
			if i == 0 || ii == 0 {
				lcs[i][ii] = 0
			} else if src[i-1].String() == patterns[ii-1].String() {
				c := lcs[i-1][ii-1] + 1
				lcs[i][ii] = c

				if c == min {
					// index = append(index, []int{i, ii})
				}

			} else {
				lcs[i][ii] = 0
			}
		}
	}

	test := make(map[int][][]int)
	for i := len(lcs) - 1; i > 0; i-- {
		for ii := len(lcs[i]) - 1; ii > 0; ii-- {
			if lcs[i][ii] >= min {
				c := lcs[i][ii]
				// fmt.Println(i-1, ii-1)
				// fmt.Println(i-c, ii-c)
				if _, ok := test[i-c]; !ok {
					test[i-c] = append(test[i-c], []int{i - c, ii - c})
					test[i-c] = append(test[i-c], []int{i - 1, ii - 1})
				}
			}
		}
	}

	// for i, s := range index {
	// 	matches[i] = append(matches[i], []int{s[0] - min, s[1] - min})
	// 	matches[i] = append(matches[i], s)

	// 	l0, l1 := len(lcs)-min, len(lcs[1])-min
	// 	max := lcs[s[0]][s[1]]
	// 	s0, s1 := s[0]+1, s[1]+1

	// 	if s0 > l0 || s1 > l1 {
	// 		continue
	// 	}

	// 	c := lcs[s0][s1]

	// 	for c > max {
	// 		matches[i] = pushPop(matches[i], []int{s0, s1})
	// 		max = c
	// 		s0++
	// 		s1++

	// 		if s0 > l0 || s1 > l1 {
	// 			break
	// 		}

	// 		c = lcs[s0+1][s1+1]
	// 	}
	// }

	return test
}

func pushPop(s [][]int, v []int) [][]int {
	if len(s) >= 2 {
		s = s[:len(s)-1]
	}
	s = append(s, v)
	return s
}
