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

// EnumArrays struct for EnumArrays
type EnumArrays struct {
	JustSymbol *string `json:"just_symbol,omitempty"`
	ArrayEnum *[]string `json:"array_enum,omitempty"`
}

// NewEnumArrays instantiates a new EnumArrays object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnumArrays() *EnumArrays {
	this := EnumArrays{}
	return &this
}

// NewEnumArraysWithDefaults instantiates a new EnumArrays object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnumArraysWithDefaults() *EnumArrays {
	this := EnumArrays{}
	return &this
}

// GetJustSymbol returns the JustSymbol field value if set, zero value otherwise.
func (o *EnumArrays) GetJustSymbol()  {
	if o == nil || o.JustSymbol == nil {
		var ret 
		return ret
	}
	return *o.JustSymbol
}

// GetJustSymbolOk returns a tuple with the JustSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumArrays) GetJustSymbolOk() (*, bool) {
	if o == nil || o.JustSymbol == nil {
		return nil, false
	}
	return o.JustSymbol, true
}

// HasJustSymbol returns a boolean if a field has been set.
func (o *EnumArrays) HasJustSymbol() bool {
	if o != nil && o.JustSymbol != nil {
		return true
	}

	return false
}

// SetJustSymbol gets a reference to the given string and assigns it to the JustSymbol field.
func (o *EnumArrays) SetJustSymbol(v ) {
	o.JustSymbol = &v
}

// GetArrayEnum returns the ArrayEnum field value if set, zero value otherwise.
func (o *EnumArrays) GetArrayEnum()  {
	if o == nil || o.ArrayEnum == nil {
		var ret 
		return ret
	}
	return *o.ArrayEnum
}

// GetArrayEnumOk returns a tuple with the ArrayEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnumArrays) GetArrayEnumOk() (*, bool) {
	if o == nil || o.ArrayEnum == nil {
		return nil, false
	}
	return o.ArrayEnum, true
}

// HasArrayEnum returns a boolean if a field has been set.
func (o *EnumArrays) HasArrayEnum() bool {
	if o != nil && o.ArrayEnum != nil {
		return true
	}

	return false
}

// SetArrayEnum gets a reference to the given []string and assigns it to the ArrayEnum field.
func (o *EnumArrays) SetArrayEnum(v ) {
	o.ArrayEnum = &v
}

func (o EnumArrays) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JustSymbol != nil {
		toSerialize["just_symbol"] = o.JustSymbol
	}
	if o.ArrayEnum != nil {
		toSerialize["array_enum"] = o.ArrayEnum
	}
	return json.Marshal(toSerialize)
}

type NullableEnumArrays struct {
	value *EnumArrays
	isSet bool
}

func (v NullableEnumArrays) Get() *EnumArrays {
	return v.value
}

func (v *NullableEnumArrays) Set(val *EnumArrays) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumArrays) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumArrays) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumArrays(val *EnumArrays) *NullableEnumArrays {
	return &NullableEnumArrays{value: val, isSet: true}
}

func (v NullableEnumArrays) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumArrays) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


