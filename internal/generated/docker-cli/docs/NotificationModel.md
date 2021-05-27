# NotificationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 

## Methods

### NewNotificationModel

`func NewNotificationModel() *NotificationModel`

NewNotificationModel instantiates a new NotificationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationModelWithDefaults

`func NewNotificationModelWithDefaults() *NotificationModel`

NewNotificationModelWithDefaults instantiates a new NotificationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NotificationModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NotificationModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBody

`func (o *NotificationModel) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *NotificationModel) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *NotificationModel) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *NotificationModel) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetLevel

`func (o *NotificationModel) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *NotificationModel) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *NotificationModel) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *NotificationModel) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


