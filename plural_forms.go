package cldr

import "fmt"

var pluralForms = map[string] string {
	"1":  "nplurals=1; plural=0;",
	"2A": "nplurals=2; plural=(n != 1);",
	"2B": "nplurals=2; plural=(n > 1);",
	"2C": "nplurals=2; plural=(n > 1);",
	"2D": "nplurals=2; plural=(n%10 != 1 || n%100 == 11);",
	"2E": "nplurals=2; plural=(n > 2 && (11 > n || n > 99));",
	"2F": "nplurals=2; plural=(n != 1);",
	"2G": "nplurals=2; plural=(n > 1);",
	"2H": "nplurals=2; plural=(n != 1);",
	"2I": "nplurals=2; plural=(n%10 == 4 || n%10 == 6 || n%10 == 9);",
	"3A": "nplurals=3; plural=(n%10 == 0 || (11 <= n%100 && n%100 <= 19)) ? 0 : (n%10 == 1 && n%100 != 11) ? 1 : 2;",
	"3B": "nplurals=3; plural=(n == 1) ? 0 : (n == 2) ? 1 : 2;",
	"3C": "nplurals=3; plural=(n == 1) ? 0 : (n == 0 || (2 <= n%100 && n%100 <= 19)) ? 1 : 2;",
	"3D": "nplurals=3; plural=(n == 1) ? 0 : (n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20)) ? 1 : 2;",
	"3E": "nplurals=3; plural=(n == 0) ? 0 : (n==1) ? 1 : 2;",
	"3F": "nplurals=3; plural=(n == 0) ? 0 : (n==1) ? 1 : 2;",
	"3G": "nplurals=3; plural=(0 <= n && n <= 1) ? 0 : (2 <= n && n <= 10) ? 1 : 2;",
	"3H": "nplurals=3; plural=(n%10 == 1 && n%100 != 11) ? 0 : (n%10 >= 2 && n%10 <= 4 && (n%100 < 12 || n%100 > 14)) ? 1 : 2;",
	"4A": "nplurals=4; plural=(n == 1) ? 0 : (n == 2) ? 1 : (n != 0 && n != 10 && n%10 == 0) ? 2 : 3;",
	"4B": "nplurals=4; plural=(n%10 == 1 && n%100 != 11) ? 0 : (n%10 >= 2 && n%10 <= 4 && (n%100 < 12 || n%100 > 14)) ? 1 : (n%10 == 0 || (n%10 >= 5 && n%10 <= 9) || (n%100 >= 11 && n%100 <= 14)) ? 2 : 3;",
	"4C": "nplurals=4; plural=(n == 1) ? 0 : (n%10 >= 2 && n%10 <= 4 && (n%100 < 12 || n%100 > 14)) ? 1 : ((n%10 >= 0 && n%10 <= 1) || (n%10 >= 5 && n%10 <= 9) || (n%100 >= 12 && n%100 <= 14)) ? 2 : 3;",
	"4D": "nplurals=4; plural=(n%100 == 1) ? 0 : (n%100 == 2) ? 1 : (n%100 >= 3 && n%100 <= 4) ? 2 : 3;",
	"4E": "nplurals=4; plural=(n == 1) ? 0 : (n == 0 || (2 <= n%100 && n%100 <= 10)) ? 1 : (11 <= n%100 && n%100 <= 19) ? 2 : 3;",
	"4F": "nplurals=4; plural=(n == 1 || n == 11) ? 0 : (n == 2 || n == 12) ? 1 : ((3 <= n && n <= 10) || (13 <= n && n <= 19)) ? 2 : 3;",
	"4G": "nplurals=4; plural=(n == 1) ? 0 : (n == 2) ? 1 : (n % 20 == 0) ? 2 : 3;",
	"4H": "nplurals=3; plural=(n%10 == 1 && n%100 != 11) ? 0 : (n%10 >= 2 && n%10 <= 9 && (n%100 < 12 || n%100 > 19)) ? 1 : 2;",
	"4I": "nplurals=3; plural=(n == 1) ? 0 : (n >= 2 && n <= 4) ? 1 : 2;",
	"5A": "nplurals=5; plural=(n == 1) ? 0 : (n == 2) ? 1 : (3 <= n && n <= 6) ? 2 : (7 <= n && n <= 10) ? 3 : 4;",
	"5B": "nplurals=5; plural=(n%10 == 1 && (n%100 != 11 && n%100 != 71 && n%100 != 91)) ? 0 : (n%10 == 2 && (n%100 != 12 && n%100 != 72 && n%100 != 92)) ? 1 : ((n%10 == 3 || n%10 == 4 || n%10 == 9) && (n%100 < 10 || n%100 > 19) && (n%100 < 70 || n%100 > 79) && (n%100 < 90 || n%100 > 99)) ? 2 : (n != 0 && n%1000000 == 0) ? 3 : 4;",
	"6A": "nplurals=6; plural=(n==0) ? 0 : (n==1) ? 1 : (n==2) ? 2 : (n%100 >= 3 && n%100 <= 10) ? 3 : (n%100 >= 11) ? 4 : 5;",
	"6B": "nplurals=6; plural=(n==0) ? 0 : (n==1) ? 1 : (n==2) ? 2 : (n==3) ? 3 : (n==6) ? 4 : 5;",
	"6C": "nplurals=6; plural=(n==0) ? 0 : (n==1) ? 1 : (n%20 == 2 || (n%1000 == 0 && (1000 <= n%100000 && n%100000 <= 20000 || n%100000 == 40000 || n%100000 == 60000 || n%100000 == 80000)) || (n != 0 && n%1000000 == 100000)) ? 2 : (n%20 == 3) ? 3 : (n != 1 && n%20 == 1) ? 4 : 5;",
}

// GetPluralForm returns the plural form corresponding to a locale
// http://docs.translatehouse.org/projects/localization-guide/en/latest/l10n/pluralforms.html
func GetPluralForm(locale string) (string, error) {
	pluralRule, err := GetPluralRule(locale)
	if err != nil {
		return "", err
	}
	pluralForm, ok := pluralForms[pluralRule]
	if !ok {
		return "", fmt.Errorf("plural form not found for rule %s", pluralRule)
	}
	return pluralForm, nil
}
