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

// AuthorizationError struct for AuthorizationError
type AuthorizationError struct {
	Message string `json:"message"`
}

// NewAuthorizationError instantiates a new AuthorizationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationError(message string) *AuthorizationError {
	this := AuthorizationError{}
	this.Message = message
	return &this
}

// NewAuthorizationErrorWithDefaults instantiates a new AuthorizationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationErrorWithDefaults() *AuthorizationError {
	this := AuthorizationError{}
	return &this
}

// GetMessage returns the Message field value
func (o *AuthorizationError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AuthorizationError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AuthorizationError) SetMessage(v string) {
	o.Message = v
}

func (o AuthorizationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationError struct {
	value *AuthorizationError
	isSet bool
}

func (v NullableAuthorizationError) Get() *AuthorizationError {
	return v.value
}

func (v *NullableAuthorizationError) Set(val *AuthorizationError) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationError) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationError(val *AuthorizationError) *NullableAuthorizationError {
	return &NullableAuthorizationError{value: val, isSet: true}
}

func (v NullableAuthorizationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


