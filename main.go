package main

import (
	"./levenshtein"
	"fmt"
	"sort"
)

func main() {
	var values = []string{"Asheville", "Arizona"}
	var key = "Aclahoma"
	sort.Sort(levenshtein.ByLeveshtein{Value: values, Key: key})
	fmt.Print("Result: ")
	for i := 0; i < len(values); i++ {
		fmt.Print(values[i] + ",")
	}
}
