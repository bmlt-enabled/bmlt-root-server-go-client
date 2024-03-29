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

// checks if the ServiceBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceBody{}

// ServiceBody struct for ServiceBody
type ServiceBody struct {
	ParentId int32 `json:"parentId"`
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	AdminUserId int32 `json:"adminUserId"`
	AssignedUserIds []int32 `json:"assignedUserIds"`
	Url string `json:"url"`
	Helpline string `json:"helpline"`
	Email string `json:"email"`
	WorldId string `json:"worldId"`
	Id int32 `json:"id"`
}

// NewServiceBody instantiates a new ServiceBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBody(parentId int32, name string, description string, type_ string, adminUserId int32, assignedUserIds []int32, url string, helpline string, email string, worldId string, id int32) *ServiceBody {
	this := ServiceBody{}
	this.ParentId = parentId
	this.Name = name
	this.Description = description
	this.Type = type_
	this.AdminUserId = adminUserId
	this.AssignedUserIds = assignedUserIds
	this.Url = url
	this.Helpline = helpline
	this.Email = email
	this.WorldId = worldId
	this.Id = id
	return &this
}

// NewServiceBodyWithDefaults instantiates a new ServiceBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBodyWithDefaults() *ServiceBody {
	this := ServiceBody{}
	return &this
}

// GetParentId returns the ParentId field value
func (o *ServiceBody) GetParentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetParentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *ServiceBody) SetParentId(v int32) {
	o.ParentId = v
}

// GetName returns the Name field value
func (o *ServiceBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceBody) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ServiceBody) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServiceBody) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *ServiceBody) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceBody) SetType(v string) {
	o.Type = v
}

// GetAdminUserId returns the AdminUserId field value
func (o *ServiceBody) GetAdminUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AdminUserId
}

// GetAdminUserIdOk returns a tuple with the AdminUserId field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetAdminUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminUserId, true
}

// SetAdminUserId sets field value
func (o *ServiceBody) SetAdminUserId(v int32) {
	o.AdminUserId = v
}

// GetAssignedUserIds returns the AssignedUserIds field value
func (o *ServiceBody) GetAssignedUserIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AssignedUserIds
}

// GetAssignedUserIdsOk returns a tuple with the AssignedUserIds field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetAssignedUserIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedUserIds, true
}

// SetAssignedUserIds sets field value
func (o *ServiceBody) SetAssignedUserIds(v []int32) {
	o.AssignedUserIds = v
}

// GetUrl returns the Url field value
func (o *ServiceBody) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ServiceBody) SetUrl(v string) {
	o.Url = v
}

// GetHelpline returns the Helpline field value
func (o *ServiceBody) GetHelpline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Helpline
}

// GetHelplineOk returns a tuple with the Helpline field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetHelplineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Helpline, true
}

// SetHelpline sets field value
func (o *ServiceBody) SetHelpline(v string) {
	o.Helpline = v
}

// GetEmail returns the Email field value
func (o *ServiceBody) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ServiceBody) SetEmail(v string) {
	o.Email = v
}

// GetWorldId returns the WorldId field value
func (o *ServiceBody) GetWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorldId, true
}

// SetWorldId sets field value
func (o *ServiceBody) SetWorldId(v string) {
	o.WorldId = v
}

// GetId returns the Id field value
func (o *ServiceBody) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceBody) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceBody) SetId(v int32) {
	o.Id = v
}

func (o ServiceBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parentId"] = o.ParentId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["type"] = o.Type
	toSerialize["adminUserId"] = o.AdminUserId
	toSerialize["assignedUserIds"] = o.AssignedUserIds
	toSerialize["url"] = o.Url
	toSerialize["helpline"] = o.Helpline
	toSerialize["email"] = o.Email
	toSerialize["worldId"] = o.WorldId
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableServiceBody struct {
	value *ServiceBody
	isSet bool
}

func (v NullableServiceBody) Get() *ServiceBody {
	return v.value
}

func (v *NullableServiceBody) Set(val *ServiceBody) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBody) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBody(val *ServiceBody) *NullableServiceBody {
	return &NullableServiceBody{value: val, isSet: true}
}

func (v NullableServiceBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


