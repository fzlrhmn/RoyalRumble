package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/fzlrhmn/playground/RoyalRumble/util/royal"
)

func main() {
	// read input file
	si, err := readInputFile()
	if err != nil {
		panic(err)
	}

	// create royal mapping between roman and number
	r, err := royal.MapRoyalWithRoman(si)
	if err != nil {
		panic(err)
	}

	// initiate empty string slice with specific length of r
	s := make([]string, 0, len(r))

	// put all r key to slice
	for k := range r {
		s = append(s, k)
	}

	// sort s based on name and number
	sort.Slice(s, func(i, j int) bool {
		sin := strings.Split(s[i], " ")
		sjn := strings.Split(s[j], " ")
		if sin[0] != sjn[0] {
			return s[i] < s[j]
		}
		ii, _ := strconv.Atoi(sin[1])
		jj, _ := strconv.Atoi(sjn[1])
		return ii < jj
	})

	// print by sorted slice
	for _, v := range s {
		fmt.Printf("%q\n", r[v])
	}
}

func readInputFile() ([]string, error) {
	// read input file of time schedule periods
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		// throw error if it is not exist
		return nil, fmt.Errorf("Cannot read input.txt")
	}

	return strings.Split(string(data), "\n"), nil
}
