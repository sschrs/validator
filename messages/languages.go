package messages

import (
	"github.com/sschrs/validator/messages/de"
	"github.com/sschrs/validator/messages/en"
	"github.com/sschrs/validator/messages/es"
	"github.com/sschrs/validator/messages/fr"
	"github.com/sschrs/validator/messages/it"
	"github.com/sschrs/validator/messages/nl"
	"github.com/sschrs/validator/messages/pt"
	"github.com/sschrs/validator/messages/tr"
)

var Languages = map[string]map[string]string{
	"en": en.EN,
	"tr": tr.TR,
	"de": de.DE,
	"fr": fr.FR,
	"es": es.ES,
	"nl": nl.NL,
	"it": it.IT,
	"pt": pt.PT,
}
