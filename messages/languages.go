package messages

import (
	"govalidator/messages/en"
	"govalidator/messages/tr"
)

var Languages = map[string]map[string]string{
	"en": en.EN,
	"tr": tr.TR,
}
