package ast

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ast struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	decimal                []byte
	group                  []byte
	minus                  []byte
	percent                []byte
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

// New returns a new instance of translator for the 'ast' locale
func New() locales.Translator {
	return &ast{
		locale:                 "ast",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		decimal:                []byte{0xd9, 0xab},
		group:                  []byte{0xd9, 0xac},
		minus:                  []byte{0xe2, 0x80, 0x8f, 0x2d},
		percent:                []byte{0xd9, 0xaa},
		perMille:               []byte{0xd8, 0x89},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b}, {0x41, 0x4f, 0x4e}, {0x41, 0x4f, 0x52}, {0x41, 0x52, 0x41}, {0x41, 0x52, 0x4c}, {0x41, 0x52, 0x4d}, {0x41, 0x52, 0x50}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53}, {0x41, 0x24}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43}, {0x42, 0x45, 0x46}, {0x42, 0x45, 0x4c}, {0x42, 0x47, 0x4c}, {0x42, 0x47, 0x4d}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c}, {0x42, 0x4f, 0x50}, {0x42, 0x4f, 0x56}, {0x42, 0x52, 0x42}, {0x42, 0x52, 0x43}, {0x42, 0x52, 0x45}, {0x52, 0x24}, {0x42, 0x52, 0x4e}, {0x42, 0x52, 0x52}, {0x42, 0x52, 0x5a}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x24}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57}, {0x43, 0x4c, 0x45}, {0x43, 0x4c, 0x46}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44}, {0x43, 0x53, 0x4b}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d}, {0x44, 0x45, 0x4d}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53}, {0x45, 0x43, 0x56}, {0x45, 0x45, 0x4b}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41}, {0x45, 0x53, 0x42}, {0x45, 0x53, 0x50}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d}, {0x46, 0x4a, 0x44}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53}, {0x47, 0x51, 0x45}, {0x47, 0x52, 0x44}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45}, {0x47, 0x57, 0x50}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50}, {0x49, 0x4c, 0x50}, {0x49, 0x4c, 0x52, 0x20}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48}, {0x4b, 0x52, 0x4f}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c}, {0x4c, 0x54, 0x4c}, {0x4c, 0x54, 0x54}, {0x4c, 0x55, 0x43}, {0x4c, 0x55, 0x46}, {0x4c, 0x55, 0x4c}, {0x4c, 0x56, 0x4c}, {0x4c, 0x56, 0x52}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46}, {0x4d, 0x43, 0x46}, {0x4d, 0x44, 0x43}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e}, {0x4d, 0x4c, 0x46}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c}, {0x4d, 0x54, 0x50}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x24}, {0x4d, 0x58, 0x50}, {0x4d, 0x58, 0x56}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45}, {0x4d, 0x5a, 0x4d}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x24}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53}, {0x50, 0x47, 0x4b}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a}, {0x50, 0x54, 0x45}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44}, {0x52, 0x4f, 0x4c}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54}, {0x53, 0x4b, 0x4b}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52}, {0x53, 0x56, 0x43}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0xe0, 0xb8, 0xbf}, {0x54, 0x4a, 0x52}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50}, {0x54, 0x50, 0x45}, {0x54, 0x52, 0x4c}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b}, {0x55, 0x47, 0x53}, {0x55, 0x47, 0x58}, {0x24}, {0x55, 0x53, 0x4e}, {0x55, 0x53, 0x53}, {0x55, 0x59, 0x49}, {0x55, 0x59, 0x50}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e}, {0x56, 0x55, 0x56}, {0x57, 0x53, 0x54}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47}, {0x58, 0x41, 0x55}, {0x58, 0x42, 0x41}, {0x58, 0x42, 0x42}, {0x58, 0x42, 0x43}, {0x58, 0x42, 0x44}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52}, {0x58, 0x45, 0x55}, {0x58, 0x46, 0x4f}, {0x58, 0x46, 0x55}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44}, {0x43, 0x46, 0x50, 0x46}, {0x58, 0x50, 0x54}, {0x58, 0x52, 0x45}, {0x58, 0x53, 0x55}, {0x58, 0x54, 0x53}, {0x58, 0x55, 0x41}, {0x58, 0x58, 0x58}, {0x59, 0x44, 0x44}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44}, {0x59, 0x55, 0x4d}, {0x59, 0x55, 0x4e}, {0x59, 0x55, 0x52}, {0x5a, 0x41, 0x4c}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e}, {0x5a, 0x52, 0x5a}, {0x5a, 0x57, 0x44}, {0x5a, 0x57, 0x4c}, {0x5a, 0x57, 0x52}},
		currencyPositiveSuffix: []byte{0xc2, 0xa0},
		currencyNegativeSuffix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x78, 0x69, 0x6e}, {0x66, 0x65, 0x62}, {0x6d, 0x61, 0x72}, {0x61, 0x62, 0x72}, {0x6d, 0x61, 0x79}, {0x78, 0x75, 0x6e}, {0x78, 0x6e, 0x74}, {0x61, 0x67, 0x6f}, {0x73, 0x65, 0x74}, {0x6f, 0x63, 0x68}, {0x70, 0x61, 0x79}, {0x61, 0x76, 0x69}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x58}, {0x46}, {0x4d}, {0x41}, {0x4d}, {0x58}, {0x58}, {0x41}, {0x53}, {0x4f}, {0x50}, {0x41}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x64, 0x65, 0x20, 0x78, 0x69, 0x6e, 0x65, 0x72, 0x75}, {0x64, 0x65, 0x20, 0x66, 0x65, 0x62, 0x72, 0x65, 0x72, 0x75}, {0x64, 0x65, 0x20, 0x6d, 0x61, 0x72, 0x7a, 0x75}, {0x64, 0xe2, 0x80, 0x99, 0x61, 0x62, 0x72, 0x69, 0x6c}, {0x64, 0x65, 0x20, 0x6d, 0x61, 0x79, 0x75}, {0x64, 0x65, 0x20, 0x78, 0x75, 0x6e, 0x75}, {0x64, 0x65, 0x20, 0x78, 0x75, 0x6e, 0x65, 0x74, 0x75}, {0x64, 0xe2, 0x80, 0x99, 0x61, 0x67, 0x6f, 0x73, 0x74, 0x75}, {0x64, 0x65, 0x20, 0x73, 0x65, 0x74, 0x69, 0x65, 0x6d, 0x62, 0x72, 0x65}, {0x64, 0xe2, 0x80, 0x99, 0x6f, 0x63, 0x68, 0x6f, 0x62, 0x72, 0x65}, {0x64, 0x65, 0x20, 0x70, 0x61, 0x79, 0x61, 0x72, 0x65, 0x73}, {0x64, 0xe2, 0x80, 0x99, 0x61, 0x76, 0x69, 0x65, 0x6e, 0x74, 0x75}},
		daysAbbreviated:        [][]uint8{{0x64, 0x6f, 0x6d}, {0x6c, 0x6c, 0x75}, {0x6d, 0x61, 0x72}, {0x6d, 0x69, 0xc3, 0xa9}, {0x78, 0x75, 0x65}, {0x76, 0x69, 0x65}, {0x73, 0xc3, 0xa1, 0x62}},
		daysNarrow:             [][]uint8{{0x44}, {0x4c}, {0x4d}, {0x4d}, {0x58}, {0x56}, {0x53}},
		daysShort:              [][]uint8{{0x64, 0x6f}, {0x6c, 0x6c}, {0x6d, 0x61}, {0x6d, 0x69}, {0x78, 0x75}, {0x76, 0x69}, {0x73, 0xc3, 0xa1}},
		daysWide:               [][]uint8{{0x64, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x75}, {0x6c, 0x6c, 0x75, 0x6e, 0x65, 0x73}, {0x6d, 0x61, 0x72, 0x74, 0x65, 0x73}, {0x6d, 0x69, 0xc3, 0xa9, 0x72, 0x63, 0x6f, 0x6c, 0x65, 0x73}, {0x78, 0x75, 0x65, 0x76, 0x65, 0x73}, {0x76, 0x69, 0x65, 0x6e, 0x72, 0x65, 0x73}, {0x73, 0xc3, 0xa1, 0x62, 0x61, 0x64, 0x75}},
		periodsAbbreviated:     [][]uint8{{0x41, 0x4d}, {0x50, 0x4d}},
		periodsNarrow:          [][]uint8{{0x61}, {0x70}},
		periodsWide:            [][]uint8{{0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x6d, 0x61, 0xc3, 0xb1, 0x61, 0x6e, 0x61}, {0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x74, 0x61, 0x72, 0x64, 0x69}},
		erasAbbreviated:        [][]uint8{{0x61, 0x2e, 0x43, 0x2e}, {0x64, 0x2e, 0x43, 0x2e}},
		erasNarrow:             [][]uint8{{0x61, 0x43}, {0x64, 0x43}},
		erasWide:               [][]uint8{{0x61, 0x2e, 0x43, 0x2e}, {0x64, 0x65, 0x73, 0x70, 0x75, 0xc3, 0xa9, 0x73, 0x20, 0x64, 0x65, 0x20, 0x43, 0x72, 0x69, 0x73, 0x74, 0x75}},
		timezones:              map[string][]uint8{"ACWDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x6f, 0x65, 0x73, 0x74, 0x65}, "ARST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "CDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x6e, 0x6f, 0x72, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x61}, "MYT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x4d, 0x61, 0x6c, 0x61, 0x73, 0x69, 0x61}, "EAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x65, 0x73, 0x74, 0x65}, "ART": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "NZDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x75, 0x65, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61}, "ADT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x41, 0x74, 0x6c, 0xc3, 0xa1, 0x6e, 0x74, 0x69, 0x63, 0x75}, "CST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x6e, 0x6f, 0x72, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x61}, "CHADT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "TMST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc3, 0xa1, 0x6e}, "ECT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x63, 0x75, 0x61, 0x64, 0x6f, 0x72}, "BT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x42, 0x75, 0x74, 0xc3, 0xa1, 0x6e}, "COST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "ACWST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x6f, 0x65, 0x73, 0x74, 0x65}, "ChST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x6d, 0x6f, 0x72, 0x72, 0x6f}, "LHST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "ACST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "WAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x6f, 0x65, 0x73, 0x74, 0x65}, "UYST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0xc3, 0xa1, 0x69}, "MDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x65, 0x73, 0x20, 0x6d, 0x6f, 0x6e, 0x74, 0x61, 0xc3, 0xb1, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x72, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x65, 0x73}, "WIT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x65, 0x73, 0x74, 0x65}, "AWDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x6f, 0x65, 0x73, 0x74, 0x65}, "GFT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x61, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61, 0x20, 0x46, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x61}, "WITA": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "SAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x53, 0x75, 0x64, 0xc3, 0xa1, 0x66, 0x72, 0x69, 0x63, 0x61}, "ACDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "WART": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "WESZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "WIB": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x49, 0x6e, 0x64, 0x6f, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x6f, 0x65, 0x73, 0x74, 0x65}, "EST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x65, 0x73, 0x74, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x75}, "EDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x65, 0x73, 0x74, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x75}, "CHAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x61, 0x74, 0x68, 0x61, 0x6d}, "∅∅∅": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x41, 0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x61, 0x73}, "HAST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x65, 0x73}, "PST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63, 0x75, 0x20, 0x6e, 0x6f, 0x72, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x75}, "HAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "AEST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x65, 0x73, 0x74, 0x65}, "SGT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x53, 0x69, 0x6e, 0x67, 0x61, 0x70, 0x75, 0x72}, "GMT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x20, 0x64, 0x65, 0x20, 0x47, 0x72, 0x65, 0x65, 0x6e, 0x77, 0x69, 0x63, 0x68}, "MST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x65, 0x73, 0x20, 0x6d, 0x6f, 0x6e, 0x74, 0x61, 0xc3, 0xb1, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x72, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x65, 0x73}, "IST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x6c, 0x61, 0x20, 0x49, 0x6e, 0x64, 0x69, 0x61}, "AEDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x65, 0x73, 0x74, 0x65}, "TMT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x54, 0x75, 0x72, 0x6b, 0x6d, 0x65, 0x6e, 0x69, 0x73, 0x74, 0xc3, 0xa1, 0x6e}, "SRT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x53, 0x75, 0x72, 0x69, 0x6e, 0x61, 0x6d}, "HKST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x20, 0xe1, 0xb8, 0xa4, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "WARST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x6f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x72, 0x78, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x61}, "VET": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x56, 0x65, 0x6e, 0x65, 0x7a, 0x75, 0x65, 0x6c, 0x61}, "CLST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "BOT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x42, 0x6f, 0x6c, 0x69, 0x76, 0x69, 0x61}, "WAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x6f, 0x65, 0x73, 0x74, 0x65}, "JDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x20, 0x58, 0x61, 0x70, 0xc3, 0xb3, 0x6e}, "PDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x50, 0x61, 0x63, 0xc3, 0xad, 0x66, 0x69, 0x63, 0x75, 0x20, 0x6e, 0x6f, 0x72, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x75}, "LHDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x6f, 0x72, 0x64, 0x20, 0x48, 0x6f, 0x77, 0x65}, "HNT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x65, 0x77, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x6c, 0x61, 0x6e, 0x64}, "AKDT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}, "COT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x6f, 0x6c, 0x6f, 0x6d, 0x62, 0x69, 0x61}, "CLT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x43, 0x68, 0x69, 0x6c, 0x65}, "JST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x58, 0x61, 0x70, 0xc3, 0xb3, 0x6e}, "WEZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x4f, 0x63, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x61, 0x6c}, "NZST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0x4e, 0x75, 0x65, 0x76, 0x61, 0x20, 0x5a, 0x65, 0x6c, 0x61, 0x6e, 0x64, 0x61}, "MEZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "OESZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x45, 0x73, 0x74, 0x65}, "HKT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x20, 0xe1, 0xb8, 0xa4, 0x6f, 0x6e, 0x67, 0x20, 0x4b, 0x6f, 0x6e, 0x67}, "AWST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x75, 0x73, 0x74, 0x72, 0x61, 0x6c, 0x69, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x6f, 0x65, 0x73, 0x74, 0x65}, "MESZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "OEZ": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x45, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x45, 0x73, 0x74, 0x65}, "HADT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x62, 0x72, 0x61, 0x6e, 0x69, 0x65, 0x67, 0x61, 0x20, 0x64, 0x65, 0x20, 0x48, 0x61, 0x77, 0x61, 0x69, 0x69, 0x2d, 0x41, 0x6c, 0x65, 0x75, 0x74, 0x69, 0x61, 0x6e, 0x65, 0x73}, "AST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x41, 0x74, 0x6c, 0xc3, 0xa1, 0x6e, 0x74, 0x69, 0x63, 0x75}, "GYT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0x65, 0x20, 0x4c, 0x61, 0x20, 0x47, 0x75, 0x79, 0x61, 0x6e, 0x61}, "UYT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0x65, 0x6c, 0x20, 0x55, 0x72, 0x75, 0x67, 0x75, 0xc3, 0xa1, 0x69}, "CAT": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x64, 0xe2, 0x80, 0x99, 0xc3, 0x81, 0x66, 0x72, 0x69, 0x63, 0x61, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c}, "AKST": {0x48, 0x6f, 0x72, 0x61, 0x20, 0x65, 0x73, 0x74, 0xc3, 0xa1, 0x6e, 0x64, 0x61, 0x72, 0x20, 0x64, 0xe2, 0x80, 0x99, 0x41, 0x6c, 0x61, 0x73, 0x6b, 0x61}},
	}
}

