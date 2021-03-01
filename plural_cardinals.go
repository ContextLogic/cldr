package cldr

import "fmt"

var pluralCardinals = map[string] []PluralRule {
	"1":  {PluralRuleOther},
	"2A": {PluralRuleOne, PluralRuleOther},
	"2B": {PluralRuleOne, PluralRuleOther},
	"2C": {PluralRuleOne, PluralRuleOther},
	"2D": {PluralRuleOne, PluralRuleOther},
	"2E": {PluralRuleOne, PluralRuleOther},
	"2F": {PluralRuleOne, PluralRuleOther},
	"2G": {PluralRuleOne, PluralRuleOther},
	"2H": {PluralRuleOne, PluralRuleOther},
	"3A": {PluralRuleZero, PluralRuleOne, PluralRuleOther},
	"3B": {PluralRuleOne, PluralRuleTwo, PluralRuleOther},
	"3C": {PluralRuleOne, PluralRuleFew, PluralRuleOther},
	"3D": {PluralRuleOne, PluralRuleFew, PluralRuleOther},
	"3E": {PluralRuleZero, PluralRuleOne, PluralRuleOther},
	"3F": {PluralRuleZero, PluralRuleOne, PluralRuleOther},
	"3G": {PluralRuleOne, PluralRuleFew, PluralRuleOther},
	"3H": {PluralRuleOne, PluralRuleFew, PluralRuleOther},
	"4A": {PluralRuleOne, PluralRuleTwo, PluralRuleMany, PluralRuleOther},
	"4B": {PluralRuleOne, PluralRuleFew, PluralRuleMany, PluralRuleOther},
	"4C": {PluralRuleOne, PluralRuleFew, PluralRuleMany, PluralRuleOther},
	"4D": {PluralRuleOne, PluralRuleTwo, PluralRuleFew, PluralRuleOther},
	"4E": {PluralRuleOne, PluralRuleFew, PluralRuleMany, PluralRuleOther},
	"4F": {PluralRuleOne, PluralRuleTwo, PluralRuleFew, PluralRuleOther},
	"4G": {PluralRuleOne, PluralRuleTwo, PluralRuleFew, PluralRuleOther},
	"4H": {PluralRuleOne, PluralRuleFew, PluralRuleMany, PluralRuleOther},
	"4I": {PluralRuleOne, PluralRuleFew, PluralRuleMany, PluralRuleOther},
	"5A": {PluralRuleOne, PluralRuleTwo, PluralRuleFew, PluralRuleMany, PluralRuleOther},
	"5B": {PluralRuleOne, PluralRuleTwo, PluralRuleFew, PluralRuleMany, PluralRuleOther},
	"6A": {PluralRuleZero, PluralRuleOne, PluralRuleTwo, PluralRuleFew, PluralRuleMany, PluralRuleOther},
	"6B": {PluralRuleZero, PluralRuleOne, PluralRuleTwo, PluralRuleFew, PluralRuleMany, PluralRuleOther},
	"6C": {PluralRuleZero, PluralRuleOne, PluralRuleTwo, PluralRuleFew, PluralRuleMany, PluralRuleOther},
}

// GetPluralCardinals returns the plural cardinals corresponding to a locale
// http://docs.translatehouse.org/projects/localization-guide/en/latest/l10n/pluralforms.html
func GetPluralCardinals(locale string) ([]PluralRule, error) {
	pluralRule, err := GetPluralRule(locale)
	if err != nil {
		return nil, err
	}
	cardinals, ok := pluralCardinals[pluralRule]
	if !ok {
		return nil, fmt.Errorf("plural form not found for rule %s", pluralRule)
	}
	return cardinals, nil
}
