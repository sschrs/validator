package validator

import (
	"errors"
	"testing"
)

func TestChangeDefaultLanguage(t *testing.T) {
	var tests = []struct {
		input, expected string
	}{
		//Unsupported Languages
		{"ch", "en"},
		{"fa", "en"},

		// Supported Languages
		{"tr", "tr"},
		{"de", "de"},
		{"en", "en"},
	}

	for _, test := range tests {
		if output := ChangeDefaultLanguage(test.input); output != test.expected {
			t.Errorf("Test Failed! Input:%s Expected:%s Output:%s", test.input, test.expected, output)
		}
	}
}

func TestChangeTagName(t *testing.T) {
	var tests = []struct {
		input, expected string
	}{
		{"=", "validate"},
		{"validator", "validator"},
		{"tagName", "tagName"},
		{"validate", "validate"},
	}

	for _, test := range tests {
		if output := ChangeTagName(test.input); output != test.expected {
			t.Errorf("Test Failed! Input:%s Expected:%s Output:%s", test.input, test.expected, output)
		}
	}
}

func TestChangeMessage(t *testing.T) {
	var tests = []struct {
		inputValidation, inputMessage string
		expected                      error
	}{
		{"notNull", "new message for notNull", nil},
		{"eqString", "new message for eq", nil},
		{"gteNumber", "new message fot gte", nil},
		{"contains", "new message for contains", nil},
		{"something_else", "new message", errors.New("unknown validation tag")},
	}

	for _, test := range tests {
		output := ChangeMessage(test.inputValidation, test.inputMessage)
		if (output == nil && output != test.expected) || (output != nil && errors.Is(output, test.expected)) {
			t.Errorf("Test Failed: InputValidation: %s Expected: %v Output: %v", test.inputValidation, test.expected, output)
		}
	}
}

func TestChangeSpecialChars(t *testing.T) {
	var tests = []struct {
		input, expected string
	}{
		{"@#$%^&+!?.,;:=", "@#$%^&+!?.,;:"}, // "=" not allowed
		{"@#$%^&+!?", "@#$%^&+!?"},
	}
	for _, test := range tests {
		if output := ChangeSpecialChars(test.input); output != test.expected {
			t.Errorf("Test Failed! Input:%s Expected:%s Output:%s", test.input, test.expected, output)
		}
	}
}
