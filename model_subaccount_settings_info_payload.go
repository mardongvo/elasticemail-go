/*
 * Elastic Email REST API
 *
 * This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    This is the documentation for REST API. If you’d like to read our legacy documentation regarding Web API v2 click <a target=\"_blank\" href=\"https://api.elasticemail.com/public/help\">here</a>.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>
 *
 * API version: 4.0.0
 * Contact: support@elasticemail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
)

// SubaccountSettingsInfoPayload SubAccount settings
type SubaccountSettingsInfoPayload struct {
	// Settings related to sending emails
	Email *SubaccountEmailSettingsPayload `json:"Email,omitempty"`
}

// NewSubaccountSettingsInfoPayload instantiates a new SubaccountSettingsInfoPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountSettingsInfoPayload() *SubaccountSettingsInfoPayload {
	this := SubaccountSettingsInfoPayload{}
	return &this
}

// NewSubaccountSettingsInfoPayloadWithDefaults instantiates a new SubaccountSettingsInfoPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountSettingsInfoPayloadWithDefaults() *SubaccountSettingsInfoPayload {
	this := SubaccountSettingsInfoPayload{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SubaccountSettingsInfoPayload) GetEmail() SubaccountEmailSettingsPayload {
	if o == nil || o.Email == nil {
		var ret SubaccountEmailSettingsPayload
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountSettingsInfoPayload) GetEmailOk() (*SubaccountEmailSettingsPayload, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SubaccountSettingsInfoPayload) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given SubaccountEmailSettingsPayload and assigns it to the Email field.
func (o *SubaccountSettingsInfoPayload) SetEmail(v SubaccountEmailSettingsPayload) {
	o.Email = &v
}

func (o SubaccountSettingsInfoPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableSubaccountSettingsInfoPayload struct {
	value *SubaccountSettingsInfoPayload
	isSet bool
}

func (v NullableSubaccountSettingsInfoPayload) Get() *SubaccountSettingsInfoPayload {
	return v.value
}

func (v *NullableSubaccountSettingsInfoPayload) Set(val *SubaccountSettingsInfoPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountSettingsInfoPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountSettingsInfoPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountSettingsInfoPayload(val *SubaccountSettingsInfoPayload) *NullableSubaccountSettingsInfoPayload {
	return &NullableSubaccountSettingsInfoPayload{value: val, isSet: true}
}

func (v NullableSubaccountSettingsInfoPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountSettingsInfoPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


