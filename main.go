package main

import (
	"./levenshtein"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {

	var args = os.Args[1:]

	if args[0] == "test-drive" {
		testdrive()
	} else {
		simple()
	}
}

func simple() {
	var values = []string{"Asheville", "Arizona"}
	var key = "Aclahoma"
	sort.Sort(levenshtein.ByLeveshtein{Value: values, Key: key})
	fmt.Println("Result: ", values)
}

func testdrive() {
	var words = getwords()
	var key = "Aclahoma"
	fmt.Println("Words: ", words)
	fmt.Println("Key: ", key)

	var t1 = time.Now()
	sort.Sort(levenshtein.ByLeveshtein{Value: words, Key: key})
	var t2 = time.Now()

	fmt.Println("Result: ", words)
	fmt.Println("Time (s): ", float64(t2.UnixNano()-t1.UnixNano())/100000000)
}

func getwords() []string {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return words
}
