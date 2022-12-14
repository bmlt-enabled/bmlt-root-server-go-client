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

// TokenCredentials struct for TokenCredentials
type TokenCredentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// NewTokenCredentials instantiates a new TokenCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenCredentials(password string, username string) *TokenCredentials {
	this := TokenCredentials{}
	this.Password = password
	this.Username = username
	return &this
}

// NewTokenCredentialsWithDefaults instantiates a new TokenCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCredentialsWithDefaults() *TokenCredentials {
	this := TokenCredentials{}
	return &this
}

// GetPassword returns the Password field value
func (o *TokenCredentials) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *TokenCredentials) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *TokenCredentials) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *TokenCredentials) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *TokenCredentials) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *TokenCredentials) SetUsername(v string) {
	o.Username = v
}

func (o TokenCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableTokenCredentials struct {
	value *TokenCredentials
	isSet bool
}

func (v NullableTokenCredentials) Get() *TokenCredentials {
	return v.value
}

func (v *NullableTokenCredentials) Set(val *TokenCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenCredentials(val *TokenCredentials) *NullableTokenCredentials {
	return &NullableTokenCredentials{value: val, isSet: true}
}

func (v NullableTokenCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


