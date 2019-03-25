// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// UnstructuredAddenda is the unstructured addenda information
type UnstructuredAddenda struct {
	// tag
	tag string
	// AddendaLength  Addenda Length must be numeric, padded with leading zeros if less than four characters and must equal length of content in Addenda Information (e.g., if content of Addenda Information is 987 characters, Addenda Length must be 0987).
	AddendaLength string `json:"addendaLength,omitempty"`
	// Addenda
	Addenda string `json:"addenda,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewUnstructuredAddenda returns a new UnstructuredAddenda
func NewUnstructuredAddenda() UnstructuredAddenda  {
	ua := UnstructuredAddenda {
		tag: TagUnstructuredAddenda,
	}
	return ua
}

// Parse takes the input string and parses the UnstructuredAddenda values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ua *UnstructuredAddenda) Parse(record string) {
}

// String writes UnstructuredAddenda
func (ua *UnstructuredAddenda) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(8998)
	buf.WriteString(ua.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on UnstructuredAddenda and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ua *UnstructuredAddenda) Validate() error {
	if err := ua.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ua *UnstructuredAddenda) fieldInclusion() error {
	return nil
}
