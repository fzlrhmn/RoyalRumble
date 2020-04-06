package royal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoyalMapping(t *testing.T) {
	testCases := map[string]struct {
		RomanSets  []string
		NumberSets []string
		IsError    bool
	}{
		"ShouldReturnSuccess": {
			RomanSets:  []string{"Foo I", "Foo II", "Bar I", "Bar II"},
			NumberSets: []string{"Foo 1", "Foo 2", "Bar 1", "Bar 2"},
		},
		"ShouldReturnError": {
			RomanSets: []string{"Foo M"},
			IsError:   true,
		},
		"ShouldReturnErrorNotCorrectFormat": {
			RomanSets: []string{"Foo"},
			IsError:   true,
		},
		"ShouldReturnErrorRomanNumberIsEmpty": {
			RomanSets: []string{"Foo "},
			IsError:   true,
		},
	}

	for v, test := range testCases {
		t.Run(v, func(t *testing.T) {
			result, err := MapRoyalWithRoman(test.RomanSets)

			if test.IsError {
				assert.Error(t, err)
			} else {
				s := make([]string, 0, len(result))
				for k := range result {
					s = append(s, k)
				}

				assert.ElementsMatch(t, test.NumberSets, s)
			}
		})
	}
}