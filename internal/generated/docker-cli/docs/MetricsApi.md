# \MetricsApi

All URIs are relative to *http://localhost/var/run/docker-cli.sock*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitMetrics**](MetricsApi.md#SubmitMetrics) | **Post** /usage | 



## SubmitMetrics

> SubmitMetrics(ctx).MetricsCommand(metricsCommand).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    metricsCommand := *openapiclient.NewMetricsCommand() // MetricsCommand | Submits a new metrics entry

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsApi.SubmitMetrics(context.Background()).MetricsCommand(metricsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.SubmitMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsCommand** | [**MetricsCommand**](MetricsCommand.md) | Submits a new metrics entry | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