// Locale returns the current translators string locale
func (ast *ast) Locale() string {
	return ast.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ast'
func (ast *ast) PluralsCardinal() []locales.PluralRule {
	return ast.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ast'
func (ast *ast) PluralsOrdinal() []locales.PluralRule {
	return ast.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ast'
func (ast *ast) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ast'
func (ast *ast) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ast'
func (ast *ast) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ast *ast) MonthAbbreviated(month time.Month) []byte {
	return ast.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ast *ast) MonthsAbbreviated() [][]byte {
	return ast.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ast *ast) MonthNarrow(month time.Month) []byte {
	return ast.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ast *ast) MonthsNarrow() [][]byte {
	return ast.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ast *ast) MonthWide(month time.Month) []byte {
	return ast.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ast *ast) MonthsWide() [][]byte {
	return ast.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ast *ast) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return ast.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ast *ast) WeekdaysAbbreviated() [][]byte {
	return ast.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ast *ast) WeekdayNarrow(weekday time.Weekday) []byte {
	return ast.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ast *ast) WeekdaysNarrow() [][]byte {
	return ast.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ast *ast) WeekdayShort(weekday time.Weekday) []byte {
	return ast.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ast *ast) WeekdaysShort() [][]byte {
	return ast.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ast *ast) WeekdayWide(weekday time.Weekday) []byte {
	return ast.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ast *ast) WeekdaysWide() [][]byte {
	return ast.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ast' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ast.decimal) + len(ast.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ast.decimal) - 1; j >= 0; j-- {
				b = append(b, ast.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ast.group) - 1; j >= 0; j-- {
					b = append(b, ast.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ast.minus) - 1; j >= 0; j-- {
			b = append(b, ast.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ast' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ast *ast) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(ast.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ast.decimal) - 1; j >= 0; j-- {
				b = append(b, ast.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ast.minus) - 1; j >= 0; j-- {
			b = append(b, ast.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ast.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ast'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ast.currencies[currency]
	l := len(s) + len(ast.decimal) + len(ast.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ast.decimal) - 1; j >= 0; j-- {
				b = append(b, ast.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ast.group) - 1; j >= 0; j-- {
					b = append(b, ast.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ast.minus) - 1; j >= 0; j-- {
			b = append(b, ast.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ast.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ast.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ast'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ast.currencies[currency]
	l := len(s) + len(ast.decimal) + len(ast.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ast.decimal) - 1; j >= 0; j-- {
				b = append(b, ast.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ast.group) - 1; j >= 0; j-- {
					b = append(b, ast.group[j])
				}

				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(ast.minus) - 1; j >= 0; j-- {
			b = append(b, ast.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ast.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ast.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ast.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'ast'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'ast'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ast.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'ast'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ast.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'ast'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, ast.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ast.monthsWide[t.Month()]...)
	b = append(b, []byte{}...)
	b = append(b, []byte{0x27, 0x64, 0x65, 0x27, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'ast'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'ast'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'ast'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'ast'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (ast *ast) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ast.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ast.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
