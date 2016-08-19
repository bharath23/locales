package bo_CN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type bo_CN struct {
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
	currencyPositivePrefix []byte
	currencyNegativePrefix []byte
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

// New returns a new instance of translator for the 'bo_CN' locale
func New() locales.Translator {
	return &bo_CN{
		locale:                 "bo_CN",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		decimal:                []byte{0x2e},
		group:                  []byte{0x2c},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x3a},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa1}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa2}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa3}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa4}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa5}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa6}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa7}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa8}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa9}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa1, 0xe0, 0xbc, 0xa0}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa1, 0xe0, 0xbc, 0xa1}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbc, 0xa1, 0xe0, 0xbc, 0xa2}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x91, 0xe0, 0xbd, 0x84, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94, 0xe0, 0xbd, 0xbc}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0x89, 0xe0, 0xbd, 0xb2, 0xe0, 0xbd, 0xa6, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0xa6, 0xe0, 0xbd, 0xb4, 0xe0, 0xbd, 0x98, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbd, 0x9e, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa3, 0xe0, 0xbe, 0x94, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x91, 0xe0, 0xbe, 0xb2, 0xe0, 0xbd, 0xb4, 0xe0, 0xbd, 0x82, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbd, 0x91, 0xe0, 0xbd, 0xb4, 0xe0, 0xbd, 0x93, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbd, 0xa2, 0xe0, 0xbe, 0x92, 0xe0, 0xbe, 0xb1, 0xe0, 0xbd, 0x91, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x91, 0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0xb4, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbd, 0x85, 0xe0, 0xbd, 0xb4, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbd, 0x85, 0xe0, 0xbd, 0xb4, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0x85, 0xe0, 0xbd, 0xb2, 0xe0, 0xbd, 0x82, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbd, 0x85, 0xe0, 0xbd, 0xb4, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0x89, 0xe0, 0xbd, 0xb2, 0xe0, 0xbd, 0xa6, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94}},
		daysAbbreviated:        [][]uint8{{0xe0, 0xbd, 0x89, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x98, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x98, 0xe0, 0xbd, 0xb2, 0xe0, 0xbd, 0x82, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x91, 0xe0, 0xbd, 0x98, 0xe0, 0xbd, 0xa2, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0xa3, 0xe0, 0xbe, 0xb7, 0xe0, 0xbd, 0x82, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x95, 0xe0, 0xbd, 0xb4, 0xe0, 0xbd, 0xa2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbd, 0xb4, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x94, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa6, 0xe0, 0xbd, 0x84, 0xe0, 0xbd, 0xa6, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0xa4, 0xe0, 0xbd, 0xba, 0xe0, 0xbd, 0x93, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94, 0xe0, 0xbc, 0x8b}},
		daysNarrow:             [][]uint8{{0xe0, 0xbd, 0x89, 0xe0, 0xbd, 0xb2}, {0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3}, {0xe0, 0xbd, 0x98, 0xe0, 0xbd, 0xb2, 0xe0, 0xbd, 0x82}, {0xe0, 0xbd, 0xa3, 0xe0, 0xbe, 0xb7, 0xe0, 0xbd, 0x82}, {0xe0, 0xbd, 0x95, 0xe0, 0xbd, 0xb4, 0xe0, 0xbd, 0xa2}, {0xe0, 0xbd, 0xa6, 0xe0, 0xbd, 0x84, 0xe0, 0xbd, 0xa6}, {0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0xa4, 0xe0, 0xbd, 0xba, 0xe0, 0xbd, 0x93}},
		daysWide:               [][]uint8{{0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0x9f, 0xe0, 0xbd, 0xa0, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x89, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x98, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0x9f, 0xe0, 0xbd, 0xa0, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x9f, 0xe0, 0xbe, 0xb3, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0x9f, 0xe0, 0xbd, 0xa0, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x98, 0xe0, 0xbd, 0xb2, 0xe0, 0xbd, 0x82, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x91, 0xe0, 0xbd, 0x98, 0xe0, 0xbd, 0xa2, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0x9f, 0xe0, 0xbd, 0xa0, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa3, 0xe0, 0xbe, 0xb7, 0xe0, 0xbd, 0x82, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0x9f, 0xe0, 0xbd, 0xa0, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x95, 0xe0, 0xbd, 0xb4, 0xe0, 0xbd, 0xa2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x96, 0xe0, 0xbd, 0xb4, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0x9f, 0xe0, 0xbd, 0xa0, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa6, 0xe0, 0xbd, 0x84, 0xe0, 0xbd, 0xa6, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x82, 0xe0, 0xbd, 0x9f, 0xe0, 0xbd, 0xa0, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0xa4, 0xe0, 0xbd, 0xba, 0xe0, 0xbd, 0x93, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x94, 0xe0, 0xbc, 0x8b}},
		periodsAbbreviated:     [][]uint8{{0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0x94, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x91, 0xe0, 0xbe, 0xb2, 0xe0, 0xbd, 0xbc, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x95, 0xe0, 0xbe, 0xb1, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x91, 0xe0, 0xbe, 0xb2, 0xe0, 0xbd, 0xbc, 0xe0, 0xbc, 0x8b}},
		periodsWide:            [][]uint8{{0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0x94, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x91, 0xe0, 0xbe, 0xb2, 0xe0, 0xbd, 0xbc, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0x95, 0xe0, 0xbe, 0xb1, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x91, 0xe0, 0xbe, 0xb2, 0xe0, 0xbd, 0xbc, 0xe0, 0xbc, 0x8b}},
		erasAbbreviated:        [][]uint8{{0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0xa4, 0xe0, 0xbe, 0xb1, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa3, 0xe0, 0xbd, 0xbc, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0x94, 0xe0, 0xbd, 0xbc, 0xe0, 0xbd, 0x93, 0xe0, 0xbc, 0x8b}, {0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0xa4, 0xe0, 0xbe, 0xb1, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa3, 0xe0, 0xbd, 0xbc, 0xe0, 0xbc, 0x8b}},
		erasNarrow:             [][]uint8{[]uint8(nil), []uint8(nil)},
		erasWide:               [][]uint8{[]uint8(nil), []uint8(nil)},
		timezones:              map[string][]uint8{"HNT": {0x48, 0x4e, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "ART": {0x41, 0x52, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "GYT": {0x47, 0x59, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "PST": {0x50, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "WIT": {0x57, 0x49, 0x54}, "SGT": {0x53, 0x47, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "ECT": {0x45, 0x43, 0x54}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "VET": {0x56, 0x45, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "PDT": {0x50, 0x44, 0x54}, "WAT": {0x57, 0x41, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "HAST": {0x48, 0x41, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "CDT": {0x43, 0x44, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "AST": {0x41, 0x53, 0x54}, "EDT": {0x45, 0x44, 0x54}, "MST": {0x4d, 0x53, 0x54}, "SRT": {0x53, 0x52, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "IST": {0x49, 0x53, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "CAT": {0x43, 0x41, 0x54}, "CST": {0x43, 0x53, 0x54}, "MYT": {0x4d, 0x59, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "AKST": {0x41, 0x4b, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "ADT": {0x41, 0x44, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "EAT": {0x45, 0x41, 0x54}, "BT": {0x42, 0x54}, "GFT": {0x47, 0x46, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "UYT": {0x55, 0x59, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "HADT": {0x48, 0x41, 0x44, 0x54}, "EST": {0x45, 0x53, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "COT": {0x43, 0x4f, 0x54}},
	}
}

// Locale returns the current translators string locale
func (bo *bo_CN) Locale() string {
	return bo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bo_CN'
func (bo *bo_CN) PluralsCardinal() []locales.PluralRule {
	return bo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bo_CN'
func (bo *bo_CN) PluralsOrdinal() []locales.PluralRule {
	return bo.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bo_CN'
func (bo *bo_CN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bo_CN'
func (bo *bo_CN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bo_CN'
func (bo *bo_CN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (bo *bo_CN) MonthAbbreviated(month time.Month) []byte {
	return bo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (bo *bo_CN) MonthsAbbreviated() [][]byte {
	return bo.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (bo *bo_CN) MonthNarrow(month time.Month) []byte {
	return bo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (bo *bo_CN) MonthsNarrow() [][]byte {
	return bo.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (bo *bo_CN) MonthWide(month time.Month) []byte {
	return bo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (bo *bo_CN) MonthsWide() [][]byte {
	return bo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (bo *bo_CN) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return bo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (bo *bo_CN) WeekdaysAbbreviated() [][]byte {
	return bo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (bo *bo_CN) WeekdayNarrow(weekday time.Weekday) []byte {
	return bo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (bo *bo_CN) WeekdaysNarrow() [][]byte {
	return bo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (bo *bo_CN) WeekdayShort(weekday time.Weekday) []byte {
	return bo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (bo *bo_CN) WeekdaysShort() [][]byte {
	return bo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (bo *bo_CN) WeekdayWide(weekday time.Weekday) []byte {
	return bo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (bo *bo_CN) WeekdaysWide() [][]byte {
	return bo.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bo_CN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(bo.decimal) + len(bo.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bo.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(bo.minus) - 1; j >= 0; j-- {
			b = append(b, bo.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bo_CN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bo *bo_CN) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(bo.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bo.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(bo.minus) - 1; j >= 0; j-- {
			b = append(b, bo.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bo.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bo_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bo.currencies[currency]
	l := len(s) + len(bo.decimal) + len(bo.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bo.group[0])
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

	for j := len(bo.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, bo.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(bo.minus) - 1; j >= 0; j-- {
			b = append(b, bo.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bo_CN'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bo.currencies[currency]
	l := len(s) + len(bo.decimal) + len(bo.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, bo.group[0])
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

		for j := len(bo.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, bo.currencyNegativePrefix[j])
		}

		for j := len(bo.minus) - 1; j >= 0; j-- {
			b = append(b, bo.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(bo.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, bo.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'bo_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'bo_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtDateMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20, 0xe0, 0xbd, 0xa3, 0xe0, 0xbd, 0xbc, 0xe0, 0xbd, 0xa0, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b}...)
	b = append(b, bo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0xe0, 0xbd, 0x9a, 0xe0, 0xbd, 0xba, 0xe0, 0xbd, 0xa6, 0xe0, 0xbc, 0x8b}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'bo_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, []byte{0xe0, 0xbd, 0xa6, 0xe0, 0xbe, 0xa4, 0xe0, 0xbe, 0xb1, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0xa3, 0xe0, 0xbd, 0xbc, 0xe0, 0xbc, 0x8b}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bo.monthsWide[t.Month()]...)
	b = append(b, []byte{0xe0, 0xbd, 0xa0, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x9a, 0xe0, 0xbd, 0xba, 0xe0, 0xbd, 0xa6, 0xe0, 0xbc, 0x8b}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'bo_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bo.monthsWide[t.Month()]...)
	b = append(b, []byte{0xe0, 0xbd, 0xa0, 0xe0, 0xbd, 0xb2, 0xe0, 0xbc, 0x8b, 0xe0, 0xbd, 0x9a, 0xe0, 0xbd, 0xba, 0xe0, 0xbd, 0xa6, 0xe0, 0xbc, 0x8b}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, bo.daysWide[t.Weekday()]...)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'bo_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, bo.periodsAbbreviated[0]...)
	} else {
		b = append(b, bo.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'bo_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, bo.periodsAbbreviated[0]...)
	} else {
		b = append(b, bo.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'bo_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, bo.periodsAbbreviated[0]...)
	} else {
		b = append(b, bo.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'bo_CN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (bo *bo_CN) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, bo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, bo.periodsAbbreviated[0]...)
	} else {
		b = append(b, bo.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := bo.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
