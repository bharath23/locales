package fr

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type fr struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'fr' locale
func New() locales.Translator {
	return &fr{
		locale:                 "fr",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{6, 2},
		decimal:                "٫",
		group:                  "٬",
		minus:                  "‏−",
		percent:                "٪",
		perMille:               "؉",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED", "AFA", "AFN", "ALK ", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "$AR", "ATS ", "$AU", "AWG", "AZM", "AZN", "BAD ", "BAM", "BAN ", "BBD", "BDT", "BEC ", "FB", "BEL ", "BGL ", "BGM ", "BGN", "BGO ", "BHD", "BIF", "$BM", "$BN", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "$BS", "BTN", "BUK", "BWP", "BYB ", "BYR", "$BZ", "$CA", "CDF", "CHE ", "CHF", "CHW ", "CLE", "CLF", "$CL", "CNX", "CNY", "$CO", "COU", "CRC", "CSD ", "CSK ", "CUC", "CUP", "CVE", "£CY", "CZK", "DDM ", "DEM ", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK ", "EGP", "ERN", "ESA ", "ESB ", "ESP ", "ETB", "€", "FIM ", "$FJ", "£FK", "F", "£GB", "GEK", "GEL", "GHC ", "GHS", "£GI", "GMD", "GNF", "GNS ", "GQE", "GRD ", "GTQ", "GWE ", "GWP ", "GYD", "HKD", "HNL", "HRD ", "HRK", "HTG", "HUF", "IDR", "£IE", "£IL", "ILR", "₪", "₹", "IQD", "IRR", "ISJ ", "ISK", "₤IT", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "₩", "KWD", "KYD", "KZT", "LAK", "£LB", "LKR", "LRD", "lLS", "LTL", "LTT ", "LUC ", "LUF ", "LUL ", "LVL", "LVR ", "LYD", "MAD", "fMA", "MCF ", "MDC ", "MDL", "MGA", "Fmg", "MKD", "MKN ", "MLF ", "MMK", "MNT", "MOP", "MRO", "MTL ", "£MT", "MUR", "MVP", "MVR", "MWK", "$MX", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "$NA", "NGN", "NIC", "NIO", "NLG ", "NOK", "NPR", "$NZ", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ ", "PTE ", "PYG", "QAR", "$RH", "ROL ", "RON", "RSD", "RUB", "RUR ", "RWF", "SAR", "$SB", "SCR", "SDD", "SDG", "SDP", "SEK", "$SG", "SHP", "SIT ", "SKK ", "SLL", "SOS", "$SR", "SRG", "SSP", "STD", "SUR ", "SVC", "SYP", "SZL", "THB", "TJR ", "TJS", "TMM ", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "$TT", "TWD", "TZS", "UAH", "UAK ", "UGS", "UGX", "$US", "USN", "USS", "UYI", "UYP", "$UY", "UZS", "VEB", "VEF", "₫", "VNN", "VUV", "WS$", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "DTS", "XEU ", "XFO", "XFU", "CFA", "XPD", "FCFP", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
		daysAbbreviated:        []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."},
		daysNarrow:             []string{"D", "L", "M", "M", "J", "V", "S"},
		daysShort:              []string{"di", "lu", "ma", "me", "je", "ve", "sa"},
		daysWide:               []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"AM", "PM"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"av. J.-C.", "ap. J.-C."},
		erasNarrow:             []string{"av. J.-C.", "ap. J.-C."},
		erasWide:               []string{"avant Jésus-Christ", "après Jésus-Christ"},
		timezones:              map[string]string{"BOT": "heure de Bolivie", "UYST": "heure d’été de l’Uruguay", "GYT": "heure du Guyana", "EAT": "heure normale d’Afrique de l’Est", "NZST": "heure normale de la Nouvelle-Zélande", "WITA": "heure du Centre indonésien", "WEZ": "heure normale d’Europe de l’Ouest", "ACWST": "heure normale du centre-ouest de l’Australie", "CST": "heure normale du centre nord-américain", "EST": "heure normale de l’Est nord-américain", "BT": "heure du Bhoutan", "PDT": "heure d’été du Pacifique", "HKT": "heure normale de Hong Kong", "CLST": "heure d’été du Chili", "CAT": "heure normale d’Afrique centrale", "LHST": "heure normale de Lord Howe", "MST": "heure normale des Rocheuses", "NZDT": "heure d’été de la Nouvelle-Zélande", "CHAST": "heure normale des îles Chatham", "IST": "heure de l’Inde", "WESZ": "heure d’été d’Europe de l’Ouest", "MDT": "heure d’été des Rocheuses", "ACDT": "heure d’été du centre de l’Australie", "WAST": "heure d’été d’Afrique de l’Ouest", "∅∅∅": "heure d’été du Pérou", "AKST": "heure normale de l’Alaska", "JST": "heure normale du Japon", "ACST": "heure normale du centre de l’Australie", "AWST": "heure normale de l’Ouest de l’Australie", "AWDT": "heure d’été de l’Ouest de l’Australie", "ARST": "heure d’été de l’Argentine", "MESZ": "heure d’été d’Europe centrale", "UYT": "heure normale de l’Uruguay", "COST": "heure d’été de Colombie", "AST": "heure normale de l’Atlantique", "WIT": "heure de l’Est indonésien", "AEDT": "heure d’été de l’Est de l’Australie", "JDT": "heure d’été du Japon", "EDT": "heure d’été de l’Est", "COT": "heure normale de Colombie", "ADT": "heure d’été de l’Atlantique", "OESZ": "heure d’été d’Europe de l’Est", "GMT": "heure moyenne de Greenwich", "ACWDT": "heure d’été du centre-ouest de l’Australie", "SGT": "heure de Singapour", "CLT": "heure normale du Chili", "TMT": "heure normale du Turkménistan", "WIB": "heure de l’Ouest indonésien", "SRT": "heure du Suriname", "ChST": "heure des Chamorro", "HAT": "heure d’été de Terre-Neuve", "CDT": "heure d’été du Centre", "HKST": "heure d’été de Hong Kong", "HADT": "heure d’été d’Hawaii - Aléoutiennes", "CHADT": "heure d’été des îles Chatham", "HAST": "heure normale d’Hawaii - Aléoutiennes", "AEST": "heure normale de l’Est de l’Australie", "SAST": "heure normale d’Afrique méridionale", "ECT": "heure de l’Équateur", "LHDT": "heure d’été de Lord Howe", "MYT": "heure de la Malaisie", "PST": "heure normale du Pacifique nord-américain", "WAT": "heure normale d’Afrique de l’Ouest", "ART": "heure normale d’Argentine", "WART": "heure normale de l’Ouest argentin", "VET": "heure du Venezuela", "MEZ": "heure normale d’Europe centrale", "TMST": "heure d’été du Turkménistan", "AKDT": "heure d’été de l’Alaska", "OEZ": "heure normale d’Europe de l’Est", "WARST": "heure d’été de l’Ouest argentin", "HNT": "heure normale de Terre-Neuve", "GFT": "heure de la Guyane française"},
	}
}

