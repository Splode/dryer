package strings

import (
	"reflect"
	"testing"
)

type testCase struct {
	given    []string
	expected [][]string
}

var cases = []testCase{
	{
		given:    []string{"a", "b", "c"},
		expected: [][]string{{"a", "b"}, {"a", "c"}, {"b", "c"}},
	},
	{
		given:    []string{"a", "b", "c", "d"},
		expected: [][]string{{"a", "b"}, {"a", "c"}, {"a", "d"}, {"b", "c"}, {"b", "d"}, {"c", "d"}},
	},
	{
		given: []string{"a", "b", "c", "d", "e"},
		expected: [][]string{
			{"a", "b"}, {"a", "c"}, {"a", "d"}, {"a", "e"}, {"b", "c"},
			{"b", "d"}, {"b", "e"}, {"c", "d"}, {"c", "e"}, {"d", "e"},
		},
	},
	{
		given: []string{"a", "b", "c", "d", "e", "f"},
		expected: [][]string{
			{"a", "b"}, {"a", "c"}, {"a", "d"}, {"a", "e"}, {"a", "f"},
			{"b", "c"}, {"b", "d"}, {"b", "e"}, {"b", "f"}, {"c", "d"},
			{"c", "e"}, {"c", "f"}, {"d", "e"}, {"d", "f"}, {"e", "f"},
		},
	},
}

func TestUniqueMatrix(t *testing.T) {
	for _, tc := range cases {
		rec := UniqueMatrix(tc.given...)
		exp := tc.expected

		if eq := reflect.DeepEqual(rec, exp); !eq {
			t.Errorf("expected: %v, received: %v", rec, exp)
		}
	}
}
