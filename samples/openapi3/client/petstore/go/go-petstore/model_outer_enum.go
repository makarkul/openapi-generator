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

// OuterEnum the model 'OuterEnum'
type OuterEnum string

// List of OuterEnum
const (
	OUTERENUM_PLACED OuterEnum = "placed"
	OUTERENUM_APPROVED OuterEnum = "approved"
	OUTERENUM_DELIVERED OuterEnum = "delivered"
)

func (v *OuterEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OuterEnum(value)
	for _, existing := range []OuterEnum{ "placed", "approved", "delivered",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OuterEnum", value)
}

// Ptr returns reference to OuterEnum value
func (v OuterEnum) Ptr() *OuterEnum {
	return &v
}

type NullableOuterEnum struct {
	value *OuterEnum
	isSet bool
}

func (v NullableOuterEnum) Get() *OuterEnum {
	return v.value
}

func (v *NullableOuterEnum) Set(val *OuterEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOuterEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOuterEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOuterEnum(val *OuterEnum) *NullableOuterEnum {
	return &NullableOuterEnum{value: val, isSet: true}
}

func (v NullableOuterEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOuterEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

