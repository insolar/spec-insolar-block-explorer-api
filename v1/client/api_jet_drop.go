/*
 * Insolar Explorer API
 *
 * [Insolar Explorer](https://github.com/insolar/block-explorer)'s REST API documentation.  Insolar Explorer is a service that allows users to search for and view the contents of individual transactions, records, lifelines, jet drops and jets.  * Record—minimum unit of storage that contains an associated request, response, and maintenance details * Lifeline—sequence of records for object state where an object is a smart contract instance * Jet drop—unit of storage for jets * Jet—groups of lifelines  Main API endpoints include: * Search endpoint: `/api/v1/search` * Endpoints for each entity type:   * Records: `/api/v1/jet-drops/{jet_drop_id}/records`   * Lifeline: `/api/v1/lifeline/{object_reference}/records`   * Pulse: `/api/v1/pulses/{pulse_number}`   * Pulses: `/api/v1/pulses`   * Jet drop by jet ID: `/api/v1/jets/{jet_id}/jet-drops`   * Jet drop by its ID: `/api/v1/jet-drops/{jet_drop_id}`   * Jet drop by pulse number: `/api/v1/pulses/{pulse_number}/jet-drops`  A search request first gets to the search endpoint where its the entity type `value` is identified.  It then is redirected to the appropriate endpoint for detailed metadata to show the user on the Insolar Explorer Web UI.  Search results can be paginated. Pagination parameters are optional and include: * Starting point that depends on the searched entity type: `from_index`, `from_pulse_number`, or `from_jet_drop_id` * Number of entries per page: `limit` * Number of entries to skip from the starting point to show on a new page: `offset`  For example, for 10 entries per page starting from the pulse number 431920 for a set of 200 entries: `<example.com>/<api_endpoint>?<params>&from_pulse_number=431920&limit=10&offset=0`.  Search results can additionally be sorted by `index` or `pulse` in the descending or ascending order.  For example, `<example.com>/<api_endpoint>?<params>&sort_by=-index`.  Search results shown to the user on the Insolar Explorer Web UI may miss some data. The `Reload` button gets the missed data via requests to the API with additional filtering parameters: * `jet_drop_id_gt` and `jet_drop_id_lt` to obtain missing data within a range of jet drops (greater than, less then) * `pulse_number_gt` and `pulse_number_lt` to obtain missing data within a range of pulses (greater than, less then) * `timestamp_gte` and `timestamp_lte` to obtain missing data within a timespan (greater than or equal, less then or equal)  For example, `<example.com>/<api_endpoint>?<params>&pulse_number_gt=431902&pulse_number_lt=431904`. 
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
JetDropByID jet drop by ID
Gets a jet drop by &#x60;jet_drop_id&#x60;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jetDropId Jet drop ID—a combination of `jet_id` with `pulse_number`.
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
    Offset optional.Int32
    FromJetDropId optional.String
    SortBy optional.String
    JetDropIdGt optional.String
    JetDropIdLt optional.String
}

/*
JetDropsByJetID jet drops by jet ID
Gets jet drops in a jet by &#x60;jet_id&#x60;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jetId Jet ID.
 * @param optional nil or *JetDropsByJetIDOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Number of entries per page.
 * @param "Offset" (optional.Int32) -  Number of entries to skip from the starting point.
 * @param "FromJetDropId" (optional.String) -  Specific jet drop ID to paginate from—a combination of jet_id` with `pulse_number`.
 * @param "SortBy" (optional.String) -  Pulse number-based sorting direction for the result set.
 * @param "JetDropIdGt" (optional.String) -  Starting point (>) for a range of jet drops.
 * @param "JetDropIdLt" (optional.String) -  Ending point (<) for a range of jet drops.
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
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FromJetDropId.IsSet() {
		localVarQueryParams.Add("from_jet_drop_id", parameterToString(localVarOptionals.FromJetDropId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortBy.IsSet() {
		localVarQueryParams.Add("sort_by", parameterToString(localVarOptionals.SortBy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.JetDropIdGt.IsSet() {
		localVarQueryParams.Add("jet_drop_id_gt", parameterToString(localVarOptionals.JetDropIdGt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.JetDropIdLt.IsSet() {
		localVarQueryParams.Add("jet_drop_id_lt", parameterToString(localVarOptionals.JetDropIdLt.Value(), ""))
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
JetDropsByPulseNumber jet drops by pulse number
Gets jet drops in a pulse by &#x60;pulse_number&#x60;.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pulseNumber Pulse number.
 * @param optional nil or *JetDropsByPulseNumberOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Number of entries per page.
 * @param "Offset" (optional.Int32) -  Number of entries to skip from the starting point.
 * @param "FromJetDropId" (optional.String) -  Specific jet drop ID to paginate from—a combination of jet_id` with `pulse_number`.
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
