package sr

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "sr",
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
