/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PaymentNotification is the payment notification
type PaymentNotification struct {
	// PaymentNotificationIndicator  * `0 - 6` - Reserved for market practice conventions. * `7 - 9` - Reserved for bilateral agreements between Fedwire senders and receivers. 
	PaymentNotificationIndicator string `json:"paymentNotificationIndicator,omitempty"`
	// ContactNotificationElectronicAddress
	ContactNotificationElectronicAddress string `json:"contactNotificationElectronicAddress,omitempty"`
	// ContactName
	ContactName string `json:"contactName,omitempty"`
	// ContactPhoneNumber
	ContactPhoneNumber string `json:"contactPhoneNumber,omitempty"`
	// ContactMobileNumber
	ContactMobileNumber string `json:"contactMobileNumber,omitempty"`
	// FaxNumber
	FaxNumber string `json:"faxNumber,omitempty"`
	// EndToEndIdentification
	EndToEndIdentification string `json:"endToEndIdentification,omitempty"`
}
