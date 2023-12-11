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

// checks if the UserAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAllOf{}

// UserAllOf struct for UserAllOf
type UserAllOf struct {
	// User roles
	Roles []string `json:"roles,omitempty"`
}

// NewUserAllOf instantiates a new UserAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAllOf() *UserAllOf {
	this := UserAllOf{}
	return &this
}

// NewUserAllOfWithDefaults instantiates a new UserAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAllOfWithDefaults() *UserAllOf {
	this := UserAllOf{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserAllOf) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAllOf) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserAllOf) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UserAllOf) SetRoles(v []string) {
	o.Roles = v
}

func (o UserAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableUserAllOf struct {
	value *UserAllOf
	isSet bool
}

func (v NullableUserAllOf) Get() *UserAllOf {
	return v.value
}

func (v *NullableUserAllOf) Set(val *UserAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAllOf(val *UserAllOf) *NullableUserAllOf {
	return &NullableUserAllOf{value: val, isSet: true}
}

func (v NullableUserAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


