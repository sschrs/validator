# Go Validator Package
[![Go Reference](https://pkg.go.dev/badge/github.com/sschrs/validator.svg)](https://pkg.go.dev/github.com/sschrs/validator)<br>
Validator is a ligth package that allows you to validate over Go Structures and Tags. You can start using directly by adding the relevant tags to the structures you have created.<br><br>
**With the Validator package you can:**<br>
- You can perform many validation processes through the structures you have created.
- In case of validation fail, you can receive error messages in eight languages. [see. supported languages]
- You can change the validation error messages as you wish.
- You can implement your own validation functions.

**For Whom?**<br>
If you want to perform basic validations in a simple way, feel free to use this package.

## Installation
**Use go get:**<br>
`go get github.com/sschrs/validator`<br><br>
**Then import:**<br>
`import “github.com/sschrs/validator”`

## Usage and Documentation
### Basic Usage
Declare a struct with validate tags:<br>
```go
type User struct {
	Name    string `validate:"notNull,gt=1"` //not null and length must be greater than 1
	Email   string `validate:"email"`
	Phone   string `validate:"eq=11,numeric"` //numeric and length must be equal to 11
	Age     int    `validate:"gte=18"`        //age must be greater than or equal to 18
	Passord string `validate:"containsString,containsNumber,containsSpecialChars"`
}
```
Use Validate() function to validate it:
```go
func main() {
	var user User
	user.Name = "John Doe"
	user.Email = "example"
	if isValid, message := validator.Validate(&user); !isValid {
		fmt.Println(message)
	}
}
```

### Change Default Language
You can use the validator.ChangeDefaultLanguage() function to change the default language. Available languages: English, German, Turkish, French, Italian, Spanish, Dutch, Portuguese
```go
validator.ChangeDefaultLanguage("tr") //Change language to turkish
```
### Change Default Tag Name
If you want to change the tag name, you can use the validator.ChangeTagName() function. This function returns the current tag name. A string value containing '=' sign is not accepted and the tag name won't change.
```go
newTagName := validator.ChangeTagName("validator")
```
### Change Default Special Chars
The default characters for containsSpecialChars validation are: [@#$%^&+!?.,;:] If you want to change these characters, you can use the validator.ChangeSpecialChars() function. The "=" sign is not accepted and no change is made if it is sent as a parameter. This function returns new special characters.
```go
newSpecialChars := validator.ChangeSpecialChars("?=&!")
```

### Change Validation Message
You can use the validator.ChangeMessage() function to change the message body. This function takes two parameters, the first validation name and the second new message body, and returns an error if it cannot find a valid validation.
```go
err := validator.ChangeMessage("notNull", "new message for notNull validation")
```

### Add Custom Validation Func
The AddCustomValidation() function can be used to add a custom validation function. This function takes two parameters, the first validation name and the second validation function.<br>
First, let's examine the structure of the validation function. This function takes a structure of type validator.Field and a value of type string as a parameter.
```go
func validationFunc(field validator.Field,value string) (bool,string){}
```
The field structure is as follows. The name value is the name of the fields in the structure. Value is the value of this field, and tag is the tag corresponding to this field.
```go
type Field struct {
	Name  string
	Value reflect.Value
	Tag   string
}
```
The second parameter of the validation function represents the value of the validation rule separated by '=', if any.
#### Sample 1
Custom validation function that checks if the value of the structure's field contains an '&' sign.
```go
validator.AddCustomValidation("containsAmp", func(field validator.Field, value string) (bool, string) {
	if !strings.Contains(field.Value.String(), "&") {
		return false, fmt.Sprintf("%s field must containt '&'", field.Name)
	}
	return true, ""
})

type S struct {
	Text string `validate:"containsAmp"` // U can use with validation name in the tags
}
```

#### Sample 2
The custom validation function, which checks whether the number is divisible by the given value, is as follows.
```go
validator.AddCustomValidation("divisible", func(field validator.Field, _value string) (bool, string) {
	value, err := strconv.Atoi(_value)
	if err != nil {
		panic("the value can not convert to int")
	}
	if field.Value.Int()%int64(value) != 0 {
		return false, fmt.Sprintf("%s field must be divisible by %d", field.Name, value)
	}
	return true, ""
})

type S struct {
	Number int `validate:"divisible=3"`
}
```
### Validation Tags
|Tag|Validation|
|---|---|
|notNull|not null|
|email|email|
|numeric|numeric,just numbers|
|containsString|must contain at least one character|
|containsNumber|must contain at least one number|
|containsSpecialChars|must contain at least one special character(Default:[@#$%^&+!?.,;:])|
|eq|equal|
|gt|greater than|
|lt|less than|
|gte|greater than or equal|
|lte|less than or equal|
|ne|not equal|
|contains|must contain the given value|
|notContains|must not contain the given value|
|begin|must begin with given value|
|notBegin|must not begin with given value|
|end|must end with given value|
|notEnd|must not end with given value|

eq,ne,gt,gte,lt,lte -> Length for string expressions and values for numeric expressions are compared.

### Supported Languages
|Language|Code|
|--------|----|
|English|en|
|German|de|
|Turkish|tr|
|French|fr|
|Spanish|es|
|Dutch|nl|
|Portuguese|pr|
