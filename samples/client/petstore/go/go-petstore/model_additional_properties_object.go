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

// AdditionalPropertiesObject struct for AdditionalPropertiesObject
type AdditionalPropertiesObject struct {
	map[string]map[string]map[string]interface{}
	Name *string `json:"name,omitempty"`
}

// NewAdditionalPropertiesObject instantiates a new AdditionalPropertiesObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalPropertiesObject() *AdditionalPropertiesObject {
	this := AdditionalPropertiesObject{}
	return &this
}

// NewAdditionalPropertiesObjectWithDefaults instantiates a new AdditionalPropertiesObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalPropertiesObjectWithDefaults() *AdditionalPropertiesObject {
	this := AdditionalPropertiesObject{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdditionalPropertiesObject) GetName()  {
	if o == nil || o.Name == nil {
		var ret 
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalPropertiesObject) GetNameOk() (*, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdditionalPropertiesObject) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdditionalPropertiesObject) SetName(v ) {
	o.Name = &v
}

func (o AdditionalPropertiesObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedmap[string]map[string]map[string]interface{}, errmap[string]map[string]map[string]interface{} := json.Marshal(o.map[string]map[string]map[string]interface{})
	if errmap[string]map[string]map[string]interface{} != nil {
		return []byte{}, errmap[string]map[string]map[string]interface{}
	}
	errmap[string]map[string]map[string]interface{} = json.Unmarshal([]byte(serializedmap[string]map[string]map[string]interface{}), &toSerialize)
	if errmap[string]map[string]map[string]interface{} != nil {
		return []byte{}, errmap[string]map[string]map[string]interface{}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalPropertiesObject struct {
	value *AdditionalPropertiesObject
	isSet bool
}

func (v NullableAdditionalPropertiesObject) Get() *AdditionalPropertiesObject {
	return v.value
}

func (v *NullableAdditionalPropertiesObject) Set(val *AdditionalPropertiesObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalPropertiesObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalPropertiesObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalPropertiesObject(val *AdditionalPropertiesObject) *NullableAdditionalPropertiesObject {
	return &NullableAdditionalPropertiesObject{value: val, isSet: true}
}

func (v NullableAdditionalPropertiesObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalPropertiesObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


