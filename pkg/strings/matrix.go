package strings

// UniqueMatrix returns a 2-dimensional matrix of unique string combinations given any number of strings.
func UniqueMatrix(strs ...string) [][]string {
	m := make([][]string, 0)
	l := len(strs)

	if l == 1 {
		return m
	}

	for i := 0; i < l-1; i++ {
		for ii := i + 1; ii < l; ii += 1 {
			m = append(m, []string{strs[i], strs[ii]})
		}
	}

	return m
}
