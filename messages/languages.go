package messages

import (
	"validator/messages/en"
	"validator/messages/tr"
)

var Languages = map[string]map[string]string{
	"en": en.EN,
	"tr": tr.TR,
}
