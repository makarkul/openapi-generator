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

// NumberOnly struct for NumberOnly
type NumberOnly struct {
	JustNumber *float32 `json:"JustNumber,omitempty"`
}

// NewNumberOnly instantiates a new NumberOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberOnly() *NumberOnly {
	this := NumberOnly{}
	return &this
}

// NewNumberOnlyWithDefaults instantiates a new NumberOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberOnlyWithDefaults() *NumberOnly {
	this := NumberOnly{}
	return &this
}

// GetJustNumber returns the JustNumber field value if set, zero value otherwise.
func (o *NumberOnly) GetJustNumber()  {
	if o == nil || o.JustNumber == nil {
		var ret 
		return ret
	}
	return *o.JustNumber
}

// GetJustNumberOk returns a tuple with the JustNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberOnly) GetJustNumberOk() (*, bool) {
	if o == nil || o.JustNumber == nil {
		return nil, false
	}
	return o.JustNumber, true
}

// HasJustNumber returns a boolean if a field has been set.
func (o *NumberOnly) HasJustNumber() bool {
	if o != nil && o.JustNumber != nil {
		return true
	}

	return false
}

// SetJustNumber gets a reference to the given float32 and assigns it to the JustNumber field.
func (o *NumberOnly) SetJustNumber(v ) {
	o.JustNumber = &v
}

func (o NumberOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JustNumber != nil {
		toSerialize["JustNumber"] = o.JustNumber
	}
	return json.Marshal(toSerialize)
}

type NullableNumberOnly struct {
	value *NumberOnly
	isSet bool
}

func (v NullableNumberOnly) Get() *NumberOnly {
	return v.value
}

func (v *NullableNumberOnly) Set(val *NumberOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberOnly(val *NumberOnly) *NullableNumberOnly {
	return &NullableNumberOnly{value: val, isSet: true}
}

func (v NullableNumberOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


