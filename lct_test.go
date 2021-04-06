package main

import "testing"

type searchCase struct {
	src      []token
	pat      []token
	min      int
	expected int
}

var searchTests = []searchCase{
	{
		src:      []token{{tokenString: "foo"}},
		pat:      []token{{tokenString: "foo"}, {tokenString: "bar"}},
		min:      2,
		expected: 0,
	},
	{
		src:      []token{{tokenString: "foo"}, {tokenString: "bar"}},
		pat:      []token{{tokenString: "foo"}},
		min:      2,
		expected: 0,
	},
	{
		src:      []token{{tokenString: "foo"}},
		pat:      []token{{tokenString: "foo"}},
		min:      2,
		expected: 0,
	},
	{
		src:      []token{},
		pat:      []token{},
		min:      2,
		expected: 0,
	},
	{
		src:      []token{{tokenString: "foo"}, {tokenString: "bar"}},
		pat:      []token{{tokenString: "foo"}, {tokenString: "bar"}},
		min:      2,
		expected: 1,
	},
	{
		src:      []token{{tokenString: "foo"}, {tokenString: "bar"}},
		pat:      []token{{tokenString: "foo"}, {tokenString: "bar"}},
		min:      1,
		expected: 1,
	},
	{
		src:      []token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "bang"}, {tokenString: "wiz"}},
		pat:      []token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "wiz"}, {tokenString: "wiz"}},
		min:      2,
		expected: 1,
	},
	{
		src:      []token{{tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}},
		pat:      []token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "wiz"}},
		min:      2,
		expected: 3,
	},
	{
		src:      []token{{tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}},
		pat:      []token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "wiz"}},
		min:      3,
		expected: 0,
	},
	{
		src:      []token{{tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "foo"}},
		pat:      []token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}},
		min:      3,
		expected: 1,
	},
	{
		src:      []token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "wiz"}, {tokenString: "bang"}},
		pat:      []token{{tokenString: "brass"}, {tokenString: "gold"}, {tokenString: "silver"}, {tokenString: "lead"}},
		min:      3,
		expected: 0,
	},
	{
		src:      []token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "wiz"}, {tokenString: "bang"}},
		pat:      []token{{tokenString: "foo"}, {tokenString: "gold"}, {tokenString: "silver"}, {tokenString: "lead"}},
		min:      1,
		expected: 1,
	},
	{
		src: []token{
			{tokenString: "foo"},
			{tokenString: "bar"},
			{tokenString: "wiz"},
			{tokenString: "bang"},
			{tokenString: "silver"},
			{tokenString: "lead"},
		},
		pat:      []token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "silver"}, {tokenString: "lead"}},
		min:      2,
		expected: 2,
	},
	{
		src: []token{
			{tokenString: "foo"},
			{tokenString: "bar"},
			{tokenString: "wiz"},
			{tokenString: "bang"},
			{tokenString: "silver"},
			{tokenString: "lead"},
		},
		pat:      []token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "silver"}, {tokenString: "lead"}},
		min:      4,
		expected: 0,
	},
	{
		src: []token{
			{tokenString: "foo"},
			{tokenString: "bar"},
			{tokenString: "wiz"},
			{tokenString: "bang"},
			{tokenString: "silver"},
			{tokenString: "lead"},
			{tokenString: "wiz"},
		},
		pat: []token{
			{tokenString: "bang"},
			{tokenString: "foo"},
			{tokenString: "wiz"},
			{tokenString: "bang"},
			{tokenString: "silver"},
			{tokenString: "lead"},
			{tokenString: "foo"},
		},
		min:      4,
		expected: 1,
	},
	{
		src: []token{
			{tokenString: "foo"},
			{tokenString: "bar"},
			{tokenString: "wiz"},
			{tokenString: "bang"},
			{tokenString: "silver"},
			{tokenString: "lead"},
			{tokenString: "wiz"},
		},
		pat: []token{
			{tokenString: "foo"},
			{tokenString: "bar"},
			{tokenString: "wiz"},
			{tokenString: "bang"},
			{tokenString: "silver"},
			{tokenString: "lead"},
			{tokenString: "foo"},
		},
		min:      4,
		expected: 1,
	},
	{
		src: []token{
			{tokenString: "Melville"},
			{tokenString: "began"},
			{tokenString: "writing"},
			{tokenString: "Moby-Dick"},
			{tokenString: "in"},
			{tokenString: "February"},
			{tokenString: "1850"},
			{tokenString: ","},
		},
		pat: []token{
			{tokenString: "and"},
			{tokenString: "finished"},
			{tokenString: "in"},
			{tokenString: "Moby-Dick"},
			{tokenString: "in"},
			{tokenString: "February"},
			{tokenString: "1850"},
			{tokenString: ","},
		},
		min:      4,
		expected: 1,
	},
	{
		src: []token{
			{tokenString: "Melville"},
			{tokenString: "began"},
			{tokenString: "writing"},
			{tokenString: "Moby-Dick"},
			{tokenString: "in"},
			{tokenString: "February"},
			{tokenString: "1850"},
			{tokenString: ","},
		},
		pat: []token{
			{tokenString: "Melville"},
			{tokenString: "began"},
			{tokenString: "writing"},
			{tokenString: "Moby-Dick"},
			{tokenString: "."},
			{tokenString: "February"},
			{tokenString: "1850"},
			{tokenString: ","},
		},
		min:      3,
		expected: 2,
	},
	{
		src: []token{
			{tokenString: "import"},
			{tokenString: "{"},
			{tokenString: "hysterics"},
			{tokenString: "}"},
			{tokenString: "from"},
			{tokenString: "\"module\""},
		},
		pat: []token{
			{tokenString: "import"},
			{tokenString: "{"},
			{tokenString: "hysterics"},
			{tokenString: "}"},
			{tokenString: "from"},
			{tokenString: "\"module\""},
		},
		min:      3,
		expected: 1,
	},
}

func TestSearch(t *testing.T) {
	for _, tc := range searchTests {
		res := search(tokenSliceToStringer(tc.src), tokenSliceToStringer(tc.pat), tc.min)
		rec := len(res)
		exp := tc.expected

		if rec != exp {
			t.Errorf("received: %v expected: %v", rec, exp)
		}
	}
}
