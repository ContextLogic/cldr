package en_MW

import "github.com/ContextLogic/cldr"

var Locale = &cldr.Locale{
	Locale: "en_MW",
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
