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
	"encoding/json"
)

// MetricsCommand struct for MetricsCommand
type MetricsCommand struct {
	Command *string `json:"command,omitempty"`
	Context *string `json:"context,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewMetricsCommand instantiates a new MetricsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsCommand() *MetricsCommand {
	this := MetricsCommand{}
	return &this
}

// NewMetricsCommandWithDefaults instantiates a new MetricsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsCommandWithDefaults() *MetricsCommand {
	this := MetricsCommand{}
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *MetricsCommand) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsCommand) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *MetricsCommand) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *MetricsCommand) SetCommand(v string) {
	o.Command = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *MetricsCommand) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsCommand) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *MetricsCommand) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *MetricsCommand) SetContext(v string) {
	o.Context = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MetricsCommand) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsCommand) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MetricsCommand) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *MetricsCommand) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MetricsCommand) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsCommand) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MetricsCommand) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MetricsCommand) SetStatus(v string) {
	o.Status = &v
}

func (o MetricsCommand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsCommand struct {
	value *MetricsCommand
	isSet bool
}

func (v NullableMetricsCommand) Get() *MetricsCommand {
	return v.value
}

func (v *NullableMetricsCommand) Set(val *MetricsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsCommand(val *MetricsCommand) *NullableMetricsCommand {
	return &NullableMetricsCommand{value: val, isSet: true}
}

func (v NullableMetricsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


