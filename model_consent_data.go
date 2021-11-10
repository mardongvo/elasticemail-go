/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    This is the documentation for REST API. If you’d like to read our legacy documentation regarding Web API v2 click <a target=\"_blank\" href=\"https://api.elasticemail.com/public/help\">here</a>.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"time"
)

// ConsentData struct for ConsentData
type ConsentData struct {
	// IP address of consent to send this contact(s) your email. If not provided your current public IP address is used for consent.
	ConsentIP *string `json:"ConsentIP,omitempty"`
	// Date of consent to send this contact(s) your email. If not provided current date is used for consent.
	ConsentDate NullableTime `json:"ConsentDate,omitempty"`
	ConsentTracking *ConsentTracking `json:"ConsentTracking,omitempty"`
}

// NewConsentData instantiates a new ConsentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsentData() *ConsentData {
	this := ConsentData{}
	var consentTracking ConsentTracking = UNKNOWN
	this.ConsentTracking = &consentTracking
	return &this
}

// NewConsentDataWithDefaults instantiates a new ConsentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsentDataWithDefaults() *ConsentData {
	this := ConsentData{}
	var consentTracking ConsentTracking = UNKNOWN
	this.ConsentTracking = &consentTracking
	return &this
}

// GetConsentIP returns the ConsentIP field value if set, zero value otherwise.
func (o *ConsentData) GetConsentIP() string {
	if o == nil || o.ConsentIP == nil {
		var ret string
		return ret
	}
	return *o.ConsentIP
}

// GetConsentIPOk returns a tuple with the ConsentIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentData) GetConsentIPOk() (*string, bool) {
	if o == nil || o.ConsentIP == nil {
		return nil, false
	}
	return o.ConsentIP, true
}

// HasConsentIP returns a boolean if a field has been set.
func (o *ConsentData) HasConsentIP() bool {
	if o != nil && o.ConsentIP != nil {
		return true
	}

	return false
}

// SetConsentIP gets a reference to the given string and assigns it to the ConsentIP field.
func (o *ConsentData) SetConsentIP(v string) {
	o.ConsentIP = &v
}

// GetConsentDate returns the ConsentDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsentData) GetConsentDate() time.Time {
	if o == nil || o.ConsentDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ConsentDate.Get()
}

// GetConsentDateOk returns a tuple with the ConsentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsentData) GetConsentDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConsentDate.Get(), o.ConsentDate.IsSet()
}

// HasConsentDate returns a boolean if a field has been set.
func (o *ConsentData) HasConsentDate() bool {
	if o != nil && o.ConsentDate.IsSet() {
		return true
	}

	return false
}

// SetConsentDate gets a reference to the given NullableTime and assigns it to the ConsentDate field.
func (o *ConsentData) SetConsentDate(v time.Time) {
	o.ConsentDate.Set(&v)
}
// SetConsentDateNil sets the value for ConsentDate to be an explicit nil
func (o *ConsentData) SetConsentDateNil() {
	o.ConsentDate.Set(nil)
}

// UnsetConsentDate ensures that no value is present for ConsentDate, not even an explicit nil
func (o *ConsentData) UnsetConsentDate() {
	o.ConsentDate.Unset()
}

// GetConsentTracking returns the ConsentTracking field value if set, zero value otherwise.
func (o *ConsentData) GetConsentTracking() ConsentTracking {
	if o == nil || o.ConsentTracking == nil {
		var ret ConsentTracking
		return ret
	}
	return *o.ConsentTracking
}

// GetConsentTrackingOk returns a tuple with the ConsentTracking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsentData) GetConsentTrackingOk() (*ConsentTracking, bool) {
	if o == nil || o.ConsentTracking == nil {
		return nil, false
	}
	return o.ConsentTracking, true
}

// HasConsentTracking returns a boolean if a field has been set.
func (o *ConsentData) HasConsentTracking() bool {
	if o != nil && o.ConsentTracking != nil {
		return true
	}

	return false
}

// SetConsentTracking gets a reference to the given ConsentTracking and assigns it to the ConsentTracking field.
func (o *ConsentData) SetConsentTracking(v ConsentTracking) {
	o.ConsentTracking = &v
}

func (o ConsentData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsentIP != nil {
		toSerialize["ConsentIP"] = o.ConsentIP
	}
	if o.ConsentDate.IsSet() {
		toSerialize["ConsentDate"] = o.ConsentDate.Get()
	}
	if o.ConsentTracking != nil {
		toSerialize["ConsentTracking"] = o.ConsentTracking
	}
	return json.Marshal(toSerialize)
}

type NullableConsentData struct {
	value *ConsentData
	isSet bool
}

func (v NullableConsentData) Get() *ConsentData {
	return v.value
}

func (v *NullableConsentData) Set(val *ConsentData) {
	v.value = val
	v.isSet = true
}

func (v NullableConsentData) IsSet() bool {
	return v.isSet
}

func (v *NullableConsentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsentData(val *ConsentData) *NullableConsentData {
	return &NullableConsentData{value: val, isSet: true}
}

func (v NullableConsentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