// Locale returns the current translators string locale
func (fr *fr) Locale() string {
	return fr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fr'
func (fr *fr) PluralsCardinal() []locales.PluralRule {
	return fr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fr'
func (fr *fr) PluralsOrdinal() []locales.PluralRule {
	return fr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'fr'
func (fr *fr) PluralsRange() []locales.PluralRule {
	return fr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fr'
func (fr *fr) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fr'
func (fr *fr) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fr'
func (fr *fr) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := fr.CardinalPluralRule(num1, v1)
	end := fr.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fr *fr) MonthAbbreviated(month time.Month) string {
	return fr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fr *fr) MonthsAbbreviated() []string {
	return fr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fr *fr) MonthNarrow(month time.Month) string {
	return fr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fr *fr) MonthsNarrow() []string {
	return fr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fr *fr) MonthWide(month time.Month) string {
	return fr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fr *fr) MonthsWide() []string {
	return fr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fr *fr) WeekdayAbbreviated(weekday time.Weekday) string {
	return fr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fr *fr) WeekdaysAbbreviated() []string {
	return fr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fr *fr) WeekdayNarrow(weekday time.Weekday) string {
	return fr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fr *fr) WeekdaysNarrow() []string {
	return fr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fr *fr) WeekdayShort(weekday time.Weekday) string {
	return fr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fr *fr) WeekdaysShort() []string {
	return fr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fr *fr) WeekdayWide(weekday time.Weekday) string {
	return fr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fr *fr) WeekdaysWide() []string {
	return fr.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fr' and handles both Whole and Real numbers based on 'v'
func (fr *fr) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(fr.decimal) + len(fr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fr.decimal) - 1; j >= 0; j-- {
				b = append(b, fr.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fr.group) - 1; j >= 0; j-- {
					b = append(b, fr.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fr.minus) - 1; j >= 0; j-- {
			b = append(b, fr.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fr' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fr *fr) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(fr.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fr.decimal) - 1; j >= 0; j-- {
				b = append(b, fr.decimal[j])
			}
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fr.minus) - 1; j >= 0; j-- {
			b = append(b, fr.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fr.percentSuffix...)

	b = append(b, fr.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fr'
func (fr *fr) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fr.currencies[currency]
	l := len(s) + len(fr.decimal) + len(fr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fr.decimal) - 1; j >= 0; j-- {
				b = append(b, fr.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fr.group) - 1; j >= 0; j-- {
					b = append(b, fr.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fr.minus) - 1; j >= 0; j-- {
			b = append(b, fr.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, fr.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fr'
// in accounting notation.
func (fr *fr) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fr.currencies[currency]
	l := len(s) + len(fr.decimal) + len(fr.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(fr.decimal) - 1; j >= 0; j-- {
				b = append(b, fr.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fr.group) - 1; j >= 0; j-- {
					b = append(b, fr.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, fr.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, fr.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, fr.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'fr'
func (fr *fr) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'fr'
func (fr *fr) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'fr'
func (fr *fr) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'fr'
func (fr *fr) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, fr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, fr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'fr'
func (fr *fr) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'fr'
func (fr *fr) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'fr'
func (fr *fr) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'fr'
func (fr *fr) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, fr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := fr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
