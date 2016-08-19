package ro_MD

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ro_MD struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
	percentSuffix          []byte
	perMille               []byte
	timeSeparator          []byte
	inifinity              []byte
	currencies             [][]byte // idx = enum of currency code
	currencyPositiveSuffix []byte
	currencyNegativePrefix []byte
	currencyNegativeSuffix []byte
	monthsAbbreviated      [][]byte
	monthsNarrow           [][]byte
	monthsWide             [][]byte
	daysAbbreviated        [][]byte
	daysNarrow             [][]byte
	daysShort              [][]byte
	daysWide               [][]byte
	periodsAbbreviated     [][]byte
	periodsNarrow          [][]byte
	periodsShort           [][]byte
	periodsWide            [][]byte
	erasAbbreviated        [][]byte
	erasNarrow             [][]byte
	erasWide               [][]byte
	timezones              map[string][]byte
}

// New returns a new instance of translator for the 'ro_MD' locale
func New() locales.Translator {
	return &ro_MD{
		locale:                 "ro_MD",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		decimal:                []byte{0x2c},
		group:                  []byte{0x2e},
		minus:                  []byte{0x2d},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x3a},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4c}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0xc2, 0xa0, 0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x69, 0x61, 0x6e, 0x2e}, {0x66, 0x65, 0x62, 0x2e}, {0x6d, 0x61, 0x72, 0x2e}, {0x61, 0x70, 0x72, 0x2e}, {0x6d, 0x61, 0x69}, {0x69, 0x75, 0x6e, 0x2e}, {0x69, 0x75, 0x6c, 0x2e}, {0x61, 0x75, 0x67, 0x2e}, {0x73, 0x65, 0x70, 0x74, 0x2e}, {0x6f, 0x63, 0x74, 0x2e}, {0x6e, 0x6f, 0x76, 0x2e}, {0x64, 0x65, 0x63, 0x2e}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x49}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x49}, {0x49}, {0x41}, {0x53}, {0x4f}, {0x4e}, {0x44}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x69, 0x61, 0x6e, 0x75, 0x61, 0x72, 0x69, 0x65}, {0x66, 0x65, 0x62, 0x72, 0x75, 0x61, 0x72, 0x69, 0x65}, {0x6d, 0x61, 0x72, 0x74, 0x69, 0x65}, {0x61, 0x70, 0x72, 0x69, 0x6c, 0x69, 0x65}, {0x6d, 0x61, 0x69}, {0x69, 0x75, 0x6e, 0x69, 0x65}, {0x69, 0x75, 0x6c, 0x69, 0x65}, {0x61, 0x75, 0x67, 0x75, 0x73, 0x74}, {0x73, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62, 0x72, 0x69, 0x65}, {0x6f, 0x63, 0x74, 0x6f, 0x6d, 0x62, 0x72, 0x69, 0x65}, {0x6e, 0x6f, 0x69, 0x65, 0x6d, 0x62, 0x72, 0x69, 0x65}, {0x64, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x72, 0x69, 0x65}},
		daysAbbreviated:        [][]uint8{{0x44, 0x75, 0x6d}, {0x4c, 0x75, 0x6e}, {0x4d, 0x61, 0x72}, {0x4d, 0x69, 0x65}, {0x4a, 0x6f, 0x69}, {0x56, 0x69, 0x6e}, {0x53, 0xc3, 0xa2, 0x6d}},
		daysNarrow:             [][]uint8{{0x44}, {0x4c}, {0x4d, 0x61}, {0x4d, 0x69}, {0x4a}, {0x56}, {0x53}},
		daysShort:              [][]uint8{{0x44, 0x75}, {0x4c, 0x75}, {0x4d, 0x61}, {0x4d, 0x69}, {0x4a, 0x6f}, {0x56, 0x69}, {0x53, 0xc3, 0xa2}},
		daysWide:               [][]uint8{{0x64, 0x75, 0x6d, 0x69, 0x6e, 0x69, 0x63, 0xc4, 0x83}, {0x6c, 0x75, 0x6e, 0x69}, {0x6d, 0x61, 0x72, 0xc8, 0x9b, 0x69}, {0x6d, 0x69, 0x65, 0x72, 0x63, 0x75, 0x72, 0x69}, {0x6a, 0x6f, 0x69}, {0x76, 0x69, 0x6e, 0x65, 0x72, 0x69}, {0x73, 0xc3, 0xa2, 0x6d, 0x62, 0xc4, 0x83, 0x74, 0xc4, 0x83}},
		periodsAbbreviated:     [][]uint8{{0x61, 0x2e, 0x6d, 0x2e}, {0x70, 0x2e, 0x6d, 0x2e}},
		periodsNarrow:          [][]uint8{[]uint8(nil), []uint8(nil)},
		periodsWide:            [][]uint8{[]uint8(nil), []uint8(nil)},
		erasAbbreviated:        [][]uint8{[]uint8(nil), []uint8(nil)},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{[]uint8(nil), []uint8(nil)},
		timezones:              map[string][]uint8{"BOT": {0x4f, 0x72, 0x61, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x65, 0x69}, "AST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0xc3, 0xae, 0x6e, 0x20, 0x7a, 0x6f, 0x6e, 0x61, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0xc4, 0x83}, "WEZ": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x69, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x73, 0x74}, "SGT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x65}, "JST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x4a, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x65, 0x69}, "SAST": {0x4f, 0x72, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x69, 0x69, 0x20, 0x4d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x65}, "ART": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x69}, "PST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0xc3, 0xae, 0x6e, 0x20, 0x7a, 0x6f, 0x6e, 0x61, 0x20, 0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0xc4, 0x83}, "HADT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x6e, 0x65}, "BT": {0x4f, 0x72, 0x61, 0x20, 0x42, 0x68, 0x75, 0x74, 0x61, 0x6e, 0x75, 0x6c, 0x75, 0x69}, "COT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x65, 0x69}, "WITA": {0x4f, 0x72, 0x61, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x65, 0x69, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "WAST": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x69, 0x69, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "WIT": {0x4f, 0x72, 0x61, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x65, 0x69, 0x20, 0x64, 0x65, 0x20, 0x45, 0x73, 0x74}, "SRT": {0x4f, 0x72, 0x61, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x75, 0x6c, 0x75, 0x69}, "MST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0xc3, 0xae, 0x6e, 0x20, 0x7a, 0x6f, 0x6e, 0x61, 0x20, 0x6d, 0x6f, 0x6e, 0x74, 0x61, 0x6e, 0xc4, 0x83, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0xc4, 0x83}, "AWDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65, 0x69, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "IST": {0x4f, 0x72, 0x61, 0x20, 0x49, 0x6e, 0x64, 0x69, 0x65, 0x69}, "ACWST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65, 0x69, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "LHST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "CDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0xc4, 0x83, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0xc4, 0x83}, "CAT": {0x4f, 0x72, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x69, 0x69, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "WAT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x69, 0x69, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "NZDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x4e, 0x6f, 0x69, 0x69, 0x20, 0x5a, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x65}, "LHDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "CLST": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "EDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0xc4, 0x83, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0xc4, 0x83}, "GMT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x68, 0x69, 0x63, 0x68}, "GYT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "ECT": {0x4f, 0x72, 0x61, 0x20, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72, 0x75, 0x6c, 0x75, 0x69}, "AEST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65, 0x69, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "TMT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "ChST": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "WIB": {0x4f, 0x72, 0x61, 0x20, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x7a, 0x69, 0x65, 0x69, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x73, 0x74}, "UYT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x75, 0x6c, 0x75, 0x69}, "HAT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "ACST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65, 0x69, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "WARST": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x69, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "∅∅∅": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x41, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x75, 0x6c, 0x75, 0x69}, "JDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x4a, 0x61, 0x70, 0x6f, 0x6e, 0x69, 0x65, 0x69}, "CLT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "UYST": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x75, 0x6c, 0x75, 0x69}, "WART": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x69, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "AKST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "HKST": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "MYT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x79, 0x73, 0x69, 0x61}, "CST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0xc4, 0x83, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0xc4, 0x83}, "EST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0xc4, 0x83, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0xc4, 0x83}, "MESZ": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x69, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "OESZ": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x69, 0x20, 0x64, 0x65, 0x20, 0x45, 0x73, 0x74}, "VET": {0x4f, 0x72, 0x61, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x65, 0x69}, "TMST": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e}, "WESZ": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x69, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x73, 0x74}, "NZST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x4e, 0x6f, 0x69, 0x69, 0x20, 0x5a, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x65}, "MEZ": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x69, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "AEDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65, 0x69, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "HKT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x48, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "MDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0xc3, 0xae, 0x6e, 0x20, 0x7a, 0x6f, 0x6e, 0x61, 0x20, 0x6d, 0x6f, 0x6e, 0x74, 0x61, 0x6e, 0xc4, 0x83, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0xc4, 0x83}, "ADT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0xc3, 0xae, 0x6e, 0x20, 0x7a, 0x6f, 0x6e, 0x61, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0xc4, 0x83}, "ACDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65, 0x69, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x65}, "CHADT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "AKDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "HNT": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "GFT": {0x4f, 0x72, 0x61, 0x20, 0x47, 0x75, 0x69, 0x61, 0x6e, 0x65, 0x69, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x7a, 0x65}, "AWST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65, 0x69, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "ACWDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x65, 0x69, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "CHAST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "OEZ": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x61, 0x20, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x65, 0x69, 0x20, 0x64, 0x65, 0x20, 0x45, 0x73, 0x74}, "HAST": {0x4f, 0x72, 0x61, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x20, 0x64, 0x69, 0x6e, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x6e, 0x65}, "PDT": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0xc3, 0xae, 0x6e, 0x20, 0x7a, 0x6f, 0x6e, 0x61, 0x20, 0x50, 0x61, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x6e, 0x6f, 0x72, 0x64, 0x2d, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0xc4, 0x83}, "EAT": {0x4f, 0x72, 0x61, 0x20, 0x41, 0x66, 0x72, 0x69, 0x63, 0x69, 0x69, 0x20, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x65}, "COST": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x65, 0x69}, "ARST": {0x4f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x76, 0x61, 0x72, 0xc4, 0x83, 0x20, 0x61, 0x20, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x69}},
	}
}

