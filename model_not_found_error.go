/*
BMLT

BMLT Admin API Documentation

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmlt

import (
	"encoding/json"
)

// checks if the NotFoundError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotFoundError{}

// NotFoundError struct for NotFoundError
type NotFoundError struct {
	Message string `json:"message"`
}

// NewNotFoundError instantiates a new NotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotFoundError(message string) *NotFoundError {
	this := NotFoundError{}
	this.Message = message
	return &this
}

// NewNotFoundErrorWithDefaults instantiates a new NotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotFoundErrorWithDefaults() *NotFoundError {
	this := NotFoundError{}
	return &this
}

// GetMessage returns the Message field value
func (o *NotFoundError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NotFoundError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NotFoundError) SetMessage(v string) {
	o.Message = v
}

func (o NotFoundError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotFoundError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableNotFoundError struct {
	value *NotFoundError
	isSet bool
}

func (v NullableNotFoundError) Get() *NotFoundError {
	return v.value
}

func (v *NullableNotFoundError) Set(val *NotFoundError) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundError) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundError(val *NotFoundError) *NullableNotFoundError {
	return &NullableNotFoundError{value: val, isSet: true}
}

func (v NullableNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


