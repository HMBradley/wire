// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"errors"
	"fmt"
)

var (
	// Errors specific to validation

	//ErrTypeCode is returned when there's an invalid TypeCode tag {1510}
	ErrTypeCode = errors.New("is an invalid type code")
	//ErrSubTypeCode is returned when there's an invalid SubTypeCode tag {1510}
	ErrSubTypeCode = errors.New("is an invalid sub type Code")
	// ErrLocalInstrumentCode is returned when there is an invalid LocalInstrumentCode tag {3610}
	ErrLocalInstrumentCode = errors.New("is an invalid local instrument Code")
	// ErrPaymentNotificationIndicator is returned when there is an invalid PaymentNotificationIndicator {3620}
	ErrPaymentNotificationIndicator = errors.New("is an invalid payment notification indicator")
	// FileHeader errors

	// ErrRecordSize is given when there's an invalid record size
	ErrRecordSize = errors.New("is not 094")
)

// FieldError is returned for errors at a field level in a record
type FieldError struct {
	FieldName string      // field name where error happened
	Value     interface{} // value that cause error
	Err       error       // context of the error.
	Msg       string      // deprecated
}

// Error message is constructed
// FieldName Msg Value
// ToDo:
// Example1:
// Example2:
func (e *FieldError) Error() string {
	return fmt.Sprintf("%s %v %s", e.FieldName, e.Value, e.Err)
}

// Unwrap implements the base.UnwrappableError interface for FieldError
func (e *FieldError) Unwrap() error {
	return e.Err
}

func fieldError(field string, err error, values ...interface{}) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*FieldError); ok {
		return err
	}
	fe := FieldError{
		FieldName: field,
		Err:       err,
	}
	// only the first value counts
	if len(values) > 0 {
		fe.Value = values[0]
	}
	return &fe
}