// Locale returns the current translators string locale
func (ro *ro_MD) Locale() string {
	return ro.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ro_MD'
func (ro *ro_MD) PluralsCardinal() []locales.PluralRule {
	return ro.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ro_MD'
func (ro *ro_MD) PluralsOrdinal() []locales.PluralRule {
	return ro.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ro_MD'
func (ro *ro_MD) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	nMod100 := math.Mod(n, 100)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if (v != 0) || (n == 0) || (n != 1 && nMod100 >= 1 && nMod100 <= 19) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ro_MD'
func (ro *ro_MD) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ro_MD'
func (ro *ro_MD) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ro.CardinalPluralRule(num1, v1)
	end := ro.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ro *ro_MD) MonthAbbreviated(month time.Month) []byte {
	return ro.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ro *ro_MD) MonthsAbbreviated() [][]byte {
	return ro.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ro *ro_MD) MonthNarrow(month time.Month) []byte {
	return ro.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ro *ro_MD) MonthsNarrow() [][]byte {
	return ro.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ro *ro_MD) MonthWide(month time.Month) []byte {
	return ro.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ro *ro_MD) MonthsWide() [][]byte {
	return ro.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return ro.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ro *ro_MD) WeekdaysAbbreviated() [][]byte {
	return ro.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayNarrow(weekday time.Weekday) []byte {
	return ro.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ro *ro_MD) WeekdaysNarrow() [][]byte {
	return ro.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayShort(weekday time.Weekday) []byte {
	return ro.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ro *ro_MD) WeekdaysShort() [][]byte {
	return ro.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayWide(weekday time.Weekday) []byte {
	return ro.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ro *ro_MD) WeekdaysWide() [][]byte {
	return ro.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ro_MD' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ro.decimal) + len(ro.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ro.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ro.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ro_MD' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ro *ro_MD) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ro.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ro.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ro.percentSuffix...)

	b = append(b, ro.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ro_MD'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ro.currencies[currency]
	l := len(s) + len(ro.decimal) + len(ro.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ro.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ro.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ro.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ro.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ro_MD'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ro.currencies[currency]
	l := len(s) + len(ro.decimal) + len(ro.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ro.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(ro.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ro.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ro.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ro.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ro.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ro_MD'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'ro_MD'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ro.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ro_MD'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ro.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ro_MD'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ro.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ro.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ro_MD'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ro_MD'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ro_MD'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ro_MD'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ro *ro_MD) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ro.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
