package levenshtein

import "testing"

type testpair1 struct {
	values []int
	expect int
}

var tests1 = []testpair1{
	{[]int{1, 2}, 1},
	{[]int{2, 3, 4, 1, 3, 4}, 1},
	{[]int{-1, 1}, -1},
}

func TestMinimum(t *testing.T) {
	for _, pair := range tests1 {
		v := minimum(pair.values...)
		if v != pair.expect {
			t.Error(
				"For", pair.values,
				"expected", pair.expect,
				"got", v,
			)
		}
	}
}

type testpair2 struct {
	values [2]string
	expect int
}

var tests2 = []testpair2{
	{[2]string{"Asheville", "Arizona"}, 8},
	{[2]string{"Python", "Peithen"}, 3},
	{[2]string{"Orange", "Apple"}, 5},
}

func TestLevenshtein(t *testing.T) {
	for _, pair := range tests2 {
		v := levenshtein([]rune(pair.values[0]), []rune(pair.values[1]))
		if v != pair.expect {
			t.Error(
				"For", pair.values,
				"expected", pair.expect,
				"got", v,
			)
		}
	}
}
