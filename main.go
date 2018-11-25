package main

import (
	"./levenshtein"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
)

func main() {

	if len(os.Args) > 0 {
		var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
		var memprofile = flag.String("memprofile", "", "write mem profile to file")
		flag.Parse()
		if *cpuprofile != "" {
			f, err := os.Create(*cpuprofile)
			if err != nil {
				log.Fatal(err)
			}
			pprof.StartCPUProfile(f)

		}

		testdrive()

		defer pprof.StopCPUProfile()

		if *memprofile != "" {
			f, err := os.Create(*memprofile)
			if err != nil {
				log.Fatal(err)
			}

			pprof.WriteHeapProfile(f)
			f.Close()
		}

		fmt.Println("Done!")
	} else {

		simple()
	}
}

func simple() {
	var words = []string{"Asheville", "Arizona"}
	var key = "Aclahoma"
	fmt.Println("Words: ", words)
	fmt.Println("Key: ", key)

	sort.Sort(levenshtein.ByLeveshtein{Value: words, Key: key})

	fmt.Println("Result: ", words)
}

func testdrive() {
	var words = getwords()
	var key = "Aclahoma"

	sort.Sort(levenshtein.ByLeveshtein{Value: words, Key: key})
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
