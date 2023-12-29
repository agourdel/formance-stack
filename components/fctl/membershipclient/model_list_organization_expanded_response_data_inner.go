/*
Membership API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membershipclient

import (
	"encoding/json"
)

// checks if the ListOrganizationExpandedResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrganizationExpandedResponseDataInner{}

// ListOrganizationExpandedResponseDataInner struct for ListOrganizationExpandedResponseDataInner
type ListOrganizationExpandedResponseDataInner struct {
	// Organization name
	Name string `json:"name"`
	DefaultOrganizationAccess []string `json:"defaultOrganizationAccess,omitempty"`
	DefaultStackAccess []string `json:"defaultStackAccess,omitempty"`
	// Organization ID
	Id string `json:"id"`
	// Owner ID
	OwnerId string `json:"ownerId"`
	// Number of available stacks
	AvailableStacks *int32 `json:"availableStacks,omitempty"`
	// Number of available sandboxes
	AvailableSandboxes *int32 `json:"availableSandboxes,omitempty"`
	TotalStacks *int32 `json:"totalStacks,omitempty"`
	TotalUsers *int32 `json:"totalUsers,omitempty"`
	Owner *User `json:"owner,omitempty"`
}

// NewListOrganizationExpandedResponseDataInner instantiates a new ListOrganizationExpandedResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrganizationExpandedResponseDataInner(name string, id string, ownerId string) *ListOrganizationExpandedResponseDataInner {
	this := ListOrganizationExpandedResponseDataInner{}
	this.Name = name
	this.Id = id
	this.OwnerId = ownerId
	return &this
}

// NewListOrganizationExpandedResponseDataInnerWithDefaults instantiates a new ListOrganizationExpandedResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrganizationExpandedResponseDataInnerWithDefaults() *ListOrganizationExpandedResponseDataInner {
	this := ListOrganizationExpandedResponseDataInner{}
	return &this
}

// GetName returns the Name field value
func (o *ListOrganizationExpandedResponseDataInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListOrganizationExpandedResponseDataInner) SetName(v string) {
	o.Name = v
}

// GetDefaultOrganizationAccess returns the DefaultOrganizationAccess field value if set, zero value otherwise.
func (o *ListOrganizationExpandedResponseDataInner) GetDefaultOrganizationAccess() []string {
	if o == nil || IsNil(o.DefaultOrganizationAccess) {
		var ret []string
		return ret
	}
	return o.DefaultOrganizationAccess
}

// GetDefaultOrganizationAccessOk returns a tuple with the DefaultOrganizationAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInner) GetDefaultOrganizationAccessOk() ([]string, bool) {
	if o == nil || IsNil(o.DefaultOrganizationAccess) {
		return nil, false
	}
	return o.DefaultOrganizationAccess, true
}

// HasDefaultOrganizationAccess returns a boolean if a field has been set.
func (o *ListOrganizationExpandedResponseDataInner) HasDefaultOrganizationAccess() bool {
	if o != nil && !IsNil(o.DefaultOrganizationAccess) {
		return true
	}

	return false
}

// SetDefaultOrganizationAccess gets a reference to the given []string and assigns it to the DefaultOrganizationAccess field.
func (o *ListOrganizationExpandedResponseDataInner) SetDefaultOrganizationAccess(v []string) {
	o.DefaultOrganizationAccess = v
}

// GetDefaultStackAccess returns the DefaultStackAccess field value if set, zero value otherwise.
func (o *ListOrganizationExpandedResponseDataInner) GetDefaultStackAccess() []string {
	if o == nil || IsNil(o.DefaultStackAccess) {
		var ret []string
		return ret
	}
	return o.DefaultStackAccess
}

// GetDefaultStackAccessOk returns a tuple with the DefaultStackAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInner) GetDefaultStackAccessOk() ([]string, bool) {
	if o == nil || IsNil(o.DefaultStackAccess) {
		return nil, false
	}
	return o.DefaultStackAccess, true
}

// HasDefaultStackAccess returns a boolean if a field has been set.
func (o *ListOrganizationExpandedResponseDataInner) HasDefaultStackAccess() bool {
	if o != nil && !IsNil(o.DefaultStackAccess) {
		return true
	}

	return false
}

// SetDefaultStackAccess gets a reference to the given []string and assigns it to the DefaultStackAccess field.
func (o *ListOrganizationExpandedResponseDataInner) SetDefaultStackAccess(v []string) {
	o.DefaultStackAccess = v
}

// GetId returns the Id field value
func (o *ListOrganizationExpandedResponseDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListOrganizationExpandedResponseDataInner) SetId(v string) {
	o.Id = v
}

// GetOwnerId returns the OwnerId field value
func (o *ListOrganizationExpandedResponseDataInner) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInner) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *ListOrganizationExpandedResponseDataInner) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetAvailableStacks returns the AvailableStacks field value if set, zero value otherwise.
func (o *ListOrganizationExpandedResponseDataInner) GetAvailableStacks() int32 {
	if o == nil || IsNil(o.AvailableStacks) {
		var ret int32
		return ret
	}
	return *o.AvailableStacks
}

// GetAvailableStacksOk returns a tuple with the AvailableStacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInner) GetAvailableStacksOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailableStacks) {
		return nil, false
	}
	return o.AvailableStacks, true
}

// HasAvailableStacks returns a boolean if a field has been set.
func (o *ListOrganizationExpandedResponseDataInner) HasAvailableStacks() bool {
	if o != nil && !IsNil(o.AvailableStacks) {
		return true
	}

	return false
}

// SetAvailableStacks gets a reference to the given int32 and assigns it to the AvailableStacks field.
func (o *ListOrganizationExpandedResponseDataInner) SetAvailableStacks(v int32) {
	o.AvailableStacks = &v
}

// GetAvailableSandboxes returns the AvailableSandboxes field value if set, zero value otherwise.
func (o *ListOrganizationExpandedResponseDataInner) GetAvailableSandboxes() int32 {
	if o == nil || IsNil(o.AvailableSandboxes) {
		var ret int32
		return ret
	}
	return *o.AvailableSandboxes
}

// GetAvailableSandboxesOk returns a tuple with the AvailableSandboxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInner) GetAvailableSandboxesOk() (*int32, bool) {
	if o == nil || IsNil(o.AvailableSandboxes) {
		return nil, false
	}
	return o.AvailableSandboxes, true
}

// HasAvailableSandboxes returns a boolean if a field has been set.
func (o *ListOrganizationExpandedResponseDataInner) HasAvailableSandboxes() bool {
	if o != nil && !IsNil(o.AvailableSandboxes) {
		return true
	}

	return false
}

// SetAvailableSandboxes gets a reference to the given int32 and assigns it to the AvailableSandboxes field.
func (o *ListOrganizationExpandedResponseDataInner) SetAvailableSandboxes(v int32) {
	o.AvailableSandboxes = &v
}

// GetTotalStacks returns the TotalStacks field value if set, zero value otherwise.
func (o *ListOrganizationExpandedResponseDataInner) GetTotalStacks() int32 {
	if o == nil || IsNil(o.TotalStacks) {
		var ret int32
		return ret
	}
	return *o.TotalStacks
}

// GetTotalStacksOk returns a tuple with the TotalStacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInner) GetTotalStacksOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalStacks) {
		return nil, false
	}
	return o.TotalStacks, true
}

// HasTotalStacks returns a boolean if a field has been set.
func (o *ListOrganizationExpandedResponseDataInner) HasTotalStacks() bool {
	if o != nil && !IsNil(o.TotalStacks) {
		return true
	}

	return false
}

// SetTotalStacks gets a reference to the given int32 and assigns it to the TotalStacks field.
func (o *ListOrganizationExpandedResponseDataInner) SetTotalStacks(v int32) {
	o.TotalStacks = &v
}

// GetTotalUsers returns the TotalUsers field value if set, zero value otherwise.
func (o *ListOrganizationExpandedResponseDataInner) GetTotalUsers() int32 {
	if o == nil || IsNil(o.TotalUsers) {
		var ret int32
		return ret
	}
	return *o.TotalUsers
}

// GetTotalUsersOk returns a tuple with the TotalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInner) GetTotalUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalUsers) {
		return nil, false
	}
	return o.TotalUsers, true
}

// HasTotalUsers returns a boolean if a field has been set.
func (o *ListOrganizationExpandedResponseDataInner) HasTotalUsers() bool {
	if o != nil && !IsNil(o.TotalUsers) {
		return true
	}

	return false
}

// SetTotalUsers gets a reference to the given int32 and assigns it to the TotalUsers field.
func (o *ListOrganizationExpandedResponseDataInner) SetTotalUsers(v int32) {
	o.TotalUsers = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ListOrganizationExpandedResponseDataInner) GetOwner() User {
	if o == nil || IsNil(o.Owner) {
		var ret User
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrganizationExpandedResponseDataInner) GetOwnerOk() (*User, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ListOrganizationExpandedResponseDataInner) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given User and assigns it to the Owner field.
func (o *ListOrganizationExpandedResponseDataInner) SetOwner(v User) {
	o.Owner = &v
}

func (o ListOrganizationExpandedResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrganizationExpandedResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.DefaultOrganizationAccess) {
		toSerialize["defaultOrganizationAccess"] = o.DefaultOrganizationAccess
	}
	if !IsNil(o.DefaultStackAccess) {
		toSerialize["defaultStackAccess"] = o.DefaultStackAccess
	}
	toSerialize["id"] = o.Id
	toSerialize["ownerId"] = o.OwnerId
	if !IsNil(o.AvailableStacks) {
		toSerialize["availableStacks"] = o.AvailableStacks
	}
	if !IsNil(o.AvailableSandboxes) {
		toSerialize["availableSandboxes"] = o.AvailableSandboxes
	}
	if !IsNil(o.TotalStacks) {
		toSerialize["totalStacks"] = o.TotalStacks
	}
	if !IsNil(o.TotalUsers) {
		toSerialize["totalUsers"] = o.TotalUsers
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableListOrganizationExpandedResponseDataInner struct {
	value *ListOrganizationExpandedResponseDataInner
	isSet bool
}

func (v NullableListOrganizationExpandedResponseDataInner) Get() *ListOrganizationExpandedResponseDataInner {
	return v.value
}

func (v *NullableListOrganizationExpandedResponseDataInner) Set(val *ListOrganizationExpandedResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrganizationExpandedResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrganizationExpandedResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrganizationExpandedResponseDataInner(val *ListOrganizationExpandedResponseDataInner) *NullableListOrganizationExpandedResponseDataInner {
	return &NullableListOrganizationExpandedResponseDataInner{value: val, isSet: true}
}

func (v NullableListOrganizationExpandedResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrganizationExpandedResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


