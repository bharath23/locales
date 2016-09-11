package sn_ZW

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type sn_ZW struct {
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

// New returns a new instance of translator for the 'sn_ZW' locale
func New() locales.Translator {
	return &sn_ZW{
		locale:                 "sn_ZW",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		timeSeparator:          ":",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Ndi", "Kuk", "Kur", "Kub", "Chv", "Chk", "Chg", "Nya", "Gun", "Gum", "Mb", "Zvi"},
		monthsNarrow:           []string{"", "N", "K", "K", "K", "C", "C", "C", "N", "G", "G", "M", "Z"},
		monthsWide:             []string{"", "Ndira", "Kukadzi", "Kurume", "Kubvumbi", "Chivabvu", "Chikumi", "Chikunguru", "Nyamavhuvhu", "Gunyana", "Gumiguru", "Mbudzi", "Zvita"},
		daysAbbreviated:        []string{"Svo", "Muv", "Chip", "Chit", "Chin", "Chis", "Mug"},
		daysNarrow:             []string{"S", "M", "C", "C", "C", "C", "M"},
		daysWide:               []string{"Svondo", "Muvhuro", "Chipiri", "Chitatu", "China", "Chishanu", "Mugovera"},
		erasAbbreviated:        []string{"BC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Kristo asati auya", "Kristo ashaya"},
		timezones:              map[string]string{"MESZ": "MESZ", "TMT": "TMT", "TMST": "TMST", "∅∅∅": "∅∅∅", "MYT": "MYT", "HAST": "HAST", "LHST": "LHST", "WART": "WART", "AEDT": "AEDT", "SAST": "SAST", "MEZ": "MEZ", "GMT": "GMT", "PDT": "PDT", "ARST": "ARST", "MST": "MST", "HAT": "HAT", "COST": "COST", "UYST": "UYST", "AST": "AST", "MDT": "MDT", "CLST": "CLST", "WIT": "WIT", "OESZ": "OESZ", "IST": "IST", "PST": "PST", "HNT": "HNT", "HKT": "HKT", "OEZ": "OEZ", "CDT": "CDT", "EDT": "EDT", "SRT": "SRT", "GYT": "GYT", "JST": "JST", "BT": "BT", "ACDT": "ACDT", "AWDT": "AWDT", "CLT": "CLT", "HADT": "HADT", "AKDT": "AKDT", "WEZ": "WEZ", "LHDT": "LHDT", "AWST": "AWST", "ACWDT": "ACWDT", "WAT": "WAT", "NZST": "NZST", "NZDT": "NZDT", "ART": "ART", "CAT": "CAT", "GFT": "GFT", "EST": "EST", "WARST": "WARST", "WAST": "WAST", "COT": "COT", "ChST": "ChST", "EAT": "EAT", "HKST": "HKST", "CHAST": "CHAST", "SGT": "SGT", "WITA": "WITA", "JDT": "JDT", "BOT": "BOT", "ADT": "ADT", "CHADT": "CHADT", "AEST": "AEST", "ECT": "ECT", "WESZ": "WESZ", "UYT": "UYT", "CST": "CST", "WIB": "WIB", "AKST": "AKST", "VET": "VET", "ACWST": "ACWST", "ACST": "ACST"},
	}
}

// Locale returns the current translators string locale
func (sn *sn_ZW) Locale() string {
	return sn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sn_ZW'
func (sn *sn_ZW) PluralsCardinal() []locales.PluralRule {
	return sn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sn_ZW'
func (sn *sn_ZW) PluralsOrdinal() []locales.PluralRule {
	return sn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sn_ZW'
func (sn *sn_ZW) PluralsRange() []locales.PluralRule {
	return sn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sn_ZW'
func (sn *sn_ZW) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sn_ZW'
func (sn *sn_ZW) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sn_ZW'
func (sn *sn_ZW) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sn *sn_ZW) MonthAbbreviated(month time.Month) string {
	return sn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sn *sn_ZW) MonthsAbbreviated() []string {
	return sn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sn *sn_ZW) MonthNarrow(month time.Month) string {
	return sn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sn *sn_ZW) MonthsNarrow() []string {
	return sn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sn *sn_ZW) MonthWide(month time.Month) string {
	return sn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sn *sn_ZW) MonthsWide() []string {
	return sn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sn *sn_ZW) WeekdayAbbreviated(weekday time.Weekday) string {
	return sn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sn *sn_ZW) WeekdaysAbbreviated() []string {
	return sn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sn *sn_ZW) WeekdayNarrow(weekday time.Weekday) string {
	return sn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sn *sn_ZW) WeekdaysNarrow() []string {
	return sn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sn *sn_ZW) WeekdayShort(weekday time.Weekday) string {
	return sn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sn *sn_ZW) WeekdaysShort() []string {
	return sn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sn *sn_ZW) WeekdayWide(weekday time.Weekday) string {
	return sn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sn *sn_ZW) WeekdaysWide() []string {
	return sn.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sn_ZW' and handles both Whole and Real numbers based on 'v'
func (sn *sn_ZW) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sn_ZW' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sn *sn_ZW) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sn_ZW'
func (sn *sn_ZW) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sn.currencies[currency]
	l := len(s) + len(sn.decimal) + len(sn.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sn.group[0])
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
		b = append(b, sn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sn_ZW'
// in accounting notation.
func (sn *sn_ZW) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sn.currencies[currency]
	l := len(s) + len(sn.decimal) + len(sn.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sn.group[0])
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

		b = append(b, sn.currencyNegativePrefix[0])

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
			b = append(b, sn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sn.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sn_ZW'
func (sn *sn_ZW) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'sn_ZW'
func (sn *sn_ZW) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sn.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sn_ZW'
func (sn *sn_ZW) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sn_ZW'
func (sn *sn_ZW) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sn.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sn_ZW'
func (sn *sn_ZW) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sn.periodsAbbreviated[0]...)
	} else {
		b = append(b, sn.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sn_ZW'
func (sn *sn_ZW) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sn.periodsAbbreviated[0]...)
	} else {
		b = append(b, sn.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sn_ZW'
func (sn *sn_ZW) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sn.periodsAbbreviated[0]...)
	} else {
		b = append(b, sn.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sn_ZW'
func (sn *sn_ZW) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sn.periodsAbbreviated[0]...)
	} else {
		b = append(b, sn.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sn.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
