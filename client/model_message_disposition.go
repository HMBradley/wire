/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MessageDisposition
type MessageDisposition struct {
	// formatVersion identifies the format version 30 
	FormatVersion string `json:"formatVersion,omitempty"`
	// testProductionCode identifies if test or production.  * `T` - Test * `P` - Production 
	TestProductionCode string `json:"testProductionCode,omitempty"`
	// MessageDuplicationCode  * ` ` - Original Message * `R` - Retrieval of an original message * `P` - Resend 
	MessageDuplicationCode string `json:"messageDuplicationCode,omitempty"`
	// MessageStatusIndicator  Outgoing Messages * `0` - In process or Intercepted * `2` - Successful with Accounting (Value) * `3` - Rejected due to Error Condition * `7` - Successful without Accounting (Non-Value)  Incoming Messages * `N` - Successful with Accounting (Value) * `S` - Successful without Accounting (Non-Value) 
	MessageStatusIndicator string `json:"messageStatusIndicator,omitempty"`
}
