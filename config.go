package govalidator

import (
	"errors"
	"govalidator/messages"
)

var (
	defaultLanguage     = "en"
	defaultTagName      = "validate"
	defaultTagSeperator = ","
	specialChars        = "@#$%^&+=!?.,;:"
)

func ChangeDefaultLanguage(lang string) {
	defaultLanguage = lang
}

func ChangeTagName(tagName string) {
	defaultTagName = tagName
}

func ChangeMessage(validation, newMessage string) error {
	if messages.Languages[defaultLanguage][validation] != "" {
		messages.Languages[defaultLanguage][validation] = newMessage
		return nil
	} else {
		return errors.New("unknown validation tag")
	}
}

func ChangeSpecialChars(chars string) {
	specialChars = chars
}
