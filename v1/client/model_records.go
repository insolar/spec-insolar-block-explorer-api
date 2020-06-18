/*
 * Insolar Explorer API
 *
 * [Insolar Explorer](https://github.com/insolar/block-explorer)'s REST API documentation.  Insolar Explorer is a service that allows users to search for and view the contents of individual transactions, Records, Lifelines, Jet Drops and Jets.  * Record—minimum unit of storage that contains an associated request, response, and maintenance details * Lifeline—sequence of Records for object state where an object is a smart contract instance * Jet Drop—unit of storage for Jets * Jet—groups of Lifelines 
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// Records Response codes.
type Records struct {
	// Total results.
	Total int64 `json:"total,omitempty"`
	Result []ObjectLifelineResponse200Result `json:"result,omitempty"`
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
	Link string `json:"link,omitempty"`
	ValidationFailures []PulsesResponse200ValidationFailures `json:"validation_failures,omitempty"`
}
