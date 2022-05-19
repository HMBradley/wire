// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &FIBeneficiaryFIAdvice{}

// FIBeneficiaryFIAdvice is the financial institution beneficiary financial institution
type FIBeneficiaryFIAdvice struct {
	// tag
	tag string
	// Advice
	Advice Advice `json:"advice,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryFIAdvice returns a new FIBeneficiaryFIAdvice
func NewFIBeneficiaryFIAdvice() *FIBeneficiaryFIAdvice {
	fibfia := &FIBeneficiaryFIAdvice{
		tag: TagFIBeneficiaryFIAdvice,
	}
	return fibfia
}

// Parse takes the input string and parses the FIBeneficiaryFIAdvice values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fibfia *FIBeneficiaryFIAdvice) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 13 {
		return 0, NewTagWrongLengthErr(13, len(record))
	}

	var err error
	var length, read int

	if fibfia.tag, read, err = fibfia.parseTag(record); err != nil {
		return 0, fieldError("FIBeneficiaryFIAdvice.Tag", err)
	}
	length += read

	if read, err = fibfia.Advice.Parse(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
}

func (fibfia *FIBeneficiaryFIAdvice) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiaryFIAdvice
	aux := struct {
		*Alias
	}{
		(*Alias)(fibfia),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fibfia.tag = TagFIBeneficiaryFIAdvice
	return nil
}

// String writes FIBeneficiaryFIAdvice
func (fibfia *FIBeneficiaryFIAdvice) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(200)

	buf.WriteString(fibfia.tag)
	buf.WriteString(fibfia.Advice.String(isCompressed))

	return buf.String()
}

// Validate performs WIRE format rule checks on FIBeneficiaryFIAdvice and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fibfia *FIBeneficiaryFIAdvice) Validate() error {
	if fibfia.tag != TagFIBeneficiaryFIAdvice {
		return fieldError("tag", ErrValidTagForType, fibfia.tag)
	}
	if err := fibfia.isAdviceCode(fibfia.Advice.AdviceCode); err != nil {
		return fieldError("AdviceCode", err, fibfia.Advice.AdviceCode)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineOne); err != nil {
		return fieldError("LineOne", err, fibfia.Advice.LineOne)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineTwo); err != nil {
		return fieldError("LineTwo", err, fibfia.Advice.LineTwo)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineThree); err != nil {
		return fieldError("LineThree", err, fibfia.Advice.LineThree)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineFour); err != nil {
		return fieldError("LineFour", err, fibfia.Advice.LineFour)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineFive); err != nil {
		return fieldError("LineFive", err, fibfia.Advice.LineFive)
	}
	if err := fibfia.isAlphanumeric(fibfia.Advice.LineSix); err != nil {
		return fieldError("LineSix", err, fibfia.Advice.LineSix)
	}
	return nil
}
