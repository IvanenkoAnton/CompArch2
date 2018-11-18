package levenshtein

type ByLeveshtein struct {
	Value []string
	Key   string
}

func (a ByLeveshtein) Len() int { return len(a.Value) }
func (a ByLeveshtein) Less(i, j int) bool {
	return levenshtein([]rune(a.Value[i]), []rune(a.Key)) < levenshtein([]rune(a.Value[j]), []rune(a.Key))
}
func (a ByLeveshtein) Swap(i, j int) { a.Value[i], a.Value[j] = a.Value[j], a.Value[i] }

func levenshtein(str1, str2 []rune) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, s1len+1)
	for y := 1; y <= s1len; y++ {
		column[y] = y
	}

	calccolumn(column, s1len, s2len, str1, str2)

	return column[s1len]
}

func calccolumn(column []int, s1len, s2len int, str1, str2 []rune) {
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}
			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
}

func minimum(inputs ...int) int {
	min := inputs[0]
	for _, value := range inputs {
		if value < min {
			min = value
		}
	}
	return min
}
