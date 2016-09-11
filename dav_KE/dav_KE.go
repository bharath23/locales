package dav_KE

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type dav_KE struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
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

// New returns a new instance of translator for the 'dav_KE' locale
func New() locales.Translator {
	return &dav_KE{
		locale:                 "dav_KE",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		timeSeparator:          ":",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Imb", "Kaw", "Kad", "Kan", "Kas", "Kar", "Mfu", "Wun", "Ike", "Iku", "Imw", "Iwi"},
		monthsNarrow:           []string{"", "I", "K", "K", "K", "K", "K", "M", "W", "I", "I", "I", "I"},
		monthsWide:             []string{"", "Mori ghwa imbiri", "Mori ghwa kawi", "Mori ghwa kadadu", "Mori ghwa kana", "Mori ghwa kasanu", "Mori ghwa karandadu", "Mori ghwa mfungade", "Mori ghwa wunyanya", "Mori ghwa ikenda", "Mori ghwa ikumi", "Mori ghwa ikumi na imweri", "Mori ghwa ikumi na iwi"},
		daysAbbreviated:        []string{"Jum", "Jim", "Kaw", "Kad", "Kan", "Kas", "Ngu"},
		daysNarrow:             []string{"J", "J", "K", "K", "K", "K", "N"},
		daysWide:               []string{"Ituku ja jumwa", "Kuramuka jimweri", "Kuramuka kawi", "Kuramuka kadadu", "Kuramuka kana", "Kuramuka kasanu", "Kifula nguwo"},
		periodsAbbreviated:     []string{"Luma lwa K", "luma lwa p"},
		periodsWide:            []string{"Luma lwa K", "luma lwa p"},
		erasAbbreviated:        []string{"KK", "BK"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Kabla ya Kristo", "Baada ya Kristo"},
		timezones:              map[string]string{"WAST": "WAST", "CAT": "CAT", "GFT": "GFT", "AKDT": "AKDT", "WESZ": "WESZ", "CLT": "CLT", "ECT": "ECT", "EST": "EST", "WARST": "WARST", "SRT": "SRT", "ART": "ART", "WIT": "WIT", "WITA": "WITA", "OESZ": "OESZ", "IST": "IST", "PDT": "PDT", "MDT": "MDT", "AEDT": "AEDT", "SAST": "SAST", "OEZ": "OEZ", "BOT": "BOT", "PST": "PST", "COST": "COST", "NZST": "NZST", "VET": "VET", "ChST": "ChST", "HNT": "HNT", "TMT": "TMT", "WIB": "WIB", "UYST": "UYST", "ACWST": "ACWST", "ACWDT": "ACWDT", "CDT": "CDT", "TMST": "TMST", "ACST": "ACST", "GYT": "GYT", "ADT": "ADT", "SGT": "SGT", "BT": "BT", "LHST": "LHST", "LHDT": "LHDT", "JDT": "JDT", "ACDT": "ACDT", "AST": "AST", "MESZ": "MESZ", "CHAST": "CHAST", "AKST": "AKST", "HKST": "HKST", "HADT": "HADT", "WART": "WART", "MST": "MST", "ARST": "ARST", "AWST": "AWST", "HKT": "HKT", "∅∅∅": "∅∅∅", "AEST": "AEST", "JST": "JST", "UYT": "UYT", "NZDT": "NZDT", "MEZ": "MEZ", "CHADT": "CHADT", "GMT": "GMT", "EAT": "EAT", "COT": "COT", "CST": "CST", "WAT": "WAT", "AWDT": "AWDT", "CLST": "CLST", "MYT": "MYT", "WEZ": "WEZ", "HAT": "HAT", "HAST": "HAST", "EDT": "EDT"},
	}
}

// Locale returns the current translators string locale
func (dav *dav_KE) Locale() string {
	return dav.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'dav_KE'
func (dav *dav_KE) PluralsCardinal() []locales.PluralRule {
	return dav.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'dav_KE'
func (dav *dav_KE) PluralsOrdinal() []locales.PluralRule {
	return dav.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'dav_KE'
func (dav *dav_KE) PluralsRange() []locales.PluralRule {
	return dav.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'dav_KE'
func (dav *dav_KE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'dav_KE'
func (dav *dav_KE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'dav_KE'
func (dav *dav_KE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (dav *dav_KE) MonthAbbreviated(month time.Month) string {
	return dav.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (dav *dav_KE) MonthsAbbreviated() []string {
	return dav.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (dav *dav_KE) MonthNarrow(month time.Month) string {
	return dav.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (dav *dav_KE) MonthsNarrow() []string {
	return dav.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (dav *dav_KE) MonthWide(month time.Month) string {
	return dav.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (dav *dav_KE) MonthsWide() []string {
	return dav.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (dav *dav_KE) WeekdayAbbreviated(weekday time.Weekday) string {
	return dav.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (dav *dav_KE) WeekdaysAbbreviated() []string {
	return dav.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (dav *dav_KE) WeekdayNarrow(weekday time.Weekday) string {
	return dav.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (dav *dav_KE) WeekdaysNarrow() []string {
	return dav.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (dav *dav_KE) WeekdayShort(weekday time.Weekday) string {
	return dav.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (dav *dav_KE) WeekdaysShort() []string {
	return dav.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (dav *dav_KE) WeekdayWide(weekday time.Weekday) string {
	return dav.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (dav *dav_KE) WeekdaysWide() []string {
	return dav.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'dav_KE' and handles both Whole and Real numbers based on 'v'
func (dav *dav_KE) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'dav_KE' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (dav *dav_KE) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'dav_KE'
func (dav *dav_KE) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := dav.currencies[currency]
	l := len(s) + len(dav.decimal) + len(dav.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dav.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, dav.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, dav.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, dav.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'dav_KE'
// in accounting notation.
func (dav *dav_KE) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := dav.currencies[currency]
	l := len(s) + len(dav.decimal) + len(dav.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dav.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, dav.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, dav.currencyNegativePrefix[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, dav.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, dav.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'dav_KE'
func (dav *dav_KE) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'dav_KE'
func (dav *dav_KE) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, dav.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'dav_KE'
func (dav *dav_KE) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, dav.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'dav_KE'
func (dav *dav_KE) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, dav.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, dav.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'dav_KE'
func (dav *dav_KE) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dav.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'dav_KE'
func (dav *dav_KE) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dav.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dav.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'dav_KE'
func (dav *dav_KE) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dav.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dav.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'dav_KE'
func (dav *dav_KE) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dav.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dav.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := dav.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
