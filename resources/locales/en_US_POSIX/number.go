package en_US_POSIX

import "github.com/theplant/i18n/cldr"

var (
	symbols = cldr.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: "0/00"}
	formats = cldr.NumberFormats{Decimal: "#0.######", Currency: "¤\u00a0#0.00", Percent: "#0%"}
)
