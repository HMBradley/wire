// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &CurrencyInstructedAmount{}

// CurrencyInstructedAmount is the currency instructed amount
type CurrencyInstructedAmount struct {
	// tag
	tag string
	// SwiftFieldTag
	SwiftFieldTag string `json:"swiftFieldTag"`
	// Amount is the instructed amount
	// Amount  Must begin with at least one numeric character (0-9) and contain only one decimal comma marker
	// (e.g., $1,234.56 should be entered as 1234,56 and $0.99 should be entered as
	Amount string `json:"amount"`
	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewCurrencyInstructedAmount returns a new CurrencyInstructedAmount
func NewCurrencyInstructedAmount() *CurrencyInstructedAmount {
	cia := &CurrencyInstructedAmount{
		tag: TagCurrencyInstructedAmount,
	}
	return cia
}

// Parse takes the input string and parses the CurrencyInstructedAmount values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (cia *CurrencyInstructedAmount) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 8 {
		return 0, NewTagWrongLengthErr(8, len(record))
	}

	var err error
	var length, read int

	if cia.tag, read, err = cia.parseTag(record); err != nil {
		return 0, fieldError("CurrencyInstructedAmount.Tag", err)
	}
	length += read

	if cia.SwiftFieldTag, read, err = cia.parseVariableStringField(record[length:], 5); err != nil {
		return 0, fieldError("SwiftFieldTag", err)
	}
	length += read

	if cia.Amount, read, err = cia.parseVariableStringField(record[length:], 18); err != nil {
		return 0, fieldError("Amount", err)
	}
	length += read

	return length, nil
}

func (cia *CurrencyInstructedAmount) UnmarshalJSON(data []byte) error {
	type Alias CurrencyInstructedAmount
	aux := struct {
		*Alias
	}{
		(*Alias)(cia),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	cia.tag = TagCurrencyInstructedAmount
	return nil
}

// String writes CurrencyInstructedAmount
func (cia *CurrencyInstructedAmount) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(29)

	buf.WriteString(cia.tag)
	buf.WriteString(cia.SwiftFieldTagField(isCompressed))
	buf.WriteString(cia.AmountField(isCompressed))

	return buf.String()
}

// Validate performs WIRE format rule checks on CurrencyInstructedAmount and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (cia *CurrencyInstructedAmount) Validate() error {
	if cia.tag != TagCurrencyInstructedAmount {
		return fieldError("tag", ErrValidTagForType, cia.tag)
	}
	if err := cia.isAlphanumeric(cia.SwiftFieldTag); err != nil {
		return fieldError("SwiftFieldTag", err, cia.SwiftFieldTag)
	}
	if err := cia.isAmount(cia.Amount); err != nil {
		return fieldError("Amount", err, cia.Amount)
	}
	return nil
}

// SwiftFieldTagField gets a string of the SwiftFieldTag field
func (cia *CurrencyInstructedAmount) SwiftFieldTagField(isCompressed bool) string {
	return cia.alphaVariableField(cia.SwiftFieldTag, 5, isCompressed)
}

// ToDo: The spec isn't clear if this is padded with zeros or not, so for now it is

// AmountField gets a string of the AmountTag field
func (cia *CurrencyInstructedAmount) AmountField(isCompressed bool) string {
	return cia.alphaVariableField(cia.Amount, 18, isCompressed)
}
