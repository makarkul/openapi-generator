/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
)

// AdditionalPropertiesInteger struct for AdditionalPropertiesInteger
type AdditionalPropertiesInteger struct {
	map[string]int32
	Name *string `json:"name,omitempty"`
}

// NewAdditionalPropertiesInteger instantiates a new AdditionalPropertiesInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalPropertiesInteger() *AdditionalPropertiesInteger {
	this := AdditionalPropertiesInteger{}
	return &this
}

// NewAdditionalPropertiesIntegerWithDefaults instantiates a new AdditionalPropertiesInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalPropertiesIntegerWithDefaults() *AdditionalPropertiesInteger {
	this := AdditionalPropertiesInteger{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdditionalPropertiesInteger) GetName()  {
	if o == nil || o.Name == nil {
		var ret 
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalPropertiesInteger) GetNameOk() (*, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdditionalPropertiesInteger) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdditionalPropertiesInteger) SetName(v ) {
	o.Name = &v
}

func (o AdditionalPropertiesInteger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedmap[string]int32, errmap[string]int32 := json.Marshal(o.map[string]int32)
	if errmap[string]int32 != nil {
		return []byte{}, errmap[string]int32
	}
	errmap[string]int32 = json.Unmarshal([]byte(serializedmap[string]int32), &toSerialize)
	if errmap[string]int32 != nil {
		return []byte{}, errmap[string]int32
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalPropertiesInteger struct {
	value *AdditionalPropertiesInteger
	isSet bool
}

func (v NullableAdditionalPropertiesInteger) Get() *AdditionalPropertiesInteger {
	return v.value
}

func (v *NullableAdditionalPropertiesInteger) Set(val *AdditionalPropertiesInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalPropertiesInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalPropertiesInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalPropertiesInteger(val *AdditionalPropertiesInteger) *NullableAdditionalPropertiesInteger {
	return &NullableAdditionalPropertiesInteger{value: val, isSet: true}
}

func (v NullableAdditionalPropertiesInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalPropertiesInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


