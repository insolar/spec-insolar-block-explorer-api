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
// SearchRecordMeta Meta data.
type SearchRecordMeta struct {
	// Object reference.
	ObjectReference string `json:"object_reference,omitempty"`
	// Index is concatenation of pulse_number and order.
	Index string `json:"index,omitempty"`
}
