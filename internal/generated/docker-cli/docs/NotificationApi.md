# \NotificationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowToast**](NotificationApi.md#ShowToast) | **Post** /notifications | GUI system toast endpoint



## ShowToast

> ShowToast(ctx).NotificationModel(notificationModel).Execute()

GUI system toast endpoint



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
    notificationModel := *openapiclient.NewNotificationModel() // NotificationModel | Shows a notification in the form of a system toast

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.ShowToast(context.Background()).NotificationModel(notificationModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.ShowToast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowToastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationModel** | [**NotificationModel**](NotificationModel.md) | Shows a notification in the form of a system toast | 

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

