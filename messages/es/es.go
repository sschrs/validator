package es

var ES = map[string]string{
	"notNull":              "el campo no puede estar vacío o ser un valor nulo",
	"email":                "el campo debe ser un correo electrónico",
	"numeric":              "el campo debe ser un valor numérico",
	"containsString":       "el campo debe contener al menos un carácter",
	"containsNumber":       "el campo debe contener al menos un número",
	"containsSpecialChars": "el campo debe contener al menos uno de los caracteres especiales",
	"eqNumber":             "el campo debe ser igual a",
	"eqString":             "la longitud del campo debe ser igual a",
	"gtNumber":             "el campo debe ser mayor que",
	"gtString":             "la longitud del campo debe ser mayor que",
	"ltNumber":             "el campo debe ser menor que",
	"ltString":             "la longitud del campo debe ser menor que",
	"lteNumber":            "el campo debe ser igual o menor que",
	"lteString":            "la longitud del campo debe ser igual o menor que",
	"gteNumber":            "el campo debe ser igual o mayor que",
	"gteString":            "la longitud del campo debe ser igual o mayor que",
	"neNumber":             "el campo no debe ser igual a",
	"neString":             "la longitud del campo no debe ser igual a",
	"contains":             "el campo debe contener",
	"notContains":          "el campo no debe contener",
	"begin":                "el campo debe comenzar con",
	"notBegin":             "el campo no debe comenzar con",
	"end":                  "el campo debe terminar con",
	"notEnd":               "el campo no debe terminar con",
}
