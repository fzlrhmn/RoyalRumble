package royal

import (
	"fmt"
	"strings"

	"github.com/fzlrhmn/playground/RoyalRumble/util/roman"
)

// MapRoyalWithRoman is a function for mapping between
// royal with roman number to royal number with integer
func MapRoyalWithRoman(s []string) (map[string]string, error) {
	// initiate for royal mapping
	f := make(map[string]string)

	// create roman instance
	c := roman.NewRoman()

	// iterate through the array
	for _, v := range s {

		// split the name and roman
		ss := strings.Split(v, " ")

		// if length is not correct
		if len(ss) != 2 {
			return nil, fmt.Errorf("Royal is not in correct format")
		}

		// check if it has roman number
		if ss[1] == "" {
			return nil, fmt.Errorf("Roman number not found")
		}

		// convert roman into number
		num, err := c.ConvertToNumber(ss[1])
		if err != nil {
			return nil, err
		}

		// create key and put it into royal mapping
		k := fmt.Sprintf("%s %d", ss[0], num)
		f[k] = v
	}

	// return royal mapping
	return f, nil
}
