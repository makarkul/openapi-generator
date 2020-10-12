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

// Zebra struct for Zebra
type Zebra struct {
	Type *string `json:"type,omitempty"`
	ClassName string `json:"className"`
	AdditionalProperties map[string]interface{}
}

type _Zebra Zebra

// NewZebra instantiates a new Zebra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZebra(ClassName string, ) *Zebra {
	this := Zebra{}
	this.ClassName = ClassName
	return &this
}

// NewZebraWithDefaults instantiates a new Zebra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZebraWithDefaults() *Zebra {
	this := Zebra{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Zebra) GetType()  {
	if o == nil || o.Type == nil {
		var ret 
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zebra) GetTypeOk() (*, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Zebra) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Zebra) SetType(v ) {
	o.Type = &v
}

// GetClassName returns the ClassName field value
func (o *Zebra) GetClassName()  {
	if o == nil  {
		var ret 
		return ret
	}

	return o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value
// and a boolean to check if the value has been set.
func (o *Zebra) GetClassNameOk() (*, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClassName, true
}

// SetClassName sets field value
func (o *Zebra) SetClassName(v ) {
	o.ClassName = v
}

func (o Zebra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["className"] = o.ClassName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Zebra) UnmarshalJSON(bytes []byte) (err error) {
	varZebra := _Zebra{}

	if err = json.Unmarshal(bytes, &varZebra); err == nil {
		*o = Zebra(varZebra)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "className")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableZebra struct {
	value *Zebra
	isSet bool
}

func (v NullableZebra) Get() *Zebra {
	return v.value
}

func (v *NullableZebra) Set(val *Zebra) {
	v.value = val
	v.isSet = true
}

func (v NullableZebra) IsSet() bool {
	return v.isSet
}

func (v *NullableZebra) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZebra(val *Zebra) *NullableZebra {
	return &NullableZebra{value: val, isSet: true}
}

func (v NullableZebra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZebra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


