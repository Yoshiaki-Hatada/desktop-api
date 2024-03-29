/*
 * Docker Desktop CLI API
 *
 * This is the Docker Desktop CLI API
 *
 * API version: 0.1.0
 * Contact: support@docker.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dockercliapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type MetricsApi interface {

	/*
	 * PostMetrics Metrics endpoint
	 * Submits metrics about a given command
	 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiPostMetricsRequest
	 */
	PostMetrics(ctx _context.Context) ApiPostMetricsRequest

	/*
	 * PostMetricsExecute executes the request
	 */
	PostMetricsExecute(r ApiPostMetricsRequest) (*_nethttp.Response, error)
}

// MetricsApiService MetricsApi service
type MetricsApiService service

type ApiPostMetricsRequest struct {
	ctx            _context.Context
	ApiService     MetricsApi
	metricsCommand *MetricsCommand
}

func (r ApiPostMetricsRequest) MetricsCommand(metricsCommand MetricsCommand) ApiPostMetricsRequest {
	r.metricsCommand = &metricsCommand
	return r
}

func (r ApiPostMetricsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PostMetricsExecute(r)
}

/*
 * PostMetrics Metrics endpoint
 * Submits metrics about a given command
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostMetricsRequest
 */
func (a *MetricsApiService) PostMetrics(ctx _context.Context) ApiPostMetricsRequest {
	return ApiPostMetricsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *MetricsApiService) PostMetricsExecute(r ApiPostMetricsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetricsApiService.PostMetrics")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.metricsCommand == nil {
		return nil, reportError("metricsCommand is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.metricsCommand
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
