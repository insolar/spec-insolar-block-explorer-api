/*
 * Insolar Explorer API
 *
 * # Insolar Explorer API documentation  Welcome to Insolar documentation for a REST-like API provided by Insolar Explorer.  [Insolar Explorer](https://github.com/insolar/block-explorer) is a service for searching for and viewing the contents of individual transactions, records, lifelines, jet drops and jets.  The API allows to search for, filter and view the contents of said entities.  ## Entities description  * Record—minimum unit of storage that contains an associated request, response, and maintenance details. * Lifeline—sequence of records for object state where an object is a smart contract instance. * Jet drop—unit of storage for jets. * Jet—group of lifelines.  * Jet drop ID—combination of jet id with pulse number. * Index—combination of pulse number with order (record number in a jet drop).   ## Filtering, pagination, sorting  API provides filtering based on a range of values: greater than and less than, greater than or equal to and less than or equal to.  API provides a combination of offset and seek pagination.  Pagination can be applied using: * Combination of a starting point (`from_*`), number of entries per page (`limit`) and number of entries to skip from the starting point (`offset`). * Just `limit` to get a limited array of the latest data. * Combination of the filtering parameters `*_gt`/`*_gte` and `*_lt`/`*_lte`, and `limit`.  Some requests can be sorted in the descending (`*_desc`)  or ascending (`*_asc`) order. 
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// Record Record representation.
type Record struct {
	// Record reference.
	Reference string `json:"reference,omitempty"`
	// Reference of previous record.
	PrevRecordReference string `json:"prev_record_reference,omitempty"`
	// Object reference.
	ObjectReference string `json:"object_reference,omitempty"`
	// Prototype reference.
	PrototypeReference string `json:"prototype_reference,omitempty"`
	// Record type.
	Type string `json:"type,omitempty"`
	// Unix timestamp.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Record payload.
	Payload string `json:"payload,omitempty"`
	// Record hash.
	Hash string `json:"hash,omitempty"`
	// Pulse number.
	PulseNumber int64 `json:"pulse_number,omitempty"`
	// Jet ID.
	JetId string `json:"jet_id,omitempty"`
	// Combination of `jet_id` with `pulse_number`.
	JetDropId string `json:"jet_drop_id,omitempty"`
	// Order—record number in a `jet drop`.
	Order int64 `json:"order,omitempty"`
	// Index—combination of `pulse_number` with `order` (record number in a jet drop).
	Index string `json:"index,omitempty"`
}
