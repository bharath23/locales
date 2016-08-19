package as_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type as_IN struct {
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

// New returns a new instance of translator for the 'as_IN' locale
func New() locales.Translator {
	return &as_IN{
		locale:                 "as_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 5, 6},
		decimal:                []byte{},
		group:                  []byte{},
		minus:                  []byte{},
		percent:                []byte{},
		perMille:               []byte{},
		timeSeparator:          []byte{0x2e},
		currencies:             [][]uint8{{0x41, 0x44, 0x50, 0x20}, {0x41, 0x45, 0x44, 0x20}, {0x41, 0x46, 0x41, 0x20}, {0x41, 0x46, 0x4e, 0x20}, {0x41, 0x4c, 0x4b, 0x20}, {0x41, 0x4c, 0x4c, 0x20}, {0x41, 0x4d, 0x44, 0x20}, {0x41, 0x4e, 0x47, 0x20}, {0x41, 0x4f, 0x41, 0x20}, {0x41, 0x4f, 0x4b, 0x20}, {0x41, 0x4f, 0x4e, 0x20}, {0x41, 0x4f, 0x52, 0x20}, {0x41, 0x52, 0x41, 0x20}, {0x41, 0x52, 0x4c, 0x20}, {0x41, 0x52, 0x4d, 0x20}, {0x41, 0x52, 0x50, 0x20}, {0x41, 0x52, 0x53, 0x20}, {0x41, 0x54, 0x53, 0x20}, {0x41, 0x55, 0x44, 0x20}, {0x41, 0x57, 0x47, 0x20}, {0x41, 0x5a, 0x4d, 0x20}, {0x41, 0x5a, 0x4e, 0x20}, {0x42, 0x41, 0x44, 0x20}, {0x42, 0x41, 0x4d, 0x20}, {0x42, 0x41, 0x4e, 0x20}, {0x42, 0x42, 0x44, 0x20}, {0x42, 0x44, 0x54, 0x20}, {0x42, 0x45, 0x43, 0x20}, {0x42, 0x45, 0x46, 0x20}, {0x42, 0x45, 0x4c, 0x20}, {0x42, 0x47, 0x4c, 0x20}, {0x42, 0x47, 0x4d, 0x20}, {0x42, 0x47, 0x4e, 0x20}, {0x42, 0x47, 0x4f, 0x20}, {0x42, 0x48, 0x44, 0x20}, {0x42, 0x49, 0x46, 0x20}, {0x42, 0x4d, 0x44, 0x20}, {0x42, 0x4e, 0x44, 0x20}, {0x42, 0x4f, 0x42, 0x20}, {0x42, 0x4f, 0x4c, 0x20}, {0x42, 0x4f, 0x50, 0x20}, {0x42, 0x4f, 0x56, 0x20}, {0x42, 0x52, 0x42, 0x20}, {0x42, 0x52, 0x43, 0x20}, {0x42, 0x52, 0x45, 0x20}, {0x42, 0x52, 0x4c, 0x20}, {0x42, 0x52, 0x4e, 0x20}, {0x42, 0x52, 0x52, 0x20}, {0x42, 0x52, 0x5a, 0x20}, {0x42, 0x53, 0x44, 0x20}, {0x42, 0x54, 0x4e, 0x20}, {0x42, 0x55, 0x4b, 0x20}, {0x42, 0x57, 0x50, 0x20}, {0x42, 0x59, 0x42, 0x20}, {0x42, 0x59, 0x52, 0x20}, {0x42, 0x5a, 0x44, 0x20}, {0x43, 0x41, 0x44, 0x20}, {0x43, 0x44, 0x46, 0x20}, {0x43, 0x48, 0x45, 0x20}, {0x43, 0x48, 0x46, 0x20}, {0x43, 0x48, 0x57, 0x20}, {0x43, 0x4c, 0x45, 0x20}, {0x43, 0x4c, 0x46, 0x20}, {0x43, 0x4c, 0x50, 0x20}, {0x43, 0x4e, 0x58, 0x20}, {0x43, 0x4e, 0x59, 0x20}, {0x43, 0x4f, 0x50, 0x20}, {0x43, 0x4f, 0x55, 0x20}, {0x43, 0x52, 0x43, 0x20}, {0x43, 0x53, 0x44, 0x20}, {0x43, 0x53, 0x4b, 0x20}, {0x43, 0x55, 0x43, 0x20}, {0x43, 0x55, 0x50, 0x20}, {0x43, 0x56, 0x45, 0x20}, {0x43, 0x59, 0x50, 0x20}, {0x43, 0x5a, 0x4b, 0x20}, {0x44, 0x44, 0x4d, 0x20}, {0x44, 0x45, 0x4d, 0x20}, {0x44, 0x4a, 0x46, 0x20}, {0x44, 0x4b, 0x4b, 0x20}, {0x44, 0x4f, 0x50, 0x20}, {0x44, 0x5a, 0x44, 0x20}, {0x45, 0x43, 0x53, 0x20}, {0x45, 0x43, 0x56, 0x20}, {0x45, 0x45, 0x4b, 0x20}, {0x45, 0x47, 0x50, 0x20}, {0x45, 0x52, 0x4e, 0x20}, {0x45, 0x53, 0x41, 0x20}, {0x45, 0x53, 0x42, 0x20}, {0x45, 0x53, 0x50, 0x20}, {0x45, 0x54, 0x42, 0x20}, {0x45, 0x55, 0x52, 0x20}, {0x46, 0x49, 0x4d, 0x20}, {0x46, 0x4a, 0x44, 0x20}, {0x46, 0x4b, 0x50, 0x20}, {0x46, 0x52, 0x46, 0x20}, {0x47, 0x42, 0x50, 0x20}, {0x47, 0x45, 0x4b, 0x20}, {0x47, 0x45, 0x4c, 0x20}, {0x47, 0x48, 0x43, 0x20}, {0x47, 0x48, 0x53, 0x20}, {0x47, 0x49, 0x50, 0x20}, {0x47, 0x4d, 0x44, 0x20}, {0x47, 0x4e, 0x46, 0x20}, {0x47, 0x4e, 0x53, 0x20}, {0x47, 0x51, 0x45, 0x20}, {0x47, 0x52, 0x44, 0x20}, {0x47, 0x54, 0x51, 0x20}, {0x47, 0x57, 0x45, 0x20}, {0x47, 0x57, 0x50, 0x20}, {0x47, 0x59, 0x44, 0x20}, {0x48, 0x4b, 0x44, 0x20}, {0x48, 0x4e, 0x4c, 0x20}, {0x48, 0x52, 0x44, 0x20}, {0x48, 0x52, 0x4b, 0x20}, {0x48, 0x54, 0x47, 0x20}, {0x48, 0x55, 0x46, 0x20}, {0x49, 0x44, 0x52, 0x20}, {0x49, 0x45, 0x50, 0x20}, {0x49, 0x4c, 0x50, 0x20}, {0x49, 0x4c, 0x52, 0x20}, {0x49, 0x4c, 0x53, 0x20}, {0x49, 0x4e, 0x52, 0x20}, {0x49, 0x51, 0x44, 0x20}, {0x49, 0x52, 0x52, 0x20}, {0x49, 0x53, 0x4a, 0x20}, {0x49, 0x53, 0x4b, 0x20}, {0x49, 0x54, 0x4c, 0x20}, {0x4a, 0x4d, 0x44, 0x20}, {0x4a, 0x4f, 0x44, 0x20}, {0x4a, 0x50, 0x59, 0x20}, {0x4b, 0x45, 0x53, 0x20}, {0x4b, 0x47, 0x53, 0x20}, {0x4b, 0x48, 0x52, 0x20}, {0x4b, 0x4d, 0x46, 0x20}, {0x4b, 0x50, 0x57, 0x20}, {0x4b, 0x52, 0x48, 0x20}, {0x4b, 0x52, 0x4f, 0x20}, {0x4b, 0x52, 0x57, 0x20}, {0x4b, 0x57, 0x44, 0x20}, {0x4b, 0x59, 0x44, 0x20}, {0x4b, 0x5a, 0x54, 0x20}, {0x4c, 0x41, 0x4b, 0x20}, {0x4c, 0x42, 0x50, 0x20}, {0x4c, 0x4b, 0x52, 0x20}, {0x4c, 0x52, 0x44, 0x20}, {0x4c, 0x53, 0x4c, 0x20}, {0x4c, 0x54, 0x4c, 0x20}, {0x4c, 0x54, 0x54, 0x20}, {0x4c, 0x55, 0x43, 0x20}, {0x4c, 0x55, 0x46, 0x20}, {0x4c, 0x55, 0x4c, 0x20}, {0x4c, 0x56, 0x4c, 0x20}, {0x4c, 0x56, 0x52, 0x20}, {0x4c, 0x59, 0x44, 0x20}, {0x4d, 0x41, 0x44, 0x20}, {0x4d, 0x41, 0x46, 0x20}, {0x4d, 0x43, 0x46, 0x20}, {0x4d, 0x44, 0x43, 0x20}, {0x4d, 0x44, 0x4c, 0x20}, {0x4d, 0x47, 0x41, 0x20}, {0x4d, 0x47, 0x46, 0x20}, {0x4d, 0x4b, 0x44, 0x20}, {0x4d, 0x4b, 0x4e, 0x20}, {0x4d, 0x4c, 0x46, 0x20}, {0x4d, 0x4d, 0x4b, 0x20}, {0x4d, 0x4e, 0x54, 0x20}, {0x4d, 0x4f, 0x50, 0x20}, {0x4d, 0x52, 0x4f, 0x20}, {0x4d, 0x54, 0x4c, 0x20}, {0x4d, 0x54, 0x50, 0x20}, {0x4d, 0x55, 0x52, 0x20}, {0x4d, 0x56, 0x50, 0x20}, {0x4d, 0x56, 0x52, 0x20}, {0x4d, 0x57, 0x4b, 0x20}, {0x4d, 0x58, 0x4e, 0x20}, {0x4d, 0x58, 0x50, 0x20}, {0x4d, 0x58, 0x56, 0x20}, {0x4d, 0x59, 0x52, 0x20}, {0x4d, 0x5a, 0x45, 0x20}, {0x4d, 0x5a, 0x4d, 0x20}, {0x4d, 0x5a, 0x4e, 0x20}, {0x4e, 0x41, 0x44, 0x20}, {0x4e, 0x47, 0x4e, 0x20}, {0x4e, 0x49, 0x43, 0x20}, {0x4e, 0x49, 0x4f, 0x20}, {0x4e, 0x4c, 0x47, 0x20}, {0x4e, 0x4f, 0x4b, 0x20}, {0x4e, 0x50, 0x52, 0x20}, {0x4e, 0x5a, 0x44, 0x20}, {0x4f, 0x4d, 0x52, 0x20}, {0x50, 0x41, 0x42, 0x20}, {0x50, 0x45, 0x49, 0x20}, {0x50, 0x45, 0x4e, 0x20}, {0x50, 0x45, 0x53, 0x20}, {0x50, 0x47, 0x4b, 0x20}, {0x50, 0x48, 0x50, 0x20}, {0x50, 0x4b, 0x52, 0x20}, {0x50, 0x4c, 0x4e, 0x20}, {0x50, 0x4c, 0x5a, 0x20}, {0x50, 0x54, 0x45, 0x20}, {0x50, 0x59, 0x47, 0x20}, {0x51, 0x41, 0x52, 0x20}, {0x52, 0x48, 0x44, 0x20}, {0x52, 0x4f, 0x4c, 0x20}, {0x52, 0x4f, 0x4e, 0x20}, {0x52, 0x53, 0x44, 0x20}, {0x52, 0x55, 0x42, 0x20}, {0x52, 0x55, 0x52, 0x20}, {0x52, 0x57, 0x46, 0x20}, {0x53, 0x41, 0x52, 0x20}, {0x53, 0x42, 0x44, 0x20}, {0x53, 0x43, 0x52, 0x20}, {0x53, 0x44, 0x44, 0x20}, {0x53, 0x44, 0x47, 0x20}, {0x53, 0x44, 0x50, 0x20}, {0x53, 0x45, 0x4b, 0x20}, {0x53, 0x47, 0x44, 0x20}, {0x53, 0x48, 0x50, 0x20}, {0x53, 0x49, 0x54, 0x20}, {0x53, 0x4b, 0x4b, 0x20}, {0x53, 0x4c, 0x4c, 0x20}, {0x53, 0x4f, 0x53, 0x20}, {0x53, 0x52, 0x44, 0x20}, {0x53, 0x52, 0x47, 0x20}, {0x53, 0x53, 0x50, 0x20}, {0x53, 0x54, 0x44, 0x20}, {0x53, 0x55, 0x52, 0x20}, {0x53, 0x56, 0x43, 0x20}, {0x53, 0x59, 0x50, 0x20}, {0x53, 0x5a, 0x4c, 0x20}, {0x54, 0x48, 0x42, 0x20}, {0x54, 0x4a, 0x52, 0x20}, {0x54, 0x4a, 0x53, 0x20}, {0x54, 0x4d, 0x4d, 0x20}, {0x54, 0x4d, 0x54, 0x20}, {0x54, 0x4e, 0x44, 0x20}, {0x54, 0x4f, 0x50, 0x20}, {0x54, 0x50, 0x45, 0x20}, {0x54, 0x52, 0x4c, 0x20}, {0x54, 0x52, 0x59, 0x20}, {0x54, 0x54, 0x44, 0x20}, {0x54, 0x57, 0x44, 0x20}, {0x54, 0x5a, 0x53, 0x20}, {0x55, 0x41, 0x48, 0x20}, {0x55, 0x41, 0x4b, 0x20}, {0x55, 0x47, 0x53, 0x20}, {0x55, 0x47, 0x58, 0x20}, {0x55, 0x53, 0x44, 0x20}, {0x55, 0x53, 0x4e, 0x20}, {0x55, 0x53, 0x53, 0x20}, {0x55, 0x59, 0x49, 0x20}, {0x55, 0x59, 0x50, 0x20}, {0x55, 0x59, 0x55, 0x20}, {0x55, 0x5a, 0x53, 0x20}, {0x56, 0x45, 0x42, 0x20}, {0x56, 0x45, 0x46, 0x20}, {0x56, 0x4e, 0x44, 0x20}, {0x56, 0x4e, 0x4e, 0x20}, {0x56, 0x55, 0x56, 0x20}, {0x57, 0x53, 0x54, 0x20}, {0x58, 0x41, 0x46, 0x20}, {0x58, 0x41, 0x47, 0x20}, {0x58, 0x41, 0x55, 0x20}, {0x58, 0x42, 0x41, 0x20}, {0x58, 0x42, 0x42, 0x20}, {0x58, 0x42, 0x43, 0x20}, {0x58, 0x42, 0x44, 0x20}, {0x58, 0x43, 0x44, 0x20}, {0x58, 0x44, 0x52, 0x20}, {0x58, 0x45, 0x55, 0x20}, {0x58, 0x46, 0x4f, 0x20}, {0x58, 0x46, 0x55, 0x20}, {0x58, 0x4f, 0x46, 0x20}, {0x58, 0x50, 0x44, 0x20}, {0x58, 0x50, 0x46, 0x20}, {0x58, 0x50, 0x54, 0x20}, {0x58, 0x52, 0x45, 0x20}, {0x58, 0x53, 0x55, 0x20}, {0x58, 0x54, 0x53, 0x20}, {0x58, 0x55, 0x41, 0x20}, {0x58, 0x58, 0x58, 0x20}, {0x59, 0x44, 0x44, 0x20}, {0x59, 0x45, 0x52, 0x20}, {0x59, 0x55, 0x44, 0x20}, {0x59, 0x55, 0x4d, 0x20}, {0x59, 0x55, 0x4e, 0x20}, {0x59, 0x55, 0x52, 0x20}, {0x5a, 0x41, 0x4c, 0x20}, {0x5a, 0x41, 0x52, 0x20}, {0x5a, 0x4d, 0x4b, 0x20}, {0x5a, 0x4d, 0x57, 0x20}, {0x5a, 0x52, 0x4e, 0x20}, {0x5a, 0x52, 0x5a, 0x20}, {0x5a, 0x57, 0x44, 0x20}, {0x5a, 0x57, 0x4c, 0x20}, {0x5a, 0x57, 0x52, 0x20}},
		currencyPositivePrefix: []byte{0xc2, 0xa0},
		currencyNegativePrefix: []byte{0xc2, 0xa0},
		monthsAbbreviated:      [][]uint8{[]uint8(nil), {0xe0, 0xa6, 0x9c, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xa8, 0xe0, 0xa7, 0x81}, {0xe0, 0xa6, 0xab, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x8d, 0xe0, 0xa7, 0xb0, 0xe0, 0xa7, 0x81}, {0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9a}, {0xe0, 0xa6, 0x8f, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x8d, 0xe0, 0xa7, 0xb0, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xb2}, {0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x87}, {0xe0, 0xa6, 0x9c, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xa8}, {0xe0, 0xa6, 0x9c, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xb2, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0x87}, {0xe0, 0xa6, 0x86, 0xe0, 0xa6, 0x97}, {0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f}, {0xe0, 0xa6, 0x85, 0xe0, 0xa6, 0x95, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f, 0xe0, 0xa7, 0x8b}, {0xe0, 0xa6, 0xa8, 0xe0, 0xa6, 0xad, 0xe0, 0xa7, 0x87}, {0xe0, 0xa6, 0xa1, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x87}},
		monthsWide:             [][]uint8{[]uint8(nil), {0xe0, 0xa6, 0x9c, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xa8, 0xe0, 0xa7, 0x81, 0xe0, 0xa7, 0xb1, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0, 0xe0, 0xa7, 0x80}, {0xe0, 0xa6, 0xab, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x8d, 0xe0, 0xa7, 0xb0, 0xe0, 0xa7, 0x81, 0xe0, 0xa7, 0xb1, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0, 0xe0, 0xa7, 0x80}, {0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9a}, {0xe0, 0xa6, 0x8f, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x8d, 0xe0, 0xa7, 0xb0, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xb2}, {0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x87}, {0xe0, 0xa6, 0x9c, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xa8}, {0xe0, 0xa6, 0x9c, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xb2, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0x87}, {0xe0, 0xa6, 0x86, 0xe0, 0xa6, 0x97, 0xe0, 0xa6, 0xb7, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f}, {0xe0, 0xa6, 0x9b, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xa4, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0xb0}, {0xe0, 0xa6, 0x85, 0xe0, 0xa6, 0x95, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x9f, 0xe0, 0xa7, 0x8b, 0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0xb0}, {0xe0, 0xa6, 0xa8, 0xe0, 0xa7, 0xb1, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0xb0}, {0xe0, 0xa6, 0xa1, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0x9a, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0xae, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0xb0}},
		daysAbbreviated:        [][]uint8{{0xe0, 0xa7, 0xb0, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbf}, {0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8b, 0xe0, 0xa6, 0xae}, {0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0x99, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x97, 0xe0, 0xa6, 0xb2}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xa7}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x83, 0xe0, 0xa6, 0xb9, 0xe0, 0xa6, 0xb7, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xaa, 0xe0, 0xa6, 0xa4, 0xe0, 0xa6, 0xbf}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0x95, 0xe0, 0xa7, 0x8d, 0xe0, 0xa7, 0xb0}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa6, 0xa8, 0xe0, 0xa6, 0xbf}},
		daysWide:               [][]uint8{{0xe0, 0xa6, 0xa6, 0xe0, 0xa7, 0x87, 0xe0, 0xa6, 0x93, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0}, {0xe0, 0xa6, 0xb8, 0xe0, 0xa7, 0x8b, 0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0}, {0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0x99, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0x97, 0xe0, 0xa6, 0xb2, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0xa7, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0}, {0xe0, 0xa6, 0xac, 0xe0, 0xa7, 0x83, 0xe0, 0xa6, 0xb9, 0xe0, 0xa6, 0xb7, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xaa, 0xe0, 0xa6, 0xa4, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa7, 0x81, 0xe0, 0xa6, 0x95, 0xe0, 0xa7, 0x8d, 0xe0, 0xa7, 0xb0, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0}, {0xe0, 0xa6, 0xb6, 0xe0, 0xa6, 0xa8, 0xe0, 0xa6, 0xbf, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0}},
		periodsAbbreviated:     [][]uint8{{0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x82, 0xe0, 0xa7, 0xb0, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb9, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xa3}, {0xe0, 0xa6, 0x85, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0xb0, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb9, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xa3}},
		periodsWide:            [][]uint8{{0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0x82, 0xe0, 0xa7, 0xb0, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xac, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb9, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xa3}, {0xe0, 0xa6, 0x85, 0xe0, 0xa6, 0xaa, 0xe0, 0xa7, 0xb0, 0xe0, 0xa6, 0xbe, 0xe0, 0xa6, 0xb9, 0xe0, 0xa7, 0x8d, 0xe0, 0xa6, 0xa3}},
		timezones:              map[string][]uint8{"AKST": {0x41, 0x4b, 0x53, 0x54}, "AEDT": {0x41, 0x45, 0x44, 0x54}, "HAST": {0x48, 0x41, 0x53, 0x54}, "PST": {0x50, 0x53, 0x54}, "HNT": {0x48, 0x4e, 0x54}, "BOT": {0x42, 0x4f, 0x54}, "AWDT": {0x41, 0x57, 0x44, 0x54}, "WITA": {0x57, 0x49, 0x54, 0x41}, "EDT": {0x45, 0x44, 0x54}, "ACST": {0x41, 0x43, 0x53, 0x54}, "SAST": {0x53, 0x41, 0x53, 0x54}, "HAT": {0x48, 0x41, 0x54}, "NZDT": {0x4e, 0x5a, 0x44, 0x54}, "ECT": {0x45, 0x43, 0x54}, "MESZ": {0x4d, 0x45, 0x53, 0x5a}, "ART": {0x41, 0x52, 0x54}, "OESZ": {0x4f, 0x45, 0x53, 0x5a}, "CHADT": {0x43, 0x48, 0x41, 0x44, 0x54}, "EAT": {0x45, 0x41, 0x54}, "IST": {0xe0, 0xa6, 0xad, 0xe0, 0xa6, 0xbe, 0xe0, 0xa7, 0xb0, 0xe0, 0xa6, 0xa4, 0xe0, 0xa7, 0x80, 0xe0, 0xa6, 0xaf, 0xe0, 0xa6, 0xbc, 0x20, 0xe0, 0xa6, 0xb8, 0xe0, 0xa6, 0xae, 0xe0, 0xa6, 0xaf, 0xe0, 0xa6, 0xbc}, "ACWST": {0x41, 0x43, 0x57, 0x53, 0x54}, "CAT": {0x43, 0x41, 0x54}, "ACDT": {0x41, 0x43, 0x44, 0x54}, "COST": {0x43, 0x4f, 0x53, 0x54}, "PDT": {0x50, 0x44, 0x54}, "TMST": {0x54, 0x4d, 0x53, 0x54}, "CDT": {0x43, 0x44, 0x54}, "GMT": {0x47, 0x4d, 0x54}, "AEST": {0x41, 0x45, 0x53, 0x54}, "ACWDT": {0x41, 0x43, 0x57, 0x44, 0x54}, "GYT": {0x47, 0x59, 0x54}, "JDT": {0x4a, 0x44, 0x54}, "WAT": {0x57, 0x41, 0x54}, "MST": {0x4d, 0x53, 0x54}, "AWST": {0x41, 0x57, 0x53, 0x54}, "COT": {0x43, 0x4f, 0x54}, "MEZ": {0x4d, 0x45, 0x5a}, "SGT": {0x53, 0x47, 0x54}, "BT": {0x42, 0x54}, "HADT": {0x48, 0x41, 0x44, 0x54}, "AST": {0x41, 0x53, 0x54}, "WARST": {0x57, 0x41, 0x52, 0x53, 0x54}, "UYST": {0x55, 0x59, 0x53, 0x54}, "CHAST": {0x43, 0x48, 0x41, 0x53, 0x54}, "AKDT": {0x41, 0x4b, 0x44, 0x54}, "ARST": {0x41, 0x52, 0x53, 0x54}, "ChST": {0x43, 0x68, 0x53, 0x54}, "LHST": {0x4c, 0x48, 0x53, 0x54}, "CLT": {0x43, 0x4c, 0x54}, "VET": {0x56, 0x45, 0x54}, "EST": {0x45, 0x53, 0x54}, "WAST": {0x57, 0x41, 0x53, 0x54}, "MDT": {0x4d, 0x44, 0x54}, "WIT": {0x57, 0x49, 0x54}, "OEZ": {0x4f, 0x45, 0x5a}, "CST": {0x43, 0x53, 0x54}, "WART": {0x57, 0x41, 0x52, 0x54}, "TMT": {0x54, 0x4d, 0x54}, "WEZ": {0x57, 0x45, 0x5a}, "WESZ": {0x57, 0x45, 0x53, 0x5a}, "GFT": {0x47, 0x46, 0x54}, "∅∅∅": {0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85, 0xe2, 0x88, 0x85}, "ADT": {0x41, 0x44, 0x54}, "CLST": {0x43, 0x4c, 0x53, 0x54}, "WIB": {0x57, 0x49, 0x42}, "MYT": {0x4d, 0x59, 0x54}, "LHDT": {0x4c, 0x48, 0x44, 0x54}, "UYT": {0x55, 0x59, 0x54}, "HKT": {0x48, 0x4b, 0x54}, "HKST": {0x48, 0x4b, 0x53, 0x54}, "SRT": {0x53, 0x52, 0x54}, "NZST": {0x4e, 0x5a, 0x53, 0x54}, "JST": {0x4a, 0x53, 0x54}},
	}
}

