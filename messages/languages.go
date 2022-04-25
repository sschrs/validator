package messages

import (
	"validator/messages/de"
	"validator/messages/en"
	"validator/messages/fr"
	"validator/messages/tr"
)

var Languages = map[string]map[string]string{
	"en": en.EN,
	"tr": tr.TR,
	"de": de.DE,
	"fr": fr.FR,
}
