package my_MM

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type my_MM struct {
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
	currencyPositivePrefix string
	currencyNegativePrefix string
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

// New returns a new instance of translator for the 'my_MM' locale
func New() locales.Translator {
	return &my_MM{
		locale:                 "my_MM",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED ", "AFA ", "AFN ", "ALK ", "ALL ", "AMD ", "ANG ", "AOA ", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS ", "ATS ", "AUD ", "AWG ", "AZM ", "AZN ", "BAD ", "BAM ", "BAN ", "BBD ", "BDT ", "BEC ", "BEF ", "BEL ", "BGL ", "BGM ", "BGN ", "BGO ", "BHD ", "BIF ", "BMD ", "BND ", "BOB ", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "BRL ", "BRN ", "BRR ", "BRZ ", "BSD ", "BTN ", "BUK ", "BWP ", "BYB ", "BYR ", "BZD ", "CAD ", "CDF ", "CHE ", "CHF ", "CHW ", "CLE ", "CLF ", "CLP ", "CNX ", "CNY ", "COP ", "COU ", "CRC ", "CSD ", "CSK ", "CUC ", "CUP ", "CVE ", "CYP ", "CZK ", "DDM ", "DEM ", "DJF ", "DKK ", "DOP ", "DZD ", "ECS ", "ECV ", "EEK ", "EGP ", "ERN ", "ESA ", "ESB ", "ESP ", "ETB ", "EUR ", "FIM ", "FJD ", "FKP ", "FRF ", "GBP ", "GEK ", "GEL ", "GHC ", "GHS ", "GIP ", "GMD ", "GNF ", "GNS ", "GQE ", "GRD ", "GTQ ", "GWE ", "GWP ", "GYD ", "HKD ", "HNL ", "HRD ", "HRK ", "HTG ", "HUF ", "IDR ", "IEP ", "ILP ", "ILR ", "ILS ", "INR ", "IQD ", "IRR ", "ISJ ", "ISK ", "ITL ", "JMD ", "JOD ", "JPY ", "KES ", "KGS ", "KHR ", "KMF ", "KPW ", "KRH ", "KRO ", "KRW ", "KWD ", "KYD ", "KZT ", "LAK ", "LBP ", "LKR ", "LRD ", "LSL ", "LTL ", "LTT ", "LUC ", "LUF ", "LUL ", "LVL ", "LVR ", "LYD ", "MAD ", "MAF ", "MCF ", "MDC ", "MDL ", "MGA ", "MGF ", "MKD ", "MKN ", "MLF ", "MMK ", "MNT ", "MOP ", "MRO ", "MTL ", "MTP ", "MUR ", "MVP ", "MVR ", "MWK ", "MXN ", "MXP ", "MXV ", "MYR ", "MZE ", "MZM ", "MZN ", "NAD ", "NGN ", "NIC ", "NIO ", "NLG ", "NOK ", "NPR ", "NZD ", "OMR ", "PAB ", "PEI ", "PEN ", "PES ", "PGK ", "PHP ", "PKR ", "PLN ", "PLZ ", "PTE ", "PYG ", "QAR ", "RHD ", "ROL ", "RON ", "RSD ", "RUB ", "RUR ", "RWF ", "SAR ", "SBD ", "SCR ", "SDD ", "SDG ", "SDP ", "SEK ", "SGD ", "SHP ", "SIT ", "SKK ", "SLL ", "SOS ", "SRD ", "SRG ", "SSP ", "STD ", "SUR ", "SVC ", "SYP ", "SZL ", "THB ", "TJR ", "TJS ", "TMM ", "TMT ", "TND ", "TOP ", "TPE ", "TRL ", "TRY ", "TTD ", "TWD ", "TZS ", "UAH ", "UAK ", "UGS ", "UGX ", "USD ", "USN ", "USS ", "UYI ", "UYP ", "UYU ", "UZS ", "VEB ", "VEF ", "VND ", "VNN ", "VUV ", "WST ", "XAF ", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "XCD ", "XDR ", "XEU ", "XFO ", "XFU ", "XOF ", "XPD ", "XPF ", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER ", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR ", "ZMK ", "ZMW ", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "ဇန်", "ဖေ", "မတ်", "ဧပြီ", "မေ", "ဇွန်", "ဇူ", "ဩ", "စက်", "အောက်", "နို", "ဒီ"},
		monthsNarrow:           []string{"", "ဇ", "ဖ", "မ", "ဧ", "မ", "ဇ", "ဇ", "ဩ", "စ", "အ", "န", "ဒ"},
		monthsWide:             []string{"", "ဇန်နဝါရီ", "ဖေဖော်ဝါရီ", "မတ်", "ဧပြီ", "မေ", "ဇွန်", "ဇူလိုင်", "ဩဂုတ်", "စက်တင်ဘာ", "အောက်တိုဘာ", "နိုဝင်ဘာ", "ဒီဇင်ဘာ"},
		daysAbbreviated:        []string{"တနင်္ဂနွေ", "တနင်္လာ", "အင်္ဂါ", "ဗုဒ္ဓဟူး", "ကြာသပတေး", "သောကြာ", "စနေ"},
		daysNarrow:             []string{"တ", "တ", "အ", "ဗ", "က", "သ", "စ"},
		daysShort:              []string{"တနင်္ဂနွေ", "တနင်္လာ", "အင်္ဂါ", "ဗုဒ္ဓဟူး", "ကြာသပတေး", "သောကြာ", "စနေ"},
		daysWide:               []string{"တနင်္ဂနွေ", "တနင်္လာ", "အင်္ဂါ", "ဗုဒ္ဓဟူး", "ကြာသပတေး", "သောကြာ", "စနေ"},
		periodsAbbreviated:     []string{"နံနက်", "ညနေ"},
		periodsNarrow:          []string{"နံနက်", "ညနေ"},
		periodsWide:            []string{"နံနက်", "ညနေ"},
		erasAbbreviated:        []string{"ဘီစီ", "အေဒီ"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ခရစ်တော် မပေါ်မီကာလ", "ခရစ်တော် ပေါ်ထွန်းပြီးကာလ"},
		timezones:              map[string]string{"MDT": "မကာအို နွေရာသီ အချိန်", "AKDT": "အလာစကာနေ့ပိုင်းအချိန်", "AST": "အတ္ထလန်တစ် စံတော်ချိန်", "WIB": "အနောက်ပိုင်း အင်ဒိုနီးရှား အချိန်", "WAT": "အနောက်ပိုင်း အာဖရိက စံတော်ချိန်", "HNT": "နယူးဖောင်လန် စံတော်ချိန်", "MST": "မကာအို စံတော်ချိန်", "TMST": "တာခ်မီန့စ်တန်နွေရာသီအချိန်", "COT": "ကိုလံဘီယာ စံတော်ချိန်", "CAT": "အလယ်ပိုင်း အာဖရိက အချိန်", "HKT": "ဟောင်ကောင် စံတော်ချိန်", "GMT": "ဂရင်းနစ် စံတော်ချိန်", "UYST": "ဥရုဂွေး နွေရာသီ အချိန်", "WARST": "အနောက် အာဂျင်တီးနား နွေရာသီ အချိန်", "NZDT": "နယူးဇီလန် နေ့ပိုင်း အချိန်", "COST": "ကိုလံဘီယာ နွေရာသီ အချိန်", "∅∅∅": "အမေဇွန်နွေရာသီအချိန်", "EDT": "အရှေ့ပိုင်း အမေရိက နေ့ပိုင်း အချိန်", "PST": "ပစိဖိတ် စံတော်ချိန်", "WAST": "အနောက်ပိုင်း အာဖရိက နွေရာသီ အချိန်", "ADT": "အတ္ထလန်တစ် နေ့ပိုင်း အချိန်", "HAST": "ဟာဝိုင်အီ အာလူးရှန်စံတော်ချိန်", "AEST": "အရှေ့ပိုင်း ဩစတြေးလျှား စံတော်ချိန်", "GFT": "ပြင်သစ် ဂီယာနာ အချိန်", "AKST": "အလာစကာစံတော်ချိန်", "AWDT": "အနောက်ပိုင်း ဩစတြေးလျှား နေ့ပိုင်း အချိန်", "CLT": "ချီလီ စံတော်ချိန်", "WIT": "အရှေ့ပိုင်း အင်ဒိုနီးရှား အချိန်", "IST": "အိန္ဒြိယ စံတော်ချိန်", "BT": "ဘူတန် အချိန်", "ACWST": "အလယ်အနောက်ပိုင်း ဩစတြေးလျှား စံတော်ချိန်", "MEZ": "ဥရောပ အလယ်ပိုင်း စံတော်ချိန်", "HKST": "ဟောင်ကောင် နွေရာသီ အချိန်", "CHADT": "ချားသမ်နေ့ပိုင်းအချိန်", "OEZ": "အရှေ့ဥရောပ စံတော်ချိန်", "HAT": "နယူးဖောင်လန် နေ့ပိုင်း အချိန်", "CHAST": "ချားသမ်စံတော်ချိန်", "TMT": "တာခ်မီန့စ်တန်စံတော်ချိန်", "SRT": "စူးရီနာမ်အချိန်", "ACST": "အလယ်ပိုင်း ဩစတြေးလျှား စံတော်ချိန်", "ART": "အာဂျင်တီးနား စံတော်ချိန်", "CDT": "အလယ်ပိုင်း အမေရိက နွေရာသီ အချိန်", "SAST": "တောင်အာဖရိက အချိန်", "ACWDT": "အလယ်အနောက်ပိုင်း ဩစတြေးလျှား နေ့ပိုင်း အချိန်", "EAT": "အရှေ့ပိုင်း အာဖရိက အချိန်", "MESZ": "ဥရောပ အလယ်ပိုင်း နွေရာသီ အချိန်", "WEZ": "အနောက်ပိုင်း ဥရောပ စံတော်ချိန်", "WESZ": "အနောက်ပိုင်း ဥရောပ နွေရာသီ အချိန်", "LHDT": "လော့ဒ်ဟောင်နေ့ပိုင်းအချိန်", "WART": "အနောက် အာဂျင်တီးနား စံတော်ချိန်", "VET": "ဗင်နီဇွဲလား အချိန်", "OESZ": "အရှေ့ဥရောပ နွေရာသီ အချိန်", "JST": "ဂျပန် စံတော်ချိန်", "UYT": "ဥရုဂွေး စံတော်ချိန်", "PDT": "ပစိဖိတ် နေ့ပိုင်း အချိန်", "AEDT": "အရှေ့ပိုင်း ဩစတြေးလျှား နေ့ပိုင်း အချိန်", "WITA": "အလယ်ပိုင်း အင်ဒိုနီးရှား အချိန်", "JDT": "ဂျပန် နေ့ပိုင်း အချိန်", "MYT": "မလေးရှား အချိန်", "EST": "အရှေ့ပိုင်း အမေရိက စံတော်ချိန်", "LHST": "လော့ဒ်ဟောင်စံတော်ချိန်", "BOT": "ဘိုလီးဘီးယား အချိန်", "HADT": "ဟာဝိုင်အီ အာလူးရှန်နေ့ပိုင်းအချိန်", "ACDT": "အလယ်ပိုင်း ဩစတြေးလျှား နေ့ပိုင်း အချိန်", "GYT": "ဂိုင်ရာနားအချိန်", "ARST": "အာဂျင်တီးနား နွေရာသီ အချိန်", "CLST": "ချီလီ နွေရာသီ အချိန်", "CST": "အလယ်ပိုင်း အမေရိက စံတော်ချိန်", "SGT": "စင်္ကာပူ စံတော်ချိန်", "ECT": "အီကွေဒေါ အချိန်", "ChST": "ချာမိုရိုအချိန်", "AWST": "အနောက်ပိုင်း ဩစတြေးလျှား စံတော်ချိန်", "NZST": "နယူးဇီလန် စံတော်ချိန်"},
	}
}

// Locale returns the current translators string locale
func (my *my_MM) Locale() string {
	return my.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'my_MM'
func (my *my_MM) PluralsCardinal() []locales.PluralRule {
	return my.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'my_MM'
func (my *my_MM) PluralsOrdinal() []locales.PluralRule {
	return my.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'my_MM'
func (my *my_MM) PluralsRange() []locales.PluralRule {
	return my.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'my_MM'
func (my *my_MM) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'my_MM'
func (my *my_MM) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'my_MM'
func (my *my_MM) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (my *my_MM) MonthAbbreviated(month time.Month) string {
	return my.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (my *my_MM) MonthsAbbreviated() []string {
	return my.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (my *my_MM) MonthNarrow(month time.Month) string {
	return my.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (my *my_MM) MonthsNarrow() []string {
	return my.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (my *my_MM) MonthWide(month time.Month) string {
	return my.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (my *my_MM) MonthsWide() []string {
	return my.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (my *my_MM) WeekdayAbbreviated(weekday time.Weekday) string {
	return my.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (my *my_MM) WeekdaysAbbreviated() []string {
	return my.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (my *my_MM) WeekdayNarrow(weekday time.Weekday) string {
	return my.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (my *my_MM) WeekdaysNarrow() []string {
	return my.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (my *my_MM) WeekdayShort(weekday time.Weekday) string {
	return my.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (my *my_MM) WeekdaysShort() []string {
	return my.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (my *my_MM) WeekdayWide(weekday time.Weekday) string {
	return my.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (my *my_MM) WeekdaysWide() []string {
	return my.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'my_MM' and handles both Whole and Real numbers based on 'v'
func (my *my_MM) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(my.decimal) + len(my.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, my.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, my.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, my.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'my_MM' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (my *my_MM) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(my.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, my.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, my.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, my.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'my_MM'
func (my *my_MM) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := my.currencies[currency]
	l := len(s) + len(my.decimal) + len(my.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, my.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, my.group[0])
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

	for j := len(my.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, my.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, my.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, my.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'my_MM'
// in accounting notation.
func (my *my_MM) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := my.currencies[currency]
	l := len(s) + len(my.decimal) + len(my.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, my.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, my.group[0])
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

		for j := len(my.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, my.currencyNegativePrefix[j])
		}

		b = append(b, my.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(my.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, my.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, my.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'my_MM'
func (my *my_MM) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'my_MM'
func (my *my_MM) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, my.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'my_MM'
func (my *my_MM) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, my.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'my_MM'
func (my *my_MM) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, my.daysWide[t.Weekday()]...)
	b = append(b, []byte{0xe1, 0x81, 0x8a, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, my.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'my_MM'
func (my *my_MM) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, my.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'my_MM'
func (my *my_MM) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, my.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, my.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'my_MM'
func (my *my_MM) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, my.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, my.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'my_MM'
func (my *my_MM) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, my.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, my.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := my.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
