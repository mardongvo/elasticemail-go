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

// TransactionalRecipient List of transactional recipients
type TransactionalRecipient struct {
	// List of recipients (visible to others)
	To *[]string `json:"To,omitempty"`
	// List of Carbon Copy recipients (visible to others)
	CC *[]string `json:"CC,omitempty"`
	// List of Blind Carbon Copy recipients (hidden from other recipients)
	BCC *[]string `json:"BCC,omitempty"`
}

// NewTransactionalRecipient instantiates a new TransactionalRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionalRecipient() *TransactionalRecipient {
	this := TransactionalRecipient{}
	return &this
}

// NewTransactionalRecipientWithDefaults instantiates a new TransactionalRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionalRecipientWithDefaults() *TransactionalRecipient {
	this := TransactionalRecipient{}
	return &this
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *TransactionalRecipient) GetTo() []string {
	if o == nil || o.To == nil {
		var ret []string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionalRecipient) GetToOk() (*[]string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *TransactionalRecipient) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given []string and assigns it to the To field.
func (o *TransactionalRecipient) SetTo(v []string) {
	o.To = &v
}

// GetCC returns the CC field value if set, zero value otherwise.
func (o *TransactionalRecipient) GetCC() []string {
	if o == nil || o.CC == nil {
		var ret []string
		return ret
	}
	return *o.CC
}

// GetCCOk returns a tuple with the CC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionalRecipient) GetCCOk() (*[]string, bool) {
	if o == nil || o.CC == nil {
		return nil, false
	}
	return o.CC, true
}

// HasCC returns a boolean if a field has been set.
func (o *TransactionalRecipient) HasCC() bool {
	if o != nil && o.CC != nil {
		return true
	}

	return false
}

// SetCC gets a reference to the given []string and assigns it to the CC field.
func (o *TransactionalRecipient) SetCC(v []string) {
	o.CC = &v
}

// GetBCC returns the BCC field value if set, zero value otherwise.
func (o *TransactionalRecipient) GetBCC() []string {
	if o == nil || o.BCC == nil {
		var ret []string
		return ret
	}
	return *o.BCC
}

// GetBCCOk returns a tuple with the BCC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionalRecipient) GetBCCOk() (*[]string, bool) {
	if o == nil || o.BCC == nil {
		return nil, false
	}
	return o.BCC, true
}

// HasBCC returns a boolean if a field has been set.
func (o *TransactionalRecipient) HasBCC() bool {
	if o != nil && o.BCC != nil {
		return true
	}

	return false
}

// SetBCC gets a reference to the given []string and assigns it to the BCC field.
func (o *TransactionalRecipient) SetBCC(v []string) {
	o.BCC = &v
}

func (o TransactionalRecipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.To != nil {
		toSerialize["To"] = o.To
	}
	if o.CC != nil {
		toSerialize["CC"] = o.CC
	}
	if o.BCC != nil {
		toSerialize["BCC"] = o.BCC
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionalRecipient struct {
	value *TransactionalRecipient
	isSet bool
}

func (v NullableTransactionalRecipient) Get() *TransactionalRecipient {
	return v.value
}

func (v *NullableTransactionalRecipient) Set(val *TransactionalRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionalRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionalRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionalRecipient(val *TransactionalRecipient) *NullableTransactionalRecipient {
	return &NullableTransactionalRecipient{value: val, isSet: true}
}

func (v NullableTransactionalRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionalRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


