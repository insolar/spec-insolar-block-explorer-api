/*
 * Insolar Explorer API
 *
 * # Insolar Explorer API documentation  Welcome to Insolar documentation for a REST-like API provided by Insolar Explorer.  [Insolar Explorer](https://github.com/insolar/block-explorer) is a service for searching for and viewing the contents of individual transactions, records, lifelines, jet drops and jets.  The API allows to search for, filter and view the contents of said entities.  ## Entities description  * Record—minimum unit of storage that contains an associated request, response, and service information. * Lifeline—sequence of records. A record in a lifeline is an object state and an object is a smart contract instance. * Jet drop—unit of storage for jets. * Jet—group of lifelines.  * Jet drop ID—combination of jet id and pulse number. * Index—combination of pulse number and order (record number in a jet drop).   ## Filtering, pagination, sorting  Some requests have filtering, sorting, and pagination capabilities.  ### Filtering  Filtering can be of two types: * By an attribute value: filters out from the returned array all the entities with attribute values other than specified.    For example: if you specify `type=request` in a [request](#operation/jet_drop_records) that returns an array of records of different types, the array doesn't contain the `result` and `state` records.  * By a range of attribute values: filters out from the returned array entities with attribute values that are out of the specified range.    You can specify inclusive, non-inclusive, open-ended, or closed-ended ranges depending on the request.    All range filter parameters have one the following suffixes:   * `*_gt`—greater than   * `*_lt`—less than   * `*_gte`—greater than or equal to   * `*_lte`—less than or equal to    For example: if you specify `timestamp_gte=1597409219` and `timestamp_lte=1597409241` in a [request](#operation/pulses) that returns an array of pulses, the array doesn't contain pulses older than `51063280` and younger than `51063340`.  ### Sorting  The returned array is sorted by an attribute with a monotonically increasing value (`pulse_number`, `index`), by default in the descending order. Sorting is applied after filtering, if any.  For example: if you specify `sort_by=index_asc` in a [request](#operation/object_lifeline) that returns an array of records, the array is sorted in the ascending order of `index`.  ### Pagination  Pagination can be applied after filtering and sorting, and uses the following 3 parameters:  * `limit`—number of array items to return. * `offset`—number of items to skip. * `from_*`—pagination starting point: the value of a monotonically increasing attribute. 
 *
 * API version: 1.0.0
 * Contact: dev-support@insolar.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// JetDropApiService JetDropApi service
type JetDropApiService service

/*
JetDropByID Jet drop by ID
Gets the contents of a &#x60;jet_drop&#x60; given a &#x60;jet_drop_id&#x60; as a path parameter.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jetDropId Combination of `jet_id` and `pulse_number`.
@return JetDropByIdResponse200
*/
func (a *JetDropApiService) JetDropByID(ctx _context.Context, jetDropId string) (JetDropByIdResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JetDropByIdResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/jet-drops/{jet_drop_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"jet_drop_id"+"}", _neturl.QueryEscape(parameterToString(jetDropId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v JetDropByIdResponse200
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// JetDropsByJetIDOpts Optional parameters for the method 'JetDropsByJetID'
type JetDropsByJetIDOpts struct {
    Limit optional.Int32
    SortBy optional.String
    PulseNumberGte optional.Int32
    PulseNumberGt optional.Int32
    PulseNumberLte optional.Int32
    PulseNumberLt optional.Int32
}

/*
JetDropsByJetID Jet drops by jet ID
Gets an array of jet drops given a &#x60;jet_id&#x60; as a path parameter. Optionally, specify filtering and pagination parameters. For more information, refer to the [filtering, pagination, sorting](#section/Insolar-Explorer-API-documentation/Filtering-pagination-sorting) section.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jetId Jet ID.
 * @param optional nil or *JetDropsByJetIDOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Defines a number of entries to show per page.
 * @param "SortBy" (optional.String) -  Sorts by the `pulse_number` attribute of the returned object. Can take two values that specify the sorting direction: descending (`pulse_number_desc`) or ascending (`pulse_number_asc`). 
 * @param "PulseNumberGte" (optional.Int32) -  Defines the starting point for a returned range of pulses—greater than or equal to the specified `pulse_number`.
 * @param "PulseNumberGt" (optional.Int32) -  Defines the starting point for a returned range of pulses—greater than the specified `pulse_number`.
 * @param "PulseNumberLte" (optional.Int32) -  Defines the ending point for a returned range of pulses—less than equal to the specified `pulse_number`.
 * @param "PulseNumberLt" (optional.Int32) -  Defines the ending point for a returned range of pulses—less than the specified `pulse_number`.
@return JetDropsByJetIdResponse200
*/
func (a *JetDropApiService) JetDropsByJetID(ctx _context.Context, jetId string, localVarOptionals *JetDropsByJetIDOpts) (JetDropsByJetIdResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JetDropsByJetIdResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/jets/{jet_id}/jet-drops"
	localVarPath = strings.Replace(localVarPath, "{"+"jet_id"+"}", _neturl.QueryEscape(parameterToString(jetId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortBy.IsSet() {
		localVarQueryParams.Add("sort_by", parameterToString(localVarOptionals.SortBy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PulseNumberGte.IsSet() {
		localVarQueryParams.Add("pulse_number_gte", parameterToString(localVarOptionals.PulseNumberGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PulseNumberGt.IsSet() {
		localVarQueryParams.Add("pulse_number_gt", parameterToString(localVarOptionals.PulseNumberGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PulseNumberLte.IsSet() {
		localVarQueryParams.Add("pulse_number_lte", parameterToString(localVarOptionals.PulseNumberLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PulseNumberLt.IsSet() {
		localVarQueryParams.Add("pulse_number_lt", parameterToString(localVarOptionals.PulseNumberLt.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v JetDropsByJetIdResponse200
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// JetDropsByPulseNumberOpts Optional parameters for the method 'JetDropsByPulseNumber'
type JetDropsByPulseNumberOpts struct {
    Limit optional.Int32
    Offset optional.Int32
    FromJetDropId optional.String
}

/*
JetDropsByPulseNumber Jet drops by pulse number
Gets an array of jet drops given a &#x60;pulse_number&#x60; as a path parameter. Optionally, specify filtering and pagination parameters. For more information, refer to the [filtering, pagination, sorting](#section/Insolar-Explorer-API-documentation/Filtering-pagination-sorting) section.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulseNumber Pulse number.
 * @param optional nil or *JetDropsByPulseNumberOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Defines a number of entries to show per page.
 * @param "Offset" (optional.Int32) -  Defines a number of entries to skip from the starting point (`from_*`).
 * @param "FromJetDropId" (optional.String) -  Defines a specific `jet_drop_id` to paginate from.
@return JetDropsByJetIdResponse200
*/
func (a *JetDropApiService) JetDropsByPulseNumber(ctx _context.Context, pulseNumber int64, localVarOptionals *JetDropsByPulseNumberOpts) (JetDropsByJetIdResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  JetDropsByJetIdResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/pulses/{pulse_number}/jet-drops"
	localVarPath = strings.Replace(localVarPath, "{"+"pulse_number"+"}", _neturl.QueryEscape(parameterToString(pulseNumber, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if pulseNumber < 0 {
		return localVarReturnValue, nil, reportError("pulseNumber must be greater than 0")
	}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromJetDropId.IsSet() {
		localVarQueryParams.Add("from_jet_drop_id", parameterToString(localVarOptionals.FromJetDropId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v JetDropsByJetIdResponse200
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
