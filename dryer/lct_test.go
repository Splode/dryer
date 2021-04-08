package dryer

import "testing"

type searchCase struct {
	src      []Token
	pat      []Token
	min      int
	expected int
}

var searchTests = []searchCase{
	{
		src:      []Token{{tokenString: "foo"}},
		pat:      []Token{{tokenString: "foo"}, {tokenString: "bar"}},
		min:      2,
		expected: 0,
	},
	{
		src:      []Token{{tokenString: "foo"}, {tokenString: "bar"}},
		pat:      []Token{{tokenString: "foo"}},
		min:      2,
		expected: 0,
	},
	{
		src:      []Token{{tokenString: "foo"}},
		pat:      []Token{{tokenString: "foo"}},
		min:      2,
		expected: 0,
	},
	{
		src:      []Token{},
		pat:      []Token{},
		min:      2,
		expected: 0,
	},
	{
		src:      []Token{{tokenString: "foo"}, {tokenString: "bar"}},
		pat:      []Token{{tokenString: "foo"}, {tokenString: "bar"}},
		min:      2,
		expected: 1,
	},
	{
		src:      []Token{{tokenString: "foo"}, {tokenString: "bar"}},
		pat:      []Token{{tokenString: "foo"}, {tokenString: "bar"}},
		min:      1,
		expected: 1,
	},
	{
		src:      []Token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "bang"}, {tokenString: "wiz"}},
		pat:      []Token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "wiz"}, {tokenString: "wiz"}},
		min:      2,
		expected: 1,
	},
	{
		src:      []Token{{tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}},
		pat:      []Token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "wiz"}},
		min:      2,
		expected: 3,
	},
	{
		src:      []Token{{tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}},
		pat:      []Token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "wiz"}},
		min:      3,
		expected: 0,
	},
	{
		src:      []Token{{tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "foo"}},
		pat:      []Token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "bar"}, {tokenString: "bar"}},
		min:      3,
		expected: 1,
	},
	{
		src:      []Token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "wiz"}, {tokenString: "bang"}},
		pat:      []Token{{tokenString: "brass"}, {tokenString: "gold"}, {tokenString: "silver"}, {tokenString: "lead"}},
		min:      3,
		expected: 0,
	},
	{
		src:      []Token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "wiz"}, {tokenString: "bang"}},
		pat:      []Token{{tokenString: "foo"}, {tokenString: "gold"}, {tokenString: "silver"}, {tokenString: "lead"}},
		min:      1,
		expected: 1,
	},
	{
		src: []Token{
			{tokenString: "foo"},
			{tokenString: "bar"},
			{tokenString: "wiz"},
			{tokenString: "bang"},
			{tokenString: "silver"},
			{tokenString: "lead"},
		},
		pat:      []Token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "silver"}, {tokenString: "lead"}},
		min:      2,
		expected: 2,
	},
	{
		src: []Token{
			{tokenString: "foo"},
			{tokenString: "bar"},
			{tokenString: "wiz"},
			{tokenString: "bang"},
			{tokenString: "silver"},
			{tokenString: "lead"},
		},
		pat:      []Token{{tokenString: "foo"}, {tokenString: "bar"}, {tokenString: "silver"}, {tokenString: "lead"}},
		min:      4,
		expected: 0,
	},
	{
		src: []Token{
			{tokenString: "foo"},
			{tokenString: "bar"},
			{tokenString: "wiz"},
			{tokenString: "bang"},
			{tokenString: "silver"},
			{tokenString: "lead"},
			{tokenString: "wiz"},
		},
		pat: []Token{
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
		src: []Token{
			{tokenString: "foo"},
			{tokenString: "bar"},
			{tokenString: "wiz"},
			{tokenString: "bang"},
			{tokenString: "silver"},
			{tokenString: "lead"},
			{tokenString: "wiz"},
		},
		pat: []Token{
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
		src: []Token{
			{tokenString: "Melville"},
			{tokenString: "began"},
			{tokenString: "writing"},
			{tokenString: "Moby-Dick"},
			{tokenString: "in"},
			{tokenString: "February"},
			{tokenString: "1850"},
			{tokenString: ","},
		},
		pat: []Token{
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
		src: []Token{
			{tokenString: "Melville"},
			{tokenString: "began"},
			{tokenString: "writing"},
			{tokenString: "Moby-Dick"},
			{tokenString: "in"},
			{tokenString: "February"},
			{tokenString: "1850"},
			{tokenString: ","},
		},
		pat: []Token{
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
		src: []Token{
			{tokenString: "import"},
			{tokenString: "{"},
			{tokenString: "hysterics"},
			{tokenString: "}"},
			{tokenString: "from"},
			{tokenString: "\"module\""},
		},
		pat: []Token{
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
		res := Search(tokenSliceToStringer(tc.src), tokenSliceToStringer(tc.pat), tc.min)
		rec := len(res)
		exp := tc.expected

		if rec != exp {
			t.Errorf("received: %v expected: %v", rec, exp)
		}
	}
}
