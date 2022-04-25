package validator

import (
	"errors"
	"strings"
	"validator/messages"
)

var (
	defaultLanguage     = "en"
	supportedLanguages  = []string{"en", "de", "tr", "es", "fr", "nl", "it", "pt"}
	defaultTagName      = "validate"
	defaultTagSeperator = ","
	specialChars        = "@#$%^&+!?.,;:"
)

func ChangeDefaultLanguage(lang string) string {
	for _, supportedLang := range supportedLanguages {
		if supportedLang == lang {
			defaultLanguage = lang
			return defaultLanguage
		}
	}
	return defaultLanguage
}

func ChangeTagName(tagName string) string {
	if tagName == "=" {
		return defaultTagName
	}
	defaultTagName = tagName
	return defaultTagName
}

func ChangeMessage(validation, newMessage string) error {
	if messages.Languages[defaultLanguage][validation] != "" {
		messages.Languages[defaultLanguage][validation] = newMessage
		return nil
	} else {
		return errors.New("unknown validation tag")
	}
}

func ChangeSpecialChars(chars string) string {
	if !strings.Contains(chars, "=") {
		specialChars = chars
	}
	return specialChars
}
