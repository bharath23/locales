package fi

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type fi struct {
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

// New returns a new instance of translator for the 'fi' locale
func New() locales.Translator {
	return &fi{
		locale:                 "fi",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		decimal:                []byte{0x2c},
		group:                  []byte{0xc2, 0xa0},
		minus:                  []byte{0xe2, 0x88, 0x92},
		percent:                []byte{0x25},
		perMille:               []byte{0xe2, 0x80, 0xb0},
		timeSeparator:          []byte{0x2e},
		inifinity:              []byte{0xe2, 0x88, 0x9e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x55, 0x44}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x42, 0x52, 0x4c}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x44}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58}, {0x43, 0x4e, 0x59}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x6d, 0x6b}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x44}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52}, {0x49, 0x4c, 0x53}, {0x49, 0x4e, 0x52}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0x4b, 0x52, 0x57}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x4e}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x54, 0x57, 0x44}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x24}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0x56, 0x4e, 0x44}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x58, 0x43, 0x44}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x58, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		percentSuffix:          []byte{0xc2, 0xa0},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x74, 0x61, 0x6d, 0x6d, 0x69, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x68, 0x65, 0x6c, 0x6d, 0x69, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x6d, 0x61, 0x61, 0x6c, 0x69, 0x73, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x68, 0x75, 0x68, 0x74, 0x69, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x74, 0x6f, 0x75, 0x6b, 0x6f, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x68, 0x65, 0x69, 0x6e, 0xc3, 0xa4, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x65, 0x6c, 0x6f, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x73, 0x79, 0x79, 0x73, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x6c, 0x6f, 0x6b, 0x61, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x6d, 0x61, 0x72, 0x72, 0x61, 0x73, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x6a, 0x6f, 0x75, 0x6c, 0x75, 0x6b, 0x75, 0x75, 0x74, 0x61}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x54}, {0x48}, {0x4d}, {0x48}, {0x54}, {0x4b}, {0x48}, {0x45}, {0x53}, {0x4c}, {0x4d}, {0x4a}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x74, 0x61, 0x6d, 0x6d, 0x69, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x68, 0x65, 0x6c, 0x6d, 0x69, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x6d, 0x61, 0x61, 0x6c, 0x69, 0x73, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x68, 0x75, 0x68, 0x74, 0x69, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x74, 0x6f, 0x75, 0x6b, 0x6f, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x68, 0x65, 0x69, 0x6e, 0xc3, 0xa4, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x65, 0x6c, 0x6f, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x73, 0x79, 0x79, 0x73, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x6c, 0x6f, 0x6b, 0x61, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x6d, 0x61, 0x72, 0x72, 0x61, 0x73, 0x6b, 0x75, 0x75, 0x74, 0x61}, {0x6a, 0x6f, 0x75, 0x6c, 0x75, 0x6b, 0x75, 0x75, 0x74, 0x61}},
		daysAbbreviated:        [][]uint8{{0x73, 0x75}, {0x6d, 0x61}, {0x74, 0x69}, {0x6b, 0x65}, {0x74, 0x6f}, {0x70, 0x65}, {0x6c, 0x61}},
		daysNarrow:             [][]uint8{{0x53}, {0x4d}, {0x54}, {0x4b}, {0x54}, {0x50}, {0x4c}},
		daysShort:              [][]uint8{{0x73, 0x75}, {0x6d, 0x61}, {0x74, 0x69}, {0x6b, 0x65}, {0x74, 0x6f}, {0x70, 0x65}, {0x6c, 0x61}},
		daysWide:               [][]uint8{{0x73, 0x75, 0x6e, 0x6e, 0x75, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x61}, {0x6d, 0x61, 0x61, 0x6e, 0x61, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x61}, {0x74, 0x69, 0x69, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x61}, {0x6b, 0x65, 0x73, 0x6b, 0x69, 0x76, 0x69, 0x69, 0x6b, 0x6b, 0x6f, 0x6e, 0x61}, {0x74, 0x6f, 0x72, 0x73, 0x74, 0x61, 0x69, 0x6e, 0x61}, {0x70, 0x65, 0x72, 0x6a, 0x61, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x61}, {0x6c, 0x61, 0x75, 0x61, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x61}},
		periodsAbbreviated:     [][]uint8{{0x61, 0x70, 0x2e}, {0x69, 0x70, 0x2e}},
		periodsNarrow:          [][]uint8{{0x61, 0x70, 0x2e}, {0x69, 0x70, 0x2e}},
		periodsWide:            [][]uint8{{0x61, 0x70, 0x2e}, {0x69, 0x70, 0x2e}},
		erasAbbreviated:        [][]uint8{{0x65, 0x4b, 0x72, 0x2e}, {0x6a, 0x4b, 0x72, 0x2e}},
		erasNarrow:             [][]uint8{{0x65, 0x4b}, {0x6a, 0x4b}},
		erasWide:               [][]uint8{{0x65, 0x6e, 0x6e, 0x65, 0x6e, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x75, 0x6b, 0x73, 0x65, 0x6e, 0x20, 0x73, 0x79, 0x6e, 0x74, 0x79, 0x6d, 0xc3, 0xa4, 0xc3, 0xa4}, {0x6a, 0xc3, 0xa4, 0x6c, 0x6b, 0x65, 0x65, 0x6e, 0x20, 0x4b, 0x72, 0x69, 0x73, 0x74, 0x75, 0x6b, 0x73, 0x65, 0x6e, 0x20, 0x73, 0x79, 0x6e, 0x74, 0x79, 0x6d, 0xc3, 0xa4, 0x6e}},
		timezones:              map[string][]uint8{"BT": {0x42, 0x68, 0x75, 0x74, 0x61, 0x6e, 0x69, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "WIT": {0x49, 0x74, 0xc3, 0xa4, 0x2d, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "AEST": {0x49, 0x74, 0xc3, 0xa4, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "VET": {0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "CHAST": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x69, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "GFT": {0x52, 0x61, 0x6e, 0x73, 0x6b, 0x61, 0x6e, 0x20, 0x47, 0x75, 0x61, 0x79, 0x61, 0x6e, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "MESZ": {0x4b, 0x65, 0x73, 0x6b, 0x69, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "ChST": {0x54, 0xc5, 0xa1, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "WEZ": {0x4c, 0xc3, 0xa4, 0x6e, 0x73, 0x69, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "JDT": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x69, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "OESZ": {0x49, 0x74, 0xc3, 0xa4, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "MDT": {0x4b, 0x61, 0x6c, 0x6c, 0x69, 0x6f, 0x76, 0x75, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "IST": {0x49, 0x6e, 0x74, 0x69, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "CST": {0x59, 0x68, 0x64, 0x79, 0x73, 0x76, 0x61, 0x6c, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0x6b, 0x69, 0x6e, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "HAT": {0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "SAST": {0x45, 0x74, 0x65, 0x6c, 0xc3, 0xa4, 0x2d, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "LHST": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "ACDT": {0x4b, 0x65, 0x73, 0x6b, 0x69, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "ECT": {0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72, 0x69, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "COST": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "PST": {0x59, 0x68, 0x64, 0x79, 0x73, 0x76, 0x61, 0x6c, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x54, 0x79, 0x79, 0x6e, 0x65, 0x6e, 0x6d, 0x65, 0x72, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "OEZ": {0x49, 0x74, 0xc3, 0xa4, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "CLT": {0x43, 0x68, 0x69, 0x6c, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "CHADT": {0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x69, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "AKST": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "AWST": {0x4c, 0xc3, 0xa4, 0x6e, 0x73, 0x69, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "WITA": {0x4b, 0x65, 0x73, 0x6b, 0x69, 0x2d, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "MEZ": {0x4b, 0x65, 0x73, 0x6b, 0x69, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "BOT": {0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "WAST": {0x4c, 0xc3, 0xa4, 0x6e, 0x73, 0x69, 0x2d, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "AWDT": {0x4c, 0xc3, 0xa4, 0x6e, 0x73, 0x69, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "COT": {0x4b, 0x6f, 0x6c, 0x75, 0x6d, 0x62, 0x69, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "TMST": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x69, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "GYT": {0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "ADT": {0x4b, 0x61, 0x6e, 0x61, 0x64, 0x61, 0x6e, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "UYST": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "MST": {0x4b, 0x61, 0x6c, 0x6c, 0x69, 0x6f, 0x76, 0x75, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "EAT": {0x49, 0x74, 0xc3, 0xa4, 0x2d, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "NZST": {0x55, 0x75, 0x64, 0x65, 0x6e, 0x2d, 0x53, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "SGT": {0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x6f, 0x72, 0x65, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "WESZ": {0x4c, 0xc3, 0xa4, 0x6e, 0x73, 0x69, 0x2d, 0x45, 0x75, 0x72, 0x6f, 0x6f, 0x70, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "HNT": {0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "CAT": {0x4b, 0x65, 0x73, 0x6b, 0x69, 0x2d, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "WART": {0x4c, 0xc3, 0xa4, 0x6e, 0x73, 0x69, 0x2d, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x69, 0x6e, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "PDT": {0x59, 0x68, 0x64, 0x79, 0x73, 0x76, 0x61, 0x6c, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x54, 0x79, 0x79, 0x6e, 0x65, 0x6e, 0x6d, 0x65, 0x72, 0x65, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "HAST": {0x48, 0x61, 0x76, 0x61, 0x69, 0x6a, 0x69, 0x6e, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x74, 0x69, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "AST": {0x4b, 0x61, 0x6e, 0x61, 0x64, 0x61, 0x6e, 0x20, 0x41, 0x74, 0x6c, 0x61, 0x6e, 0x74, 0x69, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "ART": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x69, 0x6e, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "∅∅∅": {0x41, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x69, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "UYT": {0x55, 0x72, 0x75, 0x67, 0x75, 0x61, 0x79, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "EST": {0x59, 0x68, 0x64, 0x79, 0x73, 0x76, 0x61, 0x6c, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x69, 0x74, 0xc3, 0xa4, 0x69, 0x6e, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "ACST": {0x4b, 0x65, 0x73, 0x6b, 0x69, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "ARST": {0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x69, 0x6e, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "HADT": {0x48, 0x61, 0x76, 0x61, 0x69, 0x6a, 0x69, 0x6e, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x74, 0x69, 0x65, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "MYT": {0x4d, 0x61, 0x6c, 0x65, 0x73, 0x69, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "GMT": {0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68, 0x69, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "HKST": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x69, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "TMT": {0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x69, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "ACWDT": {0x4c, 0xc3, 0xa4, 0x6e, 0x74, 0x69, 0x73, 0x65, 0x6e, 0x20, 0x4b, 0x65, 0x73, 0x6b, 0x69, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "CLST": {0x43, 0x68, 0x69, 0x6c, 0x65, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "CDT": {0x59, 0x68, 0x64, 0x79, 0x73, 0x76, 0x61, 0x6c, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0x6b, 0x69, 0x6e, 0x65, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "EDT": {0x59, 0x68, 0x64, 0x79, 0x73, 0x76, 0x61, 0x6c, 0x74, 0x61, 0x69, 0x6e, 0x20, 0x69, 0x74, 0xc3, 0xa4, 0x69, 0x6e, 0x65, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "SRT": {0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d, 0x69, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "AKDT": {0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "NZDT": {0x55, 0x75, 0x64, 0x65, 0x6e, 0x2d, 0x53, 0x65, 0x65, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "ACWST": {0x4c, 0xc3, 0xa4, 0x6e, 0x74, 0x69, 0x73, 0x65, 0x6e, 0x20, 0x4b, 0x65, 0x73, 0x6b, 0x69, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "WARST": {0x4c, 0xc3, 0xa4, 0x6e, 0x73, 0x69, 0x2d, 0x41, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x69, 0x6e, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "HKT": {0x48, 0x6f, 0x6e, 0x67, 0x6b, 0x6f, 0x6e, 0x67, 0x69, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "JST": {0x4a, 0x61, 0x70, 0x61, 0x6e, 0x69, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}, "AEDT": {0x49, 0x74, 0xc3, 0xa4, 0x2d, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "WIB": {0x4c, 0xc3, 0xa4, 0x6e, 0x73, 0x69, 0x2d, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x6e, 0x20, 0x61, 0x69, 0x6b, 0x61}, "LHDT": {0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65, 0x6e, 0x20, 0x6b, 0x65, 0x73, 0xc3, 0xa4, 0x61, 0x69, 0x6b, 0x61}, "WAT": {0x4c, 0xc3, 0xa4, 0x6e, 0x73, 0x69, 0x2d, 0x41, 0x66, 0x72, 0x69, 0x6b, 0x61, 0x6e, 0x20, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x61, 0x6c, 0x69, 0x61, 0x69, 0x6b, 0x61}},
	}
}

// Locale returns the current translators string locale
func (fi *fi) Locale() string {
	return fi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'fi'
func (fi *fi) PluralsCardinal() []locales.PluralRule {
	return fi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'fi'
func (fi *fi) PluralsOrdinal() []locales.PluralRule {
	return fi.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'fi'
func (fi *fi) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'fi'
func (fi *fi) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'fi'
func (fi *fi) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (fi *fi) MonthAbbreviated(month time.Month) []byte {
	return fi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (fi *fi) MonthsAbbreviated() [][]byte {
	return fi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (fi *fi) MonthNarrow(month time.Month) []byte {
	return fi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (fi *fi) MonthsNarrow() [][]byte {
	return fi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (fi *fi) MonthWide(month time.Month) []byte {
	return fi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (fi *fi) MonthsWide() [][]byte {
	return fi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (fi *fi) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return fi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (fi *fi) WeekdaysAbbreviated() [][]byte {
	return fi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (fi *fi) WeekdayNarrow(weekday time.Weekday) []byte {
	return fi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (fi *fi) WeekdaysNarrow() [][]byte {
	return fi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (fi *fi) WeekdayShort(weekday time.Weekday) []byte {
	return fi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (fi *fi) WeekdaysShort() [][]byte {
	return fi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (fi *fi) WeekdayWide(weekday time.Weekday) []byte {
	return fi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (fi *fi) WeekdaysWide() [][]byte {
	return fi.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'fi' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(fi.decimal) + len(fi.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fi.group) - 1; j >= 0; j-- {
					b = append(b, fi.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fi.minus) - 1; j >= 0; j-- {
			b = append(b, fi.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'fi' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (fi *fi) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(fi.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fi.minus) - 1; j >= 0; j-- {
			b = append(b, fi.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, fi.percentSuffix...)

	b = append(b, fi.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'fi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fi.currencies[currency]
	l := len(s) + len(fi.decimal) + len(fi.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fi.group) - 1; j >= 0; j-- {
					b = append(b, fi.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(fi.minus) - 1; j >= 0; j-- {
			b = append(b, fi.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, fi.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'fi'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := fi.currencies[currency]
	l := len(s) + len(fi.decimal) + len(fi.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, fi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(fi.group) - 1; j >= 0; j-- {
					b = append(b, fi.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(fi.minus) - 1; j >= 0; j-- {
			b = append(b, fi.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, fi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, fi.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, fi.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'fi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'fi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'fi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, fi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'fi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, []byte{0x63, 0x63, 0x63, 0x63, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, fi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'fi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'fi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'fi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'fi'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (fi *fi) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := fi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
