/*
 * Insolar Explorer API
 *
 * # Insolar Explorer API documentation  [Insolar Explorer](https://github.com/insolar/block-explorer) is a service that allows users to search for and view the contents of individual transactions, records, lifelines, jet drops and jets.  Insolar Explorer provides a REST-like API interface.  ## Basic entities  * Record—minimum unit of storage that contains an associated request, response, and maintenance details * Lifeline—sequence of records for object state where an object is a smart contract instance * Jet drop—unit of storage for jets * Jet—groups of lifelines  ## Filtering, pagination, sorting  API provides filtering based on a range of values: < > or ≤ ≥.  API provides a combination of offset and seek pagination.  Pagination can be applied using: * Combination of a starting point (`from_*`), number of entries per page (`limit`) and number of entries to skip from the starting point (`offset`). * Just `limit` to get a limited array of the latest data. * Combination of the filtering parameters `*_gt`/`*_gte` and `*_lt`/`*_lte`, and `limit`.  Some requests can be sorted in the descending (`*_desc`)  or ascending (`*_asc`) order. 
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// SearchJetDropMeta Meta data.
type SearchJetDropMeta struct {
	// Jet drop ID is a combination of `jet_id` with `pulse_number`.
	JetDropId string `json:"jet_drop_id,omitempty"`
}
