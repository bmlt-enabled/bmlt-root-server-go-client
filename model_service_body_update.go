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

// ServiceBodyUpdate struct for ServiceBodyUpdate
type ServiceBodyUpdate struct {
	ParentId int32 `json:"parentId"`
	Name string `json:"name"`
	Description string `json:"description"`
	Type string `json:"type"`
	AdminUserId int32 `json:"adminUserId"`
	AssignedUserIds []int32 `json:"assignedUserIds"`
	Url *string `json:"url,omitempty"`
	Helpline *string `json:"helpline,omitempty"`
	Email *string `json:"email,omitempty"`
	WorldId *string `json:"worldId,omitempty"`
}

// NewServiceBodyUpdate instantiates a new ServiceBodyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBodyUpdate(parentId int32, name string, description string, type_ string, adminUserId int32, assignedUserIds []int32) *ServiceBodyUpdate {
	this := ServiceBodyUpdate{}
	this.ParentId = parentId
	this.Name = name
	this.Description = description
	this.Type = type_
	this.AdminUserId = adminUserId
	this.AssignedUserIds = assignedUserIds
	return &this
}

// NewServiceBodyUpdateWithDefaults instantiates a new ServiceBodyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBodyUpdateWithDefaults() *ServiceBodyUpdate {
	this := ServiceBodyUpdate{}
	return &this
}

// GetParentId returns the ParentId field value
func (o *ServiceBodyUpdate) GetParentId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *ServiceBodyUpdate) GetParentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *ServiceBodyUpdate) SetParentId(v int32) {
	o.ParentId = v
}

// GetName returns the Name field value
func (o *ServiceBodyUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceBodyUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceBodyUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ServiceBodyUpdate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServiceBodyUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServiceBodyUpdate) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *ServiceBodyUpdate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceBodyUpdate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceBodyUpdate) SetType(v string) {
	o.Type = v
}

// GetAdminUserId returns the AdminUserId field value
func (o *ServiceBodyUpdate) GetAdminUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AdminUserId
}

// GetAdminUserIdOk returns a tuple with the AdminUserId field value
// and a boolean to check if the value has been set.
func (o *ServiceBodyUpdate) GetAdminUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminUserId, true
}

// SetAdminUserId sets field value
func (o *ServiceBodyUpdate) SetAdminUserId(v int32) {
	o.AdminUserId = v
}

// GetAssignedUserIds returns the AssignedUserIds field value
func (o *ServiceBodyUpdate) GetAssignedUserIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AssignedUserIds
}

// GetAssignedUserIdsOk returns a tuple with the AssignedUserIds field value
// and a boolean to check if the value has been set.
func (o *ServiceBodyUpdate) GetAssignedUserIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedUserIds, true
}

// SetAssignedUserIds sets field value
func (o *ServiceBodyUpdate) SetAssignedUserIds(v []int32) {
	o.AssignedUserIds = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ServiceBodyUpdate) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBodyUpdate) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ServiceBodyUpdate) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ServiceBodyUpdate) SetUrl(v string) {
	o.Url = &v
}

// GetHelpline returns the Helpline field value if set, zero value otherwise.
func (o *ServiceBodyUpdate) GetHelpline() string {
	if o == nil || o.Helpline == nil {
		var ret string
		return ret
	}
	return *o.Helpline
}

// GetHelplineOk returns a tuple with the Helpline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBodyUpdate) GetHelplineOk() (*string, bool) {
	if o == nil || o.Helpline == nil {
		return nil, false
	}
	return o.Helpline, true
}

// HasHelpline returns a boolean if a field has been set.
func (o *ServiceBodyUpdate) HasHelpline() bool {
	if o != nil && o.Helpline != nil {
		return true
	}

	return false
}

// SetHelpline gets a reference to the given string and assigns it to the Helpline field.
func (o *ServiceBodyUpdate) SetHelpline(v string) {
	o.Helpline = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ServiceBodyUpdate) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBodyUpdate) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ServiceBodyUpdate) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ServiceBodyUpdate) SetEmail(v string) {
	o.Email = &v
}

// GetWorldId returns the WorldId field value if set, zero value otherwise.
func (o *ServiceBodyUpdate) GetWorldId() string {
	if o == nil || o.WorldId == nil {
		var ret string
		return ret
	}
	return *o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBodyUpdate) GetWorldIdOk() (*string, bool) {
	if o == nil || o.WorldId == nil {
		return nil, false
	}
	return o.WorldId, true
}

// HasWorldId returns a boolean if a field has been set.
func (o *ServiceBodyUpdate) HasWorldId() bool {
	if o != nil && o.WorldId != nil {
		return true
	}

	return false
}

// SetWorldId gets a reference to the given string and assigns it to the WorldId field.
func (o *ServiceBodyUpdate) SetWorldId(v string) {
	o.WorldId = &v
}

func (o ServiceBodyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["parentId"] = o.ParentId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["adminUserId"] = o.AdminUserId
	}
	if true {
		toSerialize["assignedUserIds"] = o.AssignedUserIds
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Helpline != nil {
		toSerialize["helpline"] = o.Helpline
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.WorldId != nil {
		toSerialize["worldId"] = o.WorldId
	}
	return json.Marshal(toSerialize)
}

type NullableServiceBodyUpdate struct {
	value *ServiceBodyUpdate
	isSet bool
}

func (v NullableServiceBodyUpdate) Get() *ServiceBodyUpdate {
	return v.value
}

func (v *NullableServiceBodyUpdate) Set(val *ServiceBodyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBodyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBodyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBodyUpdate(val *ServiceBodyUpdate) *NullableServiceBodyUpdate {
	return &NullableServiceBodyUpdate{value: val, isSet: true}
}

func (v NullableServiceBodyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBodyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


