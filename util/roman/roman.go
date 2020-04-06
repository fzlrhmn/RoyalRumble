package roman

import "fmt"

// mappingNumber between roman and number.
// Limited to 50 only due to the task.
var mappingNumber = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
}

// Roman is a struct for roman
type Roman struct{}

// NewRoman to create Roman instance
func NewRoman() *Roman {
	return &Roman{}
}

// ConvertToNumber to covert roman string to integer
func (r *Roman) ConvertToNumber(s string) (int, error) {

	// define output as 0
	output := 0

	// length of roman string
	lens := len(s)

	// iterate through the length of roman number
	for i := 0; i < lens; i++ {

		// get the string by i index
		c := string(s[i])

		// if roman character exist in the mappingNumber struct
		if vc, ok := mappingNumber[c]; ok {

			// check length of lens is enough for do calculation conversion
			if i < lens-1 {

				// get the next string value
				cn := string(s[i+1])

				// get the integer of given roman character
				csn := mappingNumber[cn]

				// if first value is less than next value
				if vc < csn {
					output += csn - vc
					i++
				} else {
					// if first value is more than next value
					output += vc
				}
			} else {
				output += vc
			}
		} else {
			// return error on unrecognized roman character
			return 0, fmt.Errorf("Roman not recognized")
		}
	}

	// return number of the roman character
	return output, nil
}
