package fy

import "github.com/ContextLogic/cldr"

var (
	symbols = cldr.Symbols{Decimal: ",", Group: ".", Negative: "-", Percent: "%", PerMille: "‰"}
	formats = cldr.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##0.00;¤\u00a0#,##0.00-", Percent: "#,##0%"}
)
