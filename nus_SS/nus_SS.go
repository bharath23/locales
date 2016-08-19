package nus_SS

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type nus_SS struct {
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

// New returns a new instance of translator for the 'nus_SS' locale
func New() locales.Translator {
	return &nus_SS{
		locale:                 "nus_SS",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyNegativePrefix: []byte{0x28},
		currencyNegativeSuffix: []byte{0x29},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0x54, 0x69, 0x6f, 0x70}, {0x50, 0xc9, 0x9b, 0x74}, {0x44, 0x75, 0xc9, 0x94, 0xcc, 0xb1, 0xc9, 0x94, 0xcc, 0xb1}, {0x47, 0x75, 0x61, 0x6b}, {0x44, 0x75, 0xc3, 0xa4}, {0x4b, 0x6f, 0x72}, {0x50, 0x61, 0x79}, {0x54, 0x68, 0x6f, 0x6f}, {0x54, 0xc9, 0x9b, 0xc9, 0x9b}, {0x4c, 0x61, 0x61}, {0x4b, 0x75, 0x72}, {0x54, 0x69, 0x64}},
		monthsNarrow:           [][]uint8{[]uint8(nil), {0x54}, {0x50}, {0x44}, {0x47}, {0x44}, {0x4b}, {0x50}, {0x54}, {0x54}, {0x4c}, {0x4b}, {0x54}},
		monthsWide:             [][]uint8{[]uint8(nil), {0x54, 0x69, 0x6f, 0x70, 0x20, 0x74, 0x68, 0x61, 0x72, 0x20, 0x70, 0xc9, 0x9b, 0x74}, {0x50, 0xc9, 0x9b, 0x74}, {0x44, 0x75, 0xc9, 0x94, 0xcc, 0xb1, 0xc9, 0x94, 0xcc, 0xb1, 0xc5, 0x8b}, {0x47, 0x75, 0x61, 0x6b}, {0x44, 0x75, 0xc3, 0xa4, 0x74}, {0x4b, 0x6f, 0x72, 0x6e, 0x79, 0x6f, 0x6f, 0x74}, {0x50, 0x61, 0x79, 0x20, 0x79, 0x69, 0x65, 0xcc, 0xb1, 0x74, 0x6e, 0x69}, {0x54, 0x68, 0x6f, 0xcc, 0xb1, 0x6f, 0xcc, 0xb1, 0x72}, {0x54, 0xc9, 0x9b, 0xc9, 0x9b, 0x72}, {0x4c, 0x61, 0x61, 0x74, 0x68}, {0x4b, 0x75, 0x72}, {0x54, 0x69, 0x6f, 0xcc, 0xb1, 0x70, 0x20, 0x69, 0x6e, 0x20, 0x64, 0x69, 0xcc, 0xb1, 0x69, 0xcc, 0xb1, 0x74}},
		daysAbbreviated:        [][]uint8{{0x43, 0xc3, 0xa4, 0xc5, 0x8b}, {0x4a, 0x69, 0x65, 0x63}, {0x52, 0xc9, 0x9b, 0x77}, {0x44, 0x69, 0xc9, 0x94, 0xcc, 0xb1, 0x6b}, {0xc5, 0x8a, 0x75, 0x61, 0x61, 0x6e}, {0x44, 0x68, 0x69, 0x65, 0x65, 0x63}, {0x42, 0xc3, 0xa4, 0x6b, 0xc9, 0x9b, 0x6c}},
		daysNarrow:             [][]uint8{{0x43}, {0x4a}, {0x52}, {0x44}, {0xc5, 0x8a}, {0x44}, {0x42}},
		daysWide:               [][]uint8{{0x43, 0xc3, 0xa4, 0xc5, 0x8b, 0x20, 0x6b, 0x75, 0xc9, 0x94, 0x74, 0x68}, {0x4a, 0x69, 0x65, 0x63, 0x20, 0x6c, 0x61, 0xcc, 0xb1, 0x74}, {0x52, 0xc9, 0x9b, 0x77, 0x20, 0x6c, 0xc3, 0xa4, 0x74, 0x6e, 0x69}, {0x44, 0x69, 0xc9, 0x94, 0xcc, 0xb1, 0x6b, 0x20, 0x6c, 0xc3, 0xa4, 0x74, 0x6e, 0x69}, {0xc5, 0x8a, 0x75, 0x61, 0x61, 0x6e, 0x20, 0x6c, 0xc3, 0xa4, 0x74, 0x6e, 0x69}, {0x44, 0x68, 0x69, 0x65, 0x65, 0x63, 0x20, 0x6c, 0xc3, 0xa4, 0x74, 0x6e, 0x69}, {0x42, 0xc3, 0xa4, 0x6b, 0xc9, 0x9b, 0x6c, 0x20, 0x6c, 0xc3, 0xa4, 0x74, 0x6e, 0x69}},
		periodsAbbreviated:     [][]uint8{{0x52, 0x57}, {0x54, 0xc5, 0x8a}},
		periodsWide:            [][]uint8{{0x52, 0x57}, {0x54, 0xc5, 0x8a}},
		erasAbbreviated:        [][]uint8{{0x41, 0x59}, {0xc6, 0x90, 0x59}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{{0x41, 0x20, 0x6b, 0x61, 0xcc, 0xb1, 0x6e, 0x20, 0x59, 0x65, 0x63, 0x75, 0x20, 0x6e, 0x69, 0x20, 0x64, 0x61, 0x70}, {0xc6, 0x90, 0x20, 0x63, 0x61, 0x20, 0x59, 0x65, 0x63, 0x75, 0x20, 0x64, 0x61, 0x70}},
		timezones:              map[string][]uint8{"GYT": {0x47, 0x59, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "PDT": {0x50, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "HAT": {0x48, 0x41, 0x54}, "IST": {0x49, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "ChST": {0x43, 0x68, 0x53, 0x54}, "AST": {0x41, 0x53, 0x54}, "EST": {0x45, 0x53, 0x54}, "WAT": {0x57, 0x41, 0x54}, "WIT": {0x57, 0x49, 0x54}, "CST": {0x43, 0x53, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "GMT": {0x47, 0x4d, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "SGT": {0x53, 0x47, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "EDT": {0x45, 0x44, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "BT": {0x42, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "CAT": {0x43, 0x41, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "EAT": {0x45, 0x41, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "HNT": {0x48, 0x4e, 0x54}, "UYT": {0x55, 0x59, 0x54}, "SRT": {0x53, 0x52, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "ART": {0x41, 0x52, 0x54}, "ADT": {0x41, 0x44, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "VET": {0x56, 0x45, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "MST": {0x4d, 0x53, 0x54}, "GFT": {0x47, 0x46, 0x54}, "WIB": {0x57, 0x49, 0x42}, "MDT": {0x4d, 0x44, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "OEZ": {0x4f, 0x45, 0x5a}, "HADT": {0x48, 0x41, 0x44, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "HAST": {0x48, 0x41, 0x53, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "CDT": {0x43, 0x44, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "COT": {0x43, 0x4f, 0x54}},
	}
}

// Locale returns the current translators string locale
func (nus *nus_SS) Locale() string {
	return nus.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nus_SS'
func (nus *nus_SS) PluralsCardinal() []locales.PluralRule {
	return nus.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nus_SS'
func (nus *nus_SS) PluralsOrdinal() []locales.PluralRule {
	return nus.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nus_SS'
func (nus *nus_SS) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nus_SS'
func (nus *nus_SS) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nus_SS'
func (nus *nus_SS) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (nus *nus_SS) MonthAbbreviated(month time.Month) []byte {
	return nus.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (nus *nus_SS) MonthsAbbreviated() [][]byte {
	return nus.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (nus *nus_SS) MonthNarrow(month time.Month) []byte {
	return nus.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (nus *nus_SS) MonthsNarrow() [][]byte {
	return nus.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (nus *nus_SS) MonthWide(month time.Month) []byte {
	return nus.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (nus *nus_SS) MonthsWide() [][]byte {
	return nus.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (nus *nus_SS) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return nus.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (nus *nus_SS) WeekdaysAbbreviated() [][]byte {
	return nus.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (nus *nus_SS) WeekdayNarrow(weekday time.Weekday) []byte {
	return nus.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (nus *nus_SS) WeekdaysNarrow() [][]byte {
	return nus.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (nus *nus_SS) WeekdayShort(weekday time.Weekday) []byte {
	return nus.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (nus *nus_SS) WeekdaysShort() [][]byte {
	return nus.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (nus *nus_SS) WeekdayWide(weekday time.Weekday) []byte {
	return nus.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (nus *nus_SS) WeekdaysWide() [][]byte {
	return nus.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nus_SS' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(nus.decimal) + len(nus.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nus.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nus.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nus.minus) - 1; j >= 0; j-- {
			b = append(b, nus.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nus_SS' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nus *nus_SS) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(nus.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nus.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nus.minus) - 1; j >= 0; j-- {
			b = append(b, nus.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, nus.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nus_SS'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nus.currencies[currency]
	l := len(s) + len(nus.decimal) + len(nus.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nus.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nus.group[0])
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
		for j := len(nus.minus) - 1; j >= 0; j-- {
			b = append(b, nus.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nus.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nus_SS'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nus.currencies[currency]
	l := len(s) + len(nus.decimal) + len(nus.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nus.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nus.group[0])
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

		for j := len(nus.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, nus.currencyNegativePrefix[j])
		}

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
			b = append(b, nus.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, nus.currencyNegativeSuffix...)
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'nus_SS'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'nus_SS'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nus.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'nus_SS'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nus.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'nus_SS'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, nus.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nus.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'nus_SS'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, nus.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, nus.periodsAbbreviated[0]...)
	} else {
		b = append(b, nus.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'nus_SS'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, nus.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nus.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, nus.periodsAbbreviated[0]...)
	} else {
		b = append(b, nus.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'nus_SS'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, nus.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nus.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, nus.periodsAbbreviated[0]...)
	} else {
		b = append(b, nus.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'nus_SS'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (nus *nus_SS) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	tz, _ := t.Zone()

	if btz, ok := nus.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, nus.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nus.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, nus.periodsAbbreviated[0]...)
	} else {
		b = append(b, nus.periodsAbbreviated[1]...)
	}

	return b
}