// Locale returns the current translators string locale
func (as *as_IN) Locale() string {
	return as.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'as_IN'
func (as *as_IN) PluralsCardinal() []locales.PluralRule {
	return as.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'as_IN'
func (as *as_IN) PluralsOrdinal() []locales.PluralRule {
	return as.pluralsOrdinal
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'as_IN'
func (as *as_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'as_IN'
func (as *as_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 5 || n == 7 || n == 8 || n == 9 || n == 10 {
		return locales.PluralRuleOne
	} else if n == 2 || n == 3 {
		return locales.PluralRuleTwo
	} else if n == 4 {
		return locales.PluralRuleFew
	} else if n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'as_IN'
func (as *as_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (as *as_IN) MonthAbbreviated(month time.Month) []byte {
	return as.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (as *as_IN) MonthsAbbreviated() [][]byte {
	return as.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (as *as_IN) MonthNarrow(month time.Month) []byte {
	return as.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (as *as_IN) MonthsNarrow() [][]byte {
	return as.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (as *as_IN) MonthWide(month time.Month) []byte {
	return as.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (as *as_IN) MonthsWide() [][]byte {
	return as.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (as *as_IN) WeekdayAbbreviated(weekday time.Weekday) []byte {
	return as.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (as *as_IN) WeekdaysAbbreviated() [][]byte {
	return as.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (as *as_IN) WeekdayNarrow(weekday time.Weekday) []byte {
	return as.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (as *as_IN) WeekdaysNarrow() [][]byte {
	return as.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (as *as_IN) WeekdayShort(weekday time.Weekday) []byte {
	return as.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (as *as_IN) WeekdaysShort() [][]byte {
	return as.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (as *as_IN) WeekdayWide(weekday time.Weekday) []byte {
	return as.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (as *as_IN) WeekdaysWide() [][]byte {
	return as.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'as_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtNumber(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(as.decimal) + len(as.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(as.decimal) - 1; j >= 0; j-- {
				b = append(b, as.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				for j := len(as.group) - 1; j >= 0; j-- {
					b = append(b, as.group[j])
				}

				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(as.minus) - 1; j >= 0; j-- {
			b = append(b, as.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'as_IN' and handles both Whole and Real numbers based on 'v'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (as *as_IN) FmtPercent(num float64, v uint64) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + len(as.decimal)
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(as.decimal) - 1; j >= 0; j-- {
				b = append(b, as.decimal[j])
			}

			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(as.minus) - 1; j >= 0; j-- {
			b = append(b, as.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, as.percent...)

	return b
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'as_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtCurrency(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := as.currencies[currency]
	l := len(s) + len(as.decimal) + len(as.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(as.decimal) - 1; j >= 0; j-- {
				b = append(b, as.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				for j := len(as.group) - 1; j >= 0; j-- {
					b = append(b, as.group[j])
				}

				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(as.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, as.currencyPositivePrefix[j])
	}

	if num < 0 {
		for j := len(as.minus) - 1; j >= 0; j-- {
			b = append(b, as.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, as.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'as_IN'
// in accounting notation. returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtAccounting(num float64, v uint64, currency currency.Type) []byte {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := as.currencies[currency]
	l := len(s) + len(as.decimal) + len(as.group)*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(as.decimal) - 1; j >= 0; j-- {
				b = append(b, as.decimal[j])
			}

			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				for j := len(as.group) - 1; j >= 0; j-- {
					b = append(b, as.group[j])
				}

				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
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

		for j := len(as.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, as.currencyNegativePrefix[j])
		}

		for j := len(as.minus) - 1; j >= 0; j-- {
			b = append(b, as.minus[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(as.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, as.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, as.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return b
}

// FmtDateShort returns the short date representation of 't' for 'as_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtDateShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2d}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateMedium returns the medium date representation of 't' for 'as_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtDateMedium(t time.Time) []byte {

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
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateLong returns the long date representation of 't' for 'as_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtDateLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, as.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtDateFull returns the full date representation of 't' for 'as_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtDateFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	b = append(b, as.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, as.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return b
}

// FmtTimeShort returns the short time representation of 't' for 'as_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtTimeShort(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Hour() < 12 {
		b = append(b, as.periodsAbbreviated[0]...)
	} else {
		b = append(b, as.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeMedium returns the medium time representation of 't' for 'as_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtTimeMedium(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
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

	if t.Hour() < 12 {
		b = append(b, as.periodsAbbreviated[0]...)
	} else {
		b = append(b, as.periodsAbbreviated[1]...)
	}

	return b
}

// FmtTimeLong returns the long time representation of 't' for 'as_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtTimeLong(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
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

	if t.Hour() < 12 {
		b = append(b, as.periodsAbbreviated[0]...)
	} else {
		b = append(b, as.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return b
}

// FmtTimeFull returns the full time representation of 't' for 'as_IN'
// returned as a []byte just in case the caller wishes to add more and can help
// avoid allocations; otherwise just cast as string.
func (as *as_IN) FmtTimeFull(t time.Time) []byte {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
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

	if t.Hour() < 12 {
		b = append(b, as.periodsAbbreviated[0]...)
	} else {
		b = append(b, as.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := as.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return b
}
