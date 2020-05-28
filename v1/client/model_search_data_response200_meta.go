/*
 * Insolar Block Explorer API
 *
 * BE description 
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// SearchDataResponse200Meta Meta data.
type SearchDataResponse200Meta struct {
	// Pulse Number.
	PulseNumber int64 `json:"pulse_number,omitempty"`
	// Record reference.
	RecordRef string `json:"record_ref,omitempty"`
	// Object reference.
	ObjectReference string `json:"object_reference,omitempty"`
}