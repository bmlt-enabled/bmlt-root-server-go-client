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

// FormatAllOf struct for FormatAllOf
type FormatAllOf struct {
	Id *int32 `json:"id,omitempty"`
}

// NewFormatAllOf instantiates a new FormatAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormatAllOf() *FormatAllOf {
	this := FormatAllOf{}
	return &this
}

// NewFormatAllOfWithDefaults instantiates a new FormatAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatAllOfWithDefaults() *FormatAllOf {
	this := FormatAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FormatAllOf) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatAllOf) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FormatAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FormatAllOf) SetId(v int32) {
	o.Id = &v
}

func (o FormatAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableFormatAllOf struct {
	value *FormatAllOf
	isSet bool
}

func (v NullableFormatAllOf) Get() *FormatAllOf {
	return v.value
}

func (v *NullableFormatAllOf) Set(val *FormatAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFormatAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFormatAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormatAllOf(val *FormatAllOf) *NullableFormatAllOf {
	return &NullableFormatAllOf{value: val, isSet: true}
}

func (v NullableFormatAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormatAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


