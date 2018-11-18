package main

import "fmt"

func main() {
	var str1 = []rune("Asheville")
	var str2 = []rune("Arizona")
	fmt.Println("Distance between Asheville and Arizona:",levenshtein(str1,str2))
}

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
		column[0] = x; lastkey := x - 1
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