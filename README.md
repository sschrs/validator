# Go Validator Package
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
```
type User struct {
	Name    string `validate:"notNull,gt=1"` //not null and length must be greater than 1
	Email   string `validate:"email"`
	Phone   string `validate:"eq=11,numeric"` //numeric and length must be equal to 11
	Age     int    `validate:"gte=18"`        //age must be greater than or equal to 18
	Passord string `validate:"containsString,containsNumber,containsSpecialChars"`
}
```
Use Validate() function to validate it:
```
func main() {
	var user User
	user.Name = "John Doe"
	user.Email = "example"
	if isValid, message := validator.Validate(&user); !isValid {
		fmt.Println(message)
	}
}
```

