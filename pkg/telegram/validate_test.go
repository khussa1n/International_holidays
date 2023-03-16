package telegram

import "testing"

func TestValidateDate(t *testing.T) {

	table := []struct {
		month    string
		day      string
		expected bool
	}{
		{
			month:    "22",
			day:      "31",
			expected: false,
		},
	}

	for _, testCase := range table {
		result := ValidateDate(testCase.month, testCase.day)

		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %t, got %t", testCase.expected, result)
		}
	}
}
