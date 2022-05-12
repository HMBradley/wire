// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &BeneficiaryIntermediaryFI{}

// BeneficiaryIntermediaryFI {4000}
type BeneficiaryIntermediaryFI struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// Financial Institution
	FinancialInstitution FinancialInstitution `json:"financialInstitution,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBeneficiaryIntermediaryFI returns a new BeneficiaryIntermediaryFI
func NewBeneficiaryIntermediaryFI(isVariable bool) *BeneficiaryIntermediaryFI {
	bifi := &BeneficiaryIntermediaryFI{
		tag:              TagBeneficiaryIntermediaryFI,
		isVariableLength: isVariable,
	}
	return bifi
}

// Parse takes the input string and parses the ReceiverDepositoryInstitution values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bifi *BeneficiaryIntermediaryFI) Parse(record string) (error, int) {
	if utf8.RuneCountInString(record) < 12 {
		return NewTagWrongLengthErr(12, len(record)), 0
	}
	bifi.tag = record[:6]
	bifi.FinancialInstitution.IdentificationCode = bifi.parseStringField(record[6:7])

	length := 7
	read := 0

	bifi.FinancialInstitution.Identifier, read = bifi.parseVariableStringField(record[length:], 34)
	length += read

	bifi.FinancialInstitution.Name, read = bifi.parseVariableStringField(record[length:], 35)
	length += read

	bifi.FinancialInstitution.Address.AddressLineOne, read = bifi.parseVariableStringField(record[length:], 35)
	length += read

	bifi.FinancialInstitution.Address.AddressLineTwo, read = bifi.parseVariableStringField(record[length:], 35)
	length += read

	bifi.FinancialInstitution.Address.AddressLineThree, read = bifi.parseVariableStringField(record[length:], 35)
	length += read

	return nil, length
}

func (bifi *BeneficiaryIntermediaryFI) UnmarshalJSON(data []byte) error {
	type Alias BeneficiaryIntermediaryFI
	aux := struct {
		*Alias
	}{
		(*Alias)(bifi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	bifi.tag = TagBeneficiaryIntermediaryFI
	return nil
}

// String writes BeneficiaryIntermediaryFI
func (bifi *BeneficiaryIntermediaryFI) String() string {
	var buf strings.Builder
	buf.Grow(181)
	buf.WriteString(bifi.tag)
	buf.WriteString(bifi.IdentificationCodeField())
	buf.WriteString(bifi.IdentifierField())
	buf.WriteString(bifi.NameField())
	buf.WriteString(bifi.AddressLineOneField())
	buf.WriteString(bifi.AddressLineTwoField())
	buf.WriteString(bifi.AddressLineThreeField())
	return buf.String()
}

// Validate performs WIRE format rule checks on BeneficiaryIntermediaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// If ID Code is present, Identifier is mandatory and vice versa.
func (bifi *BeneficiaryIntermediaryFI) Validate() error {
	if err := bifi.fieldInclusion(); err != nil {
		return err
	}
	if bifi.tag != TagBeneficiaryIntermediaryFI {
		return fieldError("tag", ErrValidTagForType, bifi.tag)
	}
	if err := bifi.isIdentificationCode(bifi.FinancialInstitution.IdentificationCode); err != nil {
		return fieldError("IdentificationCode", err, bifi.FinancialInstitution.IdentificationCode)
	}
	// Can only be these Identification Codes
	switch bifi.FinancialInstitution.IdentificationCode {
	case
		"B", "C", "D", "F", "U":
	default:
		return fieldError("IdentificationCode", ErrIdentificationCode, bifi.FinancialInstitution.IdentificationCode)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Identifier); err != nil {
		return fieldError("Identifier", err, bifi.FinancialInstitution.Identifier)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Name); err != nil {
		return fieldError("Name", err, bifi.FinancialInstitution.Name)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Address.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, bifi.FinancialInstitution.Address.AddressLineOne)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Address.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, bifi.FinancialInstitution.Address.AddressLineTwo)
	}
	if err := bifi.isAlphanumeric(bifi.FinancialInstitution.Address.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, bifi.FinancialInstitution.Address.AddressLineThree)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (bifi *BeneficiaryIntermediaryFI) fieldInclusion() error {
	if bifi.FinancialInstitution.IdentificationCode != "" && bifi.FinancialInstitution.Identifier == "" {
		return fieldError("BeneficiaryIntermediaryFI.FinancialInstitution.Identifier", ErrFieldRequired)
	}
	if bifi.FinancialInstitution.IdentificationCode == "" && bifi.FinancialInstitution.Identifier != "" {
		return fieldError("BeneficiaryIntermediaryFI.FinancialInstitution.IdentificationCode", ErrFieldRequired)
	}
	return nil
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (bifi *BeneficiaryIntermediaryFI) IdentificationCodeField() string {
	return bifi.alphaField(bifi.FinancialInstitution.IdentificationCode, 1)
}

// IdentifierField gets a string of the Identifier field
func (bifi *BeneficiaryIntermediaryFI) IdentifierField() string {
	return bifi.alphaVariableField(bifi.FinancialInstitution.Identifier, 34, bifi.isVariableLength)
}

// NameField gets a string of the Name field
func (bifi *BeneficiaryIntermediaryFI) NameField() string {
	return bifi.alphaVariableField(bifi.FinancialInstitution.Name, 35, bifi.isVariableLength)
}

// AddressLineOneField gets a string of AddressLineOne field
func (bifi *BeneficiaryIntermediaryFI) AddressLineOneField() string {
	return bifi.alphaVariableField(bifi.FinancialInstitution.Address.AddressLineOne, 35, bifi.isVariableLength)
}

// AddressLineTwoField gets a string of AddressLineTwo field
func (bifi *BeneficiaryIntermediaryFI) AddressLineTwoField() string {
	return bifi.alphaVariableField(bifi.FinancialInstitution.Address.AddressLineTwo, 35, bifi.isVariableLength)
}

// AddressLineThreeField gets a string of AddressLineThree field
func (bifi *BeneficiaryIntermediaryFI) AddressLineThreeField() string {
	return bifi.alphaVariableField(bifi.FinancialInstitution.Address.AddressLineThree, 35, bifi.isVariableLength)
}
