package vai_Latn

import "github.com/ContextLogic/cldr"

var (
	symbols = cldr.Symbols{Decimal: ".", Group: ",", Negative: "", Percent: "", PerMille: ""}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", Percent: ""}
)
