/*
 * Insolar Explorer API
 *
 * # Insolar Explorer API documentation  Welcome to Insolar documentation for a REST-like API provided by Insolar Explorer.  [Insolar Explorer](https://github.com/insolar/block-explorer) is a service for searching for and viewing the contents of records, lifelines, jet drops, and jets.  Insolar Explorer provides an API that allows to search for said entities as well as request, filter, sort, and paginate their arrays.  ## Entities description  Basic Insolar entities that you can find in this specification are the following:  * Record—minimum unit of storage that contains an associated request, response, and service information.   Records can be of three types:   * Request. Most requests are caused by other requests inside the Plaftorm with the exception of original requests. The latter are invoked by users.   * Result. Each request has a corresponding execution result.   * State. Some execution results change object state. They are caused by so-called mutable requests, while immutable requests are simple read operations. * Lifeline—sequence of records. A record in a lifeline is an object state and an object is a smart contract instance. * Jet drop—unit of storage for jets. * Jet—group of lifelines.  Some of the entities have complex identifiers. For example:  * A jet drop ID is a combination of jet ID and pulse number. * A record's index is a combination of pulse number and order. * An order is a record number in a jet drop.  ## Filtering, pagination, sorting  Some requests have filtering, sorting, and pagination capabilities.  ### Filtering  Filtering can be of two types: * By an attribute value: filters out from the returned array all the entities with attribute values other than specified.    For example: if you specify `type=request` in a [request](#operation/jet_drop_records) that returns an array of records of different types, the array doesn't contain the `result` and `state` records.  * By a range of attribute values: filters out from the returned array entities with attribute values that are out of the specified range.    You can specify inclusive, non-inclusive, open-ended, or closed-ended ranges depending on the request.    All range filter parameters have one of the following suffixes:   * `*_gt`—greater than   * `*_lt`—less than   * `*_gte`—greater than or equal to   * `*_lte`—less than or equal to    For example: if you specify `timestamp_gte=1597409219` and `timestamp_lte=1597409241` in a [request](#operation/pulses) that returns an array of pulses, the array doesn't contain pulses older than `51063280` and younger than `51063340`.  ### Sorting  The returned array is sorted by an attribute with a monotonically increasing value (`pulse_number`, `index`), by default in the descending order. Sorting is applied after filtering, if any.  For example: if you specify `sort_by=index_asc` in a [request](#operation/object_lifeline) that returns an array of records, the array is sorted in the ascending order of `index`.  ### Pagination  Pagination can be applied after filtering and sorting, and uses the following 3 parameters:  * `limit`—number of array items to return. * `offset`—number of items to skip. * `from_*`—pagination starting point: the value of a monotonically increasing attribute. The `*` designates the property name to paginate by.  For example: if you specify `?limit=100&offset=5&from_index=47382564:2` in a [request](#operation/jet_drop_records) that returns an array of records, the array:  * Starts with a record with the 5th index relative to the index `47382564:2`. * Can consist of 100 entries maximum.  The number of entries may vary depending on the actual number of records requested. To show if the value of the `limit` parameter is higher or lower than the actual number of entries, each response to a paginated request contains the actual existing number in the `total` property. 
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

// PulseApiService PulseApi service
type PulseApiService service

/*
Pulse Pulse
Gets pulse by &#x60;pulse_number&#x60;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulseNumber Pulse number.
@return PulseResponse200
*/
func (a *PulseApiService) Pulse(ctx _context.Context, pulseNumber int64) (PulseResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PulseResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/pulses/{pulse_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"pulse_number"+"}", _neturl.QueryEscape(parameterToString(pulseNumber, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if pulseNumber < 0 {
		return localVarReturnValue, nil, reportError("pulseNumber must be greater than 0")
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
			var v PulseResponse200
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

// PulsesOpts Optional parameters for the method 'Pulses'
type PulsesOpts struct {
    Limit optional.Int32
    Offset optional.Int32
    FromPulseNumber optional.Int64
    TimestampGte optional.Int64
    TimestampLte optional.Int64
    PulseNumberGt optional.Int32
    PulseNumberGte optional.Int32
    PulseNumberLt optional.Int32
    PulseNumberLte optional.Int32
    SortBy optional.String
}

/*
Pulses Pulses
Gets an array of pulses.  Optionally, specify filtering, sorting, and pagination parameters. For more information, refer to the [filtering, pagination, sorting](#section/Insolar-Explorer-API-documentation/Filtering-pagination-sorting) section. 
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *PulsesOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Number of entries to show per page.
 * @param "Offset" (optional.Int32) -  Number of entries to skip from the starting point (`from_*`).
 * @param "FromPulseNumber" (optional.Int64) -  Specific `pulse_number` to paginate from.
 * @param "TimestampGte" (optional.Int64) -  Starting point for a range—greater than or equal to the specified `timestamp` in the Unix format.
 * @param "TimestampLte" (optional.Int64) -  Ending point for a range—less than or equal to the specified `timestamp` in the Unix format.
 * @param "PulseNumberGt" (optional.Int32) -  Starting point for a range of pulses—greater than the specified `pulse_number`.
 * @param "PulseNumberGte" (optional.Int32) -  Starting point for a range of pulses—greater than or equal to the specified `pulse_number`.
 * @param "PulseNumberLt" (optional.Int32) -  Ending point for a range of pulses—less than the specified `pulse_number`.
 * @param "PulseNumberLte" (optional.Int32) -  Ending point for a range of pulses—less than or equal to the specified `pulse_number`.
 * @param "SortBy" (optional.String) -  Sorting direction—ascending or descending relative to the monotonically increasing `pulse_number`.
@return PulsesResponse200
*/
func (a *PulseApiService) Pulses(ctx _context.Context, localVarOptionals *PulsesOpts) (PulsesResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PulsesResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/pulses"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromPulseNumber.IsSet() {
		localVarQueryParams.Add("from_pulse_number", parameterToString(localVarOptionals.FromPulseNumber.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimestampGte.IsSet() {
		localVarQueryParams.Add("timestamp_gte", parameterToString(localVarOptionals.TimestampGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimestampLte.IsSet() {
		localVarQueryParams.Add("timestamp_lte", parameterToString(localVarOptionals.TimestampLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PulseNumberGt.IsSet() {
		localVarQueryParams.Add("pulse_number_gt", parameterToString(localVarOptionals.PulseNumberGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PulseNumberGte.IsSet() {
		localVarQueryParams.Add("pulse_number_gte", parameterToString(localVarOptionals.PulseNumberGte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PulseNumberLt.IsSet() {
		localVarQueryParams.Add("pulse_number_lt", parameterToString(localVarOptionals.PulseNumberLt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PulseNumberLte.IsSet() {
		localVarQueryParams.Add("pulse_number_lte", parameterToString(localVarOptionals.PulseNumberLte.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortBy.IsSet() {
		localVarQueryParams.Add("sort_by", parameterToString(localVarOptionals.SortBy.Value(), ""))
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
			var v PulsesResponse200
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
