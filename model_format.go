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

// checks if the Format type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Format{}

// Format struct for Format
type Format struct {
	WorldId string `json:"worldId"`
	Type string `json:"type"`
	Translations []FormatTranslation `json:"translations"`
	Id int32 `json:"id"`
}

// NewFormat instantiates a new Format object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormat(worldId string, type_ string, translations []FormatTranslation, id int32) *Format {
	this := Format{}
	this.WorldId = worldId
	this.Type = type_
	this.Translations = translations
	this.Id = id
	return &this
}

// NewFormatWithDefaults instantiates a new Format object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormatWithDefaults() *Format {
	this := Format{}
	return &this
}

// GetWorldId returns the WorldId field value
func (o *Format) GetWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value
// and a boolean to check if the value has been set.
func (o *Format) GetWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorldId, true
}

// SetWorldId sets field value
func (o *Format) SetWorldId(v string) {
	o.WorldId = v
}

// GetType returns the Type field value
func (o *Format) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Format) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Format) SetType(v string) {
	o.Type = v
}

// GetTranslations returns the Translations field value
func (o *Format) GetTranslations() []FormatTranslation {
	if o == nil {
		var ret []FormatTranslation
		return ret
	}

	return o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value
// and a boolean to check if the value has been set.
func (o *Format) GetTranslationsOk() ([]FormatTranslation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Translations, true
}

// SetTranslations sets field value
func (o *Format) SetTranslations(v []FormatTranslation) {
	o.Translations = v
}

// GetId returns the Id field value
func (o *Format) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Format) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Format) SetId(v int32) {
	o.Id = v
}

func (o Format) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Format) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["worldId"] = o.WorldId
	toSerialize["type"] = o.Type
	toSerialize["translations"] = o.Translations
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableFormat struct {
	value *Format
	isSet bool
}

func (v NullableFormat) Get() *Format {
	return v.value
}

func (v *NullableFormat) Set(val *Format) {
	v.value = val
	v.isSet = true
}

func (v NullableFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormat(val *Format) *NullableFormat {
	return &NullableFormat{value: val, isSet: true}
}

func (v NullableFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


