package validator

import (
	"strings"
	"testing"
)

type testStructNotNull struct {
	Field string `validate:"notNull"`
}

type testStructEmail struct {
	Field string `validate:"email"`
}

type testStructNumeric struct {
	Field string `validate:"numeric"`
}

type testStructContainsString struct {
	Field string `validate:"containsString"`
}

type testStructContainsNumber struct {
	Field string `validate:"containsNumber"`
}

type testStructContainsSpecialChars struct {
	Field string `validate:"containsSpecialChars"`
}

type testStructEqString struct {
	Field string `validate:"eq=10"`
}

type testStructGtString struct {
	Field string `validate:"gt=10"`
}

type testStructLtString struct {
	Field string `validate:"lt=10"`
}

type testStructGteString struct {
	Field string `validate:"gte=10"`
}

type testStructLteString struct {
	Field string `validate:"lte=10"`
}

type testStructNeString struct {
	Field string `validate:"ne=10"`
}

type testStructEqNumber struct {
	Field int `validate:"eq=10"`
}

type testStructGtNumber struct {
	Field int `validate:"gt=10"`
}

type testStructLtNumber struct {
	Field int `validate:"lt=10"`
}

type testStructGteNumber struct {
	Field int `validate:"gte=10"`
}

type testStructLteNumber struct {
	Field int `validate:"lte=10"`
}

type testStructNeNumber struct {
	Field int `validate:"ne=10"`
}

type testStructContains struct {
	Field string `validate:"contains=engineer"`
}

type testStructNotContains struct {
	Field string `validate:"notContains=war"`
}

type testStructBegin struct {
	Field string `validate:"begin=peace"`
}

type testStructNotBegin struct {
	Field string `validate:"notBegin=war"`
}

type testStructEnd struct {
	Field string `validate:"end=peace"`
}

type testStructNotEnd struct {
	Field string `validate:"notEnd=war"`
}

