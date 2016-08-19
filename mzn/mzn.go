package mzn

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type mzn struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	decimal            []byte
	group              []byte
	minus              []byte
	percent            []byte
	perMille           []byte
	timeSeparator      []byte
	inifinity          []byte
	currencies         [][]byte // idx = enum of currency code
	monthsAbbreviated  [][]byte
	monthsNarrow       [][]byte
	monthsWide         [][]byte
	daysAbbreviated    [][]byte
	daysNarrow         [][]byte
	daysShort          [][]byte
	daysWide           [][]byte
	periodsAbbreviated [][]byte
	periodsNarrow      [][]byte
	periodsShort       [][]byte
	periodsWide        [][]byte
	erasAbbreviated    [][]byte
	erasNarrow         [][]byte
	erasWide           [][]byte
	timezones          map[string][]byte
}

// New returns a new instance of translator for the 'mzn' locale
func New() locales.Translator {
	return &mzn{
		locale:            "mzn",
		pluralsCardinal:   nil,
		pluralsOrdinal:    nil,
		decimal:           []byte{},
		group:             []byte{},
		minus:             []byte{},
		percent:           []byte{},
		perMille:          []byte{},
		timeSeparator:     []byte{0x3a},
		currencies:        [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c}, {0x41, 0x4d, 0x44}, {0x41, 0x4e, 0x47}, {0x41, 0x4f, 0x41}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44}, {0x42, 0x44, 0x54}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44}, {0x42, 0x49, 0x46}, {0x42, 0x4d, 0x44}, {0x42, 0x4e, 0x44}, {0x42, 0x4f, 0x42}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x52, 0x24}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44}, {0x42, 0x54, 0x4e}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52}, {0x42, 0x5a, 0x44}, {0x43, 0x41, 0x24}, {0x43, 0x44, 0x46}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0xc2, 0xa5}, {0x43, 0x4f, 0x50}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43}, {0x43, 0x55, 0x50}, {0x43, 0x56, 0x45}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46}, {0x44, 0x4b, 0x4b}, {0x44, 0x4f, 0x50}, {0x44, 0x5a, 0x44}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50}, {0x45, 0x52, 0x4e}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42}, {0xe2, 0x82, 0xac}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50}, {0x46, 0x52, 0x46, 0x20}, {0xc2, 0xa3}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53}, {0x47, 0x49, 0x50}, {0x47, 0x4d, 0x44}, {0x47, 0x4e, 0x46}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44}, {0x48, 0x4b, 0x24}, {0x48, 0x4e, 0x4c}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b}, {0x48, 0x54, 0x47}, {0x48, 0x55, 0x46}, {0x49, 0x44, 0x52}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0xe2, 0x82, 0xaa}, {0xe2, 0x82, 0xb9}, {0x49, 0x51, 0x44}, {0x49, 0x52, 0x52}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44}, {0x4a, 0x4f, 0x44}, {0x4a, 0x50, 0xc2, 0xa5}, {0x4b, 0x45, 0x53}, {0x4b, 0x47, 0x53}, {0x4b, 0x48, 0x52}, {0x4b, 0x4d, 0x46}, {0x4b, 0x50, 0x57}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0xe2, 0x82, 0xa9}, {0x4b, 0x57, 0x44}, {0x4b, 0x59, 0x44}, {0x4b, 0x5a, 0x54}, {0x4c, 0x41, 0x4b}, {0x4c, 0x42, 0x50}, {0x4c, 0x4b, 0x52}, {0x4c, 0x52, 0x44}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44}, {0x4d, 0x41, 0x44}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c}, {0x4d, 0x47, 0x41}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b}, {0x4d, 0x4e, 0x54}, {0x4d, 0x4f, 0x50}, {0x4d, 0x52, 0x4f}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52}, {0x4d, 0x57, 0x4b}, {0x4d, 0x58, 0x24}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e}, {0x4e, 0x41, 0x44}, {0x4e, 0x47, 0x4e}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b}, {0x4e, 0x50, 0x52}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52}, {0x50, 0x41, 0x42}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50}, {0x50, 0x4b, 0x52}, {0x50, 0x4c, 0x4e}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47}, {0x51, 0x41, 0x52}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e}, {0x52, 0x53, 0x44}, {0x52, 0x55, 0x42}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46}, {0x53, 0x41, 0x52}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b}, {0x53, 0x47, 0x44}, {0x53, 0x48, 0x50}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c}, {0x53, 0x4f, 0x53}, {0x53, 0x52, 0x44}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50}, {0x53, 0x54, 0x44}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50}, {0x53, 0x5a, 0x4c}, {0x54, 0x48, 0x42}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54}, {0x54, 0x4e, 0x44}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59}, {0x54, 0x54, 0x44}, {0x4e, 0x54, 0x24}, {0x54, 0x5a, 0x53}, {0x55, 0x41, 0x48}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58}, {0x24}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55}, {0x55, 0x5a, 0x53}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46}, {0xe2, 0x82, 0xab}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x46, 0x43, 0x46, 0x41}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x45, 0x43, 0x24}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x43, 0x46, 0x41}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		monthsAbbreviated: [][]uint8{[]uint8(nil), {0xda, 0x98, 0xd8, 0xa7, 0xd9, 0x86, 0xd9, 0x88, 0xdb, 0x8c, 0xd9, 0x87}, {0xd9, 0x81, 0xd9, 0x88, 0xd8, 0xb1, 0xdb, 0x8c, 0xd9, 0x87}, {0xd9, 0x85, 0xd8, 0xa7, 0xd8, 0xb1, 0xd8, 0xb3}, {0xd8, 0xa2, 0xd9, 0x88, 0xd8, 0xb1, 0xdb, 0x8c, 0xd9, 0x84}, {0xd9, 0x85, 0xd9, 0x87}, {0xda, 0x98, 0xd9, 0x88, 0xd8, 0xa6, 0xd9, 0x86}, {0xda, 0x98, 0xd9, 0x88, 0xd8, 0xa6, 0xdb, 0x8c, 0xd9, 0x87}, {0xd8, 0xa7, 0xd9, 0x88, 0xd8, 0xaa}, {0xd8, 0xb3, 0xd9, 0xbe, 0xd8, 0xaa, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xa7, 0xda, 0xa9, 0xd8, 0xaa, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd9, 0x86, 0xd9, 0x88, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xaf, 0xd8, 0xb3, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}},
		monthsWide:        [][]uint8{[]uint8(nil), {0xda, 0x98, 0xd8, 0xa7, 0xd9, 0x86, 0xd9, 0x88, 0xdb, 0x8c, 0xd9, 0x87}, {0xd9, 0x81, 0xd9, 0x88, 0xd8, 0xb1, 0xdb, 0x8c, 0xd9, 0x87}, {0xd9, 0x85, 0xd8, 0xa7, 0xd8, 0xb1, 0xd8, 0xb3}, {0xd8, 0xa2, 0xd9, 0x88, 0xd8, 0xb1, 0xdb, 0x8c, 0xd9, 0x84}, {0xd9, 0x85, 0xd9, 0x87}, {0xda, 0x98, 0xd9, 0x88, 0xd8, 0xa6, 0xd9, 0x86}, {0xda, 0x98, 0xd9, 0x88, 0xd8, 0xa6, 0xdb, 0x8c, 0xd9, 0x87}, {0xd8, 0xa7, 0xd9, 0x88, 0xd8, 0xaa}, {0xd8, 0xb3, 0xd9, 0xbe, 0xd8, 0xaa, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xa7, 0xda, 0xa9, 0xd8, 0xaa, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd9, 0x86, 0xd9, 0x88, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}, {0xd8, 0xaf, 0xd8, 0xb3, 0xd8, 0xa7, 0xd9, 0x85, 0xd8, 0xa8, 0xd8, 0xb1}},
		erasAbbreviated:   [][]uint8{{0xd9, 0xbe, 0x2e, 0xd9, 0x85}, {0xd9, 0x85, 0x2e}},
		erasNarrow:        [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:          [][]uint8{{0xd9, 0x82, 0xd8, 0xa8, 0xd9, 0x84, 0x20, 0xd9, 0x85, 0xdb, 0x8c, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xaf}, {0xd8, 0xa8, 0xd8, 0xb9, 0xd8, 0xaf, 0x20, 0xd9, 0x85, 0xdb, 0x8c, 0xd9, 0x84, 0xd8, 0xa7, 0xd8, 0xaf}},
		timezones:         map[string][]uint8{"HADT": {0x48, 0x41, 0x44, 0x54}, "PST": {0x50, 0x53, 0x54}, "UYT": {0x55, 0x59, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "ADT": {0x41, 0x44, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "CAT": {0x43, 0x41, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "ART": {0x41, 0x52, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "WIB": {0x57, 0x49, 0x42}, "WAST": {0x57, 0x41, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "BT": {0x42, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "GFT": {0x47, 0x46, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "WAT": {0x57, 0x41, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "WIT": {0x57, 0x49, 0x54}, "SRT": {0x53, 0x52, 0x54}, "EAT": {0x45, 0x41, 0x54}, "SGT": {0x53, 0x47, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "CDT": {0x43, 0x44, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "GYT": {0x47, 0x59, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "WITA": {0x57, 0x49, 0x54, 0x41}, "IST": {0x49, 0x53, 0x54}, "EST": {0x45, 0x53, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "CST": {0x43, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "COT": {0x43, 0x4f, 0x54}, "MST": {0x4d, 0x53, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}},
	}
}

// Locale returns the current translators string locale
func (mzn *mzn) Locale() string {
	return mzn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mzn'
func (mzn *mzn) PluralsCardinal() []locales.PluralRule {
	return mzn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mzn'
func (mzn *mzn) PluralsOrdinal() []locales.PluralRule {
	return mzn.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mzn'
func (mzn *mzn) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mzn'
func (mzn *mzn) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mzn'
func (mzn *mzn) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mzn *mzn) MonthAbbreviated(month time.Month) []byte {
	return mzn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mzn *mzn) MonthsAbbreviated() [][]byte {
	return mzn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mzn *mzn) MonthNarrow(month time.Month) []byte {
	return mzn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mzn *mzn) MonthsNarrow() [][]byte {
	return mzn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mzn *mzn) MonthWide(month time.Month) []byte {
	return mzn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mzn *mzn) MonthsWide() [][]byte {
	return mzn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mzn *mzn) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return mzn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mzn *mzn) WeekdaysAbbreviated() [][]byte {
	return mzn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mzn *mzn) WeekdayNarrow(weekday time.Weekday) []byte {
	return mzn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mzn *mzn) WeekdaysNarrow() [][]byte {
	return mzn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mzn *mzn) WeekdayShort(weekday time.Weekday) []byte {
	return mzn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mzn *mzn) WeekdaysShort() [][]byte {
	return mzn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mzn *mzn) WeekdayWide(weekday time.Weekday) []byte {
	return mzn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mzn *mzn) WeekdaysWide() [][]byte {
	return mzn.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mzn' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mzn' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mzn *mzn) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	return []byte(s)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mzn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mzn.currencies[currency]
	return append(append([]byte{}, symbol...), s...)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mzn'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mzn.currencies[currency]
	return append(append([]byte{}, symbol...), s...)
}

// FmtDateShort returns the short date representation of 't' for 'mzn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'mzn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'mzn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'mzn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'mzn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'mzn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'mzn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'mzn'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (mzn *mzn) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	return b
}
