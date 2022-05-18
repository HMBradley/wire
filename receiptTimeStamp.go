// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &ReceiptTimeStamp{}

// ReceiptTimeStamp is the receipt time stamp of the wire
type ReceiptTimeStamp struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// ReceiptDate is the receipt date
	ReceiptDate string `json:"receiptDate,omitempty"`
	// ReceiptTime is the receipt time
	ReceiptTime string `json:"receiptTime,omitempty"`
	// ApplicationIdentification
	ReceiptApplicationIdentification string `json:"receiptApplicationIdentification,omitempty"`

	// validator is composed for data validation
	// validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewReceiptTimeStamp returns a new ReceiptTimeStamp
func NewReceiptTimeStamp(isVariable bool) *ReceiptTimeStamp {
	rts := &ReceiptTimeStamp{
		tag:              TagReceiptTimeStamp,
		isVariableLength: isVariable,
	}
	return rts
}

// Parse takes the input string and parses the ReceiptTimeStamp values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rts *ReceiptTimeStamp) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 9 {
		return 0, NewTagWrongLengthErr(9, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if rts.tag, read, err = rts.parseTag(record); err != nil {
		return 0, fieldError("ReceiptTimeStamp.Tag", err)
	}
	length += read

	if rts.ReceiptDate, read, err = rts.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("ReceiptDate", err)
	}
	length += read

	if rts.ReceiptTime, read, err = rts.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("ReceiptTime", err)
	}
	length += read

	if rts.ReceiptApplicationIdentification, read, err = rts.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("ReceiptApplicationIdentification", err)
	}
	length += read

	return length, nil
}

func (rts *ReceiptTimeStamp) UnmarshalJSON(data []byte) error {
	type Alias ReceiptTimeStamp
	aux := struct {
		*Alias
	}{
		(*Alias)(rts),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rts.tag = TagReceiptTimeStamp
	return nil
}

// String writes ReceiptTimeStamp
func (rts *ReceiptTimeStamp) String() string {
	var buf strings.Builder
	buf.Grow(18)

	buf.WriteString(rts.tag)
	buf.WriteString(rts.ReceiptDateField())
	buf.WriteString(rts.ReceiptTimeField())
	buf.WriteString(rts.ReceiptApplicationIdentificationField())

	return buf.String()
}

// Validate performs WIRE format rule checks on ReceiptTimeStamp and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rts *ReceiptTimeStamp) Validate() error {
	// Currently no validation as the FED is responsible for the values
	if rts.tag != TagReceiptTimeStamp {
		return fieldError("tag", ErrValidTagForType, rts.tag)
	}
	return nil
}

// ReceiptDateField gets a string of the ReceiptDate field
func (rts *ReceiptTimeStamp) ReceiptDateField() string {
	return rts.alphaVariableField(rts.ReceiptDate, 4, rts.isVariableLength)
}

// ReceiptTimeField gets a string of the ReceiptTime field
func (rts *ReceiptTimeStamp) ReceiptTimeField() string {
	return rts.alphaVariableField(rts.ReceiptTime, 4, rts.isVariableLength)
}

// ReceiptApplicationIdentificationField gets a string of the ReceiptApplicationIdentification field
func (rts *ReceiptTimeStamp) ReceiptApplicationIdentificationField() string {
	return rts.alphaVariableField(rts.ReceiptApplicationIdentification, 4, rts.isVariableLength)
}
