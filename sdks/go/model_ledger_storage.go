/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the LedgerStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LedgerStorage{}

// LedgerStorage struct for LedgerStorage
type LedgerStorage struct {
	Driver string `json:"driver"`
	Ledgers []string `json:"ledgers"`
}

// NewLedgerStorage instantiates a new LedgerStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerStorage(driver string, ledgers []string) *LedgerStorage {
	this := LedgerStorage{}
	this.Driver = driver
	this.Ledgers = ledgers
	return &this
}

// NewLedgerStorageWithDefaults instantiates a new LedgerStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerStorageWithDefaults() *LedgerStorage {
	this := LedgerStorage{}
	return &this
}

// GetDriver returns the Driver field value
func (o *LedgerStorage) GetDriver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Driver
}

// GetDriverOk returns a tuple with the Driver field value
// and a boolean to check if the value has been set.
func (o *LedgerStorage) GetDriverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Driver, true
}

// SetDriver sets field value
func (o *LedgerStorage) SetDriver(v string) {
	o.Driver = v
}

// GetLedgers returns the Ledgers field value
func (o *LedgerStorage) GetLedgers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ledgers
}

// GetLedgersOk returns a tuple with the Ledgers field value
// and a boolean to check if the value has been set.
func (o *LedgerStorage) GetLedgersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ledgers, true
}

// SetLedgers sets field value
func (o *LedgerStorage) SetLedgers(v []string) {
	o.Ledgers = v
}

func (o LedgerStorage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LedgerStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["driver"] = o.Driver
	toSerialize["ledgers"] = o.Ledgers
	return toSerialize, nil
}

type NullableLedgerStorage struct {
	value *LedgerStorage
	isSet bool
}

func (v NullableLedgerStorage) Get() *LedgerStorage {
	return v.value
}

func (v *NullableLedgerStorage) Set(val *LedgerStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerStorage(val *LedgerStorage) *NullableLedgerStorage {
	return &NullableLedgerStorage{value: val, isSet: true}
}

func (v NullableLedgerStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

