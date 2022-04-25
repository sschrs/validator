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

### Change Default Language
You can use the validator.ChangeDefaultLanguage() function to change the default language. Available languages: English, German, Turkish, French, Italian, Spanish, Dutch, Portuguese
```
validator.ChangeDefaultLanguage("tr") //Change language to turkish
```
### Change Default Tag Name
If you want to change the tag name, you can use the validator.ChangeTagName() function. This function returns the current tag name. A string value containing '=' sign is not accepted and the tag name won't change.
```
newTagName := validator.ChangeTagName("validator")
```
### Change Default Special Chars
The default characters for containsSpecialChars validation are: [@#$%^&+!?.,;:] If you want to change these characters, you can use the validator.ChangeSpecialChars() function. The "=" sign is not accepted and no change is made if it is sent as a parameter. This function returns new special characters.
```
newSpecialChars := validator.ChangeSpecialChars("?=&!")
```

### Change Validation Message
You can use the validator.ChangeMessage() function to change the message body. This function takes two parameters, the first validation name and the second new message body, and returns an error if it cannot find a valid validation.
```
err := validator.ChangeMessage("notNull", "anan")
```
