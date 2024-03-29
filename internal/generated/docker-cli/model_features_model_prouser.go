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

// FeaturesModelProuser struct for FeaturesModelProuser
type FeaturesModelProuser struct {
	// true when it's a Pro user
	Enabled bool `json:"enabled"`
}

// NewFeaturesModelProuser instantiates a new FeaturesModelProuser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturesModelProuser(enabled bool) *FeaturesModelProuser {
	this := FeaturesModelProuser{}
	this.Enabled = enabled
	return &this
}

// NewFeaturesModelProuserWithDefaults instantiates a new FeaturesModelProuser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturesModelProuserWithDefaults() *FeaturesModelProuser {
	this := FeaturesModelProuser{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *FeaturesModelProuser) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *FeaturesModelProuser) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *FeaturesModelProuser) SetEnabled(v bool) {
	o.Enabled = v
}

func (o FeaturesModelProuser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableFeaturesModelProuser struct {
	value *FeaturesModelProuser
	isSet bool
}

func (v NullableFeaturesModelProuser) Get() *FeaturesModelProuser {
	return v.value
}

func (v *NullableFeaturesModelProuser) Set(val *FeaturesModelProuser) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturesModelProuser) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturesModelProuser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturesModelProuser(val *FeaturesModelProuser) *NullableFeaturesModelProuser {
	return &NullableFeaturesModelProuser{value: val, isSet: true}
}

func (v NullableFeaturesModelProuser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturesModelProuser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
