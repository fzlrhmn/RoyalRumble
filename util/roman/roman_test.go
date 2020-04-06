package roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertRomanToNumber(t *testing.T) {
	TestCases := map[string]struct {
		Roman          string
		ExpectedResult int
		IsError        bool
	}{
		"I => 1": {
			Roman:          "I",
			ExpectedResult: 1,
		},
		"II => 2": {
			Roman:          "II",
			ExpectedResult: 2,
		},
		"III => 3": {
			Roman:          "III",
			ExpectedResult: 3,
		},
		"XLVII => 47": {
			Roman:          "XLVII",
			ExpectedResult: 47,
		},
		"MM => 2000, should not recognized since limited to 50": {
			Roman:          "MM",
			ExpectedResult: 0,
			IsError:        true,
		},
		"RomanNotRecognized": {
			Roman:   "ASD",
			IsError: true,
		},
	}

	for v, test := range TestCases {
		roman := NewRoman()
		t.Run(v, func(t *testing.T) {
			result, err := roman.ConvertToNumber(test.Roman)
			if test.IsError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.ExpectedResult, result)
				assert.NoError(t, err)
			}
		})
	}
}
