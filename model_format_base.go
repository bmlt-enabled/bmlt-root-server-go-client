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

// checks if the FormatBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormatBase{}

// FormatBase struct for FormatBase
type FormatBase struct {
	WorldId *string `json:"worldId,omitempty"`
	Type *string `json:"type,omitempty"`
	Translations []FormatTranslation `json:"translations,omitempty"`
}

// NewFormatBase instantiates a new FormatBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormatBase() *FormatBase {
	this := FormatBase{}
	return &this
}

// NewFormatBaseWithDefaults instantiates a new FormatBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatBaseWithDefaults() *FormatBase {
	this := FormatBase{}
	return &this
}

// GetWorldId returns the WorldId field value if set, zero value otherwise.
func (o *FormatBase) GetWorldId() string {
	if o == nil || IsNil(o.WorldId) {
		var ret string
		return ret
	}
	return *o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatBase) GetWorldIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorldId) {
		return nil, false
	}
	return o.WorldId, true
}

// HasWorldId returns a boolean if a field has been set.
func (o *FormatBase) HasWorldId() bool {
	if o != nil && !IsNil(o.WorldId) {
		return true
	}

	return false
}

// SetWorldId gets a reference to the given string and assigns it to the WorldId field.
func (o *FormatBase) SetWorldId(v string) {
	o.WorldId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FormatBase) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatBase) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FormatBase) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FormatBase) SetType(v string) {
	o.Type = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *FormatBase) GetTranslations() []FormatTranslation {
	if o == nil || IsNil(o.Translations) {
		var ret []FormatTranslation
		return ret
	}
	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormatBase) GetTranslationsOk() ([]FormatTranslation, bool) {
	if o == nil || IsNil(o.Translations) {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *FormatBase) HasTranslations() bool {
	if o != nil && !IsNil(o.Translations) {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given []FormatTranslation and assigns it to the Translations field.
func (o *FormatBase) SetTranslations(v []FormatTranslation) {
	o.Translations = v
}

func (o FormatBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormatBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WorldId) {
		toSerialize["worldId"] = o.WorldId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Translations) {
		toSerialize["translations"] = o.Translations
	}
	return toSerialize, nil
}

type NullableFormatBase struct {
	value *FormatBase
	isSet bool
}

func (v NullableFormatBase) Get() *FormatBase {
	return v.value
}

func (v *NullableFormatBase) Set(val *FormatBase) {
	v.value = val
	v.isSet = true
}

func (v NullableFormatBase) IsSet() bool {
	return v.isSet
}

func (v *NullableFormatBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormatBase(val *FormatBase) *NullableFormatBase {
	return &NullableFormatBase{value: val, isSet: true}
}

func (v NullableFormatBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormatBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


