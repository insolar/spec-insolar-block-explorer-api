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
// PulseResponse200 struct for PulseResponse200
type PulseResponse200 struct {
	// Pulse number.
	PulseNumber int64 `json:"pulse_number,omitempty"`
	// Previous pulse number.
	PrevPulseNumber int64 `json:"prev_pulse_number,omitempty"`
	// Next pulse number.
	NextPulseNumber int64 `json:"next_pulse_number,omitempty"`
	// Amount of all Jet Drop in the Pulse.
	JetDropAmount int64 `json:"jet_drop_amount,omitempty"`
	// Number of all records in the Pulse.
	RecordAmount int64 `json:"record_amount,omitempty"`
	// Unix timestamp.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Pulse fullness status.
	IsComplete bool `json:"is_complete,omitempty"`
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Link string `json:"link,omitempty"`
	ValidationFailures []PulsesResponse200ValidationFailures `json:"validation_failures"`
}
