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
// JetDropByIdResponse200 struct for JetDropByIdResponse200
type JetDropByIdResponse200 struct {
	// JetDrop ID is concantination of jet_id and pulse_number.
	JetDropId string `json:"jet_drop_id,omitempty"`
	// Next jet_drop_id.
	NextJetDropId []string `json:"next_jet_drop_id,omitempty"`
	// Previous jet_drop_id.
	PrevJetDropId []string `json:"prev_jet_drop_id,omitempty"`
	// Jet ID.
	JetId string `json:"jet_id,omitempty"`
	// Pulse number.
	PulseNumber int64 `json:"pulse_number,omitempty"`
	// Number of all records in the pulse.
	RecordAmount int64 `json:"record_amount,omitempty"`
	// Unix timestamp.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Record hash.
	Hash string `json:"hash,omitempty"`
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Link string `json:"link,omitempty"`
	ValidationFailures []PulsesResponse200ValidationFailures `json:"validation_failures"`
}
