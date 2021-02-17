package nnh

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "nnh",
	Number: cldr.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar:   calendar,
	PluralRule: pluralRule,
}

func init() {
	cldr.RegisterLocale(Locale)
}