func TestNotNull(t *testing.T) {
	tests := []struct {
		testStruct testStructNotNull
		expected   bool
	}{
		{testStructNotNull{}, false},
		{testStructNotNull{""}, false},
		{testStructNotNull{"not null value"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for NotNull. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestEmail(t *testing.T) {
	tests := []struct {
		testStruct testStructEmail
		expected   bool
	}{
		{testStructEmail{}, false},
		{testStructEmail{"notemail"}, false},
		{testStructEmail{"example@example.com"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for Email. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestNumeric(t *testing.T) {
	tests := []struct {
		testStruct testStructNumeric
		expected   bool
	}{
		{testStructNumeric{"not numeric 123"}, false},
		{testStructNumeric{"10"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for Numeric. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestContainsString(t *testing.T) {
	tests := []struct {
		testStruct testStructContainsString
		expected   bool
	}{
		{testStructContainsString{"123123"}, false},
		{testStructContainsString{"123123asd"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for containsString. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestContainsNumber(t *testing.T) {
	tests := []struct {
		testStruct testStructContainsNumber
		expected   bool
	}{
		{testStructContainsNumber{"asdasd"}, false},
		{testStructContainsNumber{"asdasd123"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for containsNumber. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestContainsSpecialChars(t *testing.T) {
	tests := []struct {
		testStruct testStructContainsSpecialChars
		expected   bool
	}{
		{testStructContainsSpecialChars{"asd123"}, false},
		{testStructContainsSpecialChars{"asd123?"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for containsSpecialChars. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestEqString(t *testing.T) {
	tests := []struct {
		testStruct testStructEqString
		expected   bool
	}{
		{testStructEqString{"asdasdasd"}, false},
		{testStructEqString{"asdasdasda"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for eqString. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestEqNumber(t *testing.T) {
	tests := []struct {
		testStruct testStructEqNumber
		expected   bool
	}{
		{testStructEqNumber{120}, false},
		{testStructEqNumber{10}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for eqNumber. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestGtString(t *testing.T) {
	tests := []struct {
		testStruct testStructGtString
		expected   bool
	}{
		{testStructGtString{"asdasdasda"}, false},
		{testStructGtString{"asdasdasdaa"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for gtString. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestGtNumber(t *testing.T) {
	tests := []struct {
		testStruct testStructGtNumber
		expected   bool
	}{
		{testStructGtNumber{10}, false},
		{testStructGtNumber{11}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for gtNumber. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestLtString(t *testing.T) {
	tests := []struct {
		testStruct testStructLtString
		expected   bool
	}{
		{testStructLtString{"asdasdasda"}, false},
		{testStructLtString{"asdasdasd"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for ltString. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestLtNumber(t *testing.T) {
	tests := []struct {
		testStruct testStructLtNumber
		expected   bool
	}{
		{testStructLtNumber{10}, false},
		{testStructLtNumber{9}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for ltNumber. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestGteString(t *testing.T) {
	tests := []struct {
		testStruct testStructGteString
		expected   bool
	}{
		{testStructGteString{"asdasdasd"}, false},
		{testStructGteString{"asdasdasda"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for gteString. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestGteNumber(t *testing.T) {
	tests := []struct {
		testStruct testStructGteNumber
		expected   bool
	}{
		{testStructGteNumber{9}, false},
		{testStructGteNumber{10}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for gteNumber. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestLteString(t *testing.T) {
	tests := []struct {
		testStruct testStructLteString
		expected   bool
	}{
		{testStructLteString{"asdasdasdaa"}, false},
		{testStructLteString{"asdasdasda"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for lteString. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestLteNumber(t *testing.T) {
	tests := []struct {
		testStruct testStructLteNumber
		expected   bool
	}{
		{testStructLteNumber{11}, false},
		{testStructLteNumber{10}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for lteNumber. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestNeString(t *testing.T) {
	tests := []struct {
		testStruct testStructNeString
		expected   bool
	}{
		{testStructNeString{"asdasdasda"}, false},
		{testStructNeString{"a"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for neString. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestNeNumber(t *testing.T) {
	tests := []struct {
		testStruct testStructNeNumber
		expected   bool
	}{
		{testStructNeNumber{10}, false},
		{testStructNeNumber{1}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for neNumber. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		testStruct testStructContains
		expected   bool
	}{
		{testStructContains{"hello"}, false},
		{testStructContains{"helloengineerhello"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for contains. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestNotContains(t *testing.T) {
	tests := []struct {
		testStruct testStructNotContains
		expected   bool
	}{
		{testStructNotContains{"world war"}, false},
		{testStructNotContains{"world peace"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for notContains. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestBegin(t *testing.T) {
	tests := []struct {
		testStruct testStructBegin
		expected   bool
	}{
		{testStructBegin{"war at home,war in the world"}, false},
		{testStructBegin{"peace at home,peace in the world"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for begin. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestNotBegin(t *testing.T) {
	tests := []struct {
		testStruct testStructNotBegin
		expected   bool
	}{
		{testStructNotBegin{"war at home,war in the world"}, false},
		{testStructNotBegin{"peace at home,peace in the world"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for begin. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestEnd(t *testing.T) {
	tests := []struct {
		testStruct testStructEnd
		expected   bool
	}{
		{testStructEnd{"world war"}, false},
		{testStructEnd{"world peace"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for end. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

func TestNotEnd(t *testing.T) {
	tests := []struct {
		testStruct testStructNotEnd
		expected   bool
	}{
		{testStructNotEnd{"world war"}, false},
		{testStructNotEnd{"world peace"}, true},
	}
	for _, test := range tests {
		if ok, _ := Validate(&test.testStruct); ok != test.expected {
			t.Errorf("Test Failed for notEnd. Expected: %v Output: %v", test.expected, ok)
		}
	}
}

type testStructCustomValidation struct {
	Field string `validate:"containsAmp"`
}

func TestAddCustomValidation(t *testing.T) {
	AddCustomValidation("containsAmp", func(field Field, value string) (bool, string) {
		if ok := strings.Contains(field.Value.String(), "&"); !ok {
			return ok, "Field must contain '&'"
		}
		return true, ""
	})

	var struct1 testStructCustomValidation
	var struct2 testStructCustomValidation
	struct1.Field = "hello world"
	struct2.Field = "hello world&universe"

	test1, _ := Validate(&struct1)
	test2, _ := Validate(&struct2)
	if test1 || !test2 {
		t.Error("Test Failed: AddCustomValidation")
	}

}
