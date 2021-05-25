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

// SubaccountEmailSettings Settings related to sending emails
type SubaccountEmailSettings struct {
	// Amount of credits added to Account automatically
	MonthlyRefillCredits *int32 `json:"MonthlyRefillCredits,omitempty"`
	// True, if Account needs credits to send emails. Otherwise, false
	RequiresEmailCredits *bool `json:"RequiresEmailCredits,omitempty"`
	// Maximum size of email including attachments in MB's
	EmailSizeLimit *int32 `json:"EmailSizeLimit,omitempty"`
	// Amount of emails Account can send daily
	DailySendLimit *int32 `json:"DailySendLimit,omitempty"`
	// Maximum number of contacts the Account can have. 0 means that parent account's limit is used.
	MaxContacts *int32 `json:"MaxContacts,omitempty"`
	// Can the SubAccount purchase Private IP for themselves
	EnablePrivateIPPurchase *bool `json:"EnablePrivateIPPurchase,omitempty"`
	// Name of your custom IP Pool to be used in the sending process
	PoolName *string `json:"PoolName,omitempty"`
}

// NewSubaccountEmailSettings instantiates a new SubaccountEmailSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountEmailSettings() *SubaccountEmailSettings {
	this := SubaccountEmailSettings{}
	return &this
}

// NewSubaccountEmailSettingsWithDefaults instantiates a new SubaccountEmailSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountEmailSettingsWithDefaults() *SubaccountEmailSettings {
	this := SubaccountEmailSettings{}
	return &this
}

// GetMonthlyRefillCredits returns the MonthlyRefillCredits field value if set, zero value otherwise.
func (o *SubaccountEmailSettings) GetMonthlyRefillCredits() int32 {
	if o == nil || o.MonthlyRefillCredits == nil {
		var ret int32
		return ret
	}
	return *o.MonthlyRefillCredits
}

// GetMonthlyRefillCreditsOk returns a tuple with the MonthlyRefillCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettings) GetMonthlyRefillCreditsOk() (*int32, bool) {
	if o == nil || o.MonthlyRefillCredits == nil {
		return nil, false
	}
	return o.MonthlyRefillCredits, true
}

// HasMonthlyRefillCredits returns a boolean if a field has been set.
func (o *SubaccountEmailSettings) HasMonthlyRefillCredits() bool {
	if o != nil && o.MonthlyRefillCredits != nil {
		return true
	}

	return false
}

// SetMonthlyRefillCredits gets a reference to the given int32 and assigns it to the MonthlyRefillCredits field.
func (o *SubaccountEmailSettings) SetMonthlyRefillCredits(v int32) {
	o.MonthlyRefillCredits = &v
}

// GetRequiresEmailCredits returns the RequiresEmailCredits field value if set, zero value otherwise.
func (o *SubaccountEmailSettings) GetRequiresEmailCredits() bool {
	if o == nil || o.RequiresEmailCredits == nil {
		var ret bool
		return ret
	}
	return *o.RequiresEmailCredits
}

// GetRequiresEmailCreditsOk returns a tuple with the RequiresEmailCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettings) GetRequiresEmailCreditsOk() (*bool, bool) {
	if o == nil || o.RequiresEmailCredits == nil {
		return nil, false
	}
	return o.RequiresEmailCredits, true
}

// HasRequiresEmailCredits returns a boolean if a field has been set.
func (o *SubaccountEmailSettings) HasRequiresEmailCredits() bool {
	if o != nil && o.RequiresEmailCredits != nil {
		return true
	}

	return false
}

// SetRequiresEmailCredits gets a reference to the given bool and assigns it to the RequiresEmailCredits field.
func (o *SubaccountEmailSettings) SetRequiresEmailCredits(v bool) {
	o.RequiresEmailCredits = &v
}

// GetEmailSizeLimit returns the EmailSizeLimit field value if set, zero value otherwise.
func (o *SubaccountEmailSettings) GetEmailSizeLimit() int32 {
	if o == nil || o.EmailSizeLimit == nil {
		var ret int32
		return ret
	}
	return *o.EmailSizeLimit
}

// GetEmailSizeLimitOk returns a tuple with the EmailSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettings) GetEmailSizeLimitOk() (*int32, bool) {
	if o == nil || o.EmailSizeLimit == nil {
		return nil, false
	}
	return o.EmailSizeLimit, true
}

// HasEmailSizeLimit returns a boolean if a field has been set.
func (o *SubaccountEmailSettings) HasEmailSizeLimit() bool {
	if o != nil && o.EmailSizeLimit != nil {
		return true
	}

	return false
}

// SetEmailSizeLimit gets a reference to the given int32 and assigns it to the EmailSizeLimit field.
func (o *SubaccountEmailSettings) SetEmailSizeLimit(v int32) {
	o.EmailSizeLimit = &v
}

// GetDailySendLimit returns the DailySendLimit field value if set, zero value otherwise.
func (o *SubaccountEmailSettings) GetDailySendLimit() int32 {
	if o == nil || o.DailySendLimit == nil {
		var ret int32
		return ret
	}
	return *o.DailySendLimit
}

// GetDailySendLimitOk returns a tuple with the DailySendLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettings) GetDailySendLimitOk() (*int32, bool) {
	if o == nil || o.DailySendLimit == nil {
		return nil, false
	}
	return o.DailySendLimit, true
}

// HasDailySendLimit returns a boolean if a field has been set.
func (o *SubaccountEmailSettings) HasDailySendLimit() bool {
	if o != nil && o.DailySendLimit != nil {
		return true
	}

	return false
}

// SetDailySendLimit gets a reference to the given int32 and assigns it to the DailySendLimit field.
func (o *SubaccountEmailSettings) SetDailySendLimit(v int32) {
	o.DailySendLimit = &v
}

// GetMaxContacts returns the MaxContacts field value if set, zero value otherwise.
func (o *SubaccountEmailSettings) GetMaxContacts() int32 {
	if o == nil || o.MaxContacts == nil {
		var ret int32
		return ret
	}
	return *o.MaxContacts
}

// GetMaxContactsOk returns a tuple with the MaxContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettings) GetMaxContactsOk() (*int32, bool) {
	if o == nil || o.MaxContacts == nil {
		return nil, false
	}
	return o.MaxContacts, true
}

// HasMaxContacts returns a boolean if a field has been set.
func (o *SubaccountEmailSettings) HasMaxContacts() bool {
	if o != nil && o.MaxContacts != nil {
		return true
	}

	return false
}

// SetMaxContacts gets a reference to the given int32 and assigns it to the MaxContacts field.
func (o *SubaccountEmailSettings) SetMaxContacts(v int32) {
	o.MaxContacts = &v
}

// GetEnablePrivateIPPurchase returns the EnablePrivateIPPurchase field value if set, zero value otherwise.
func (o *SubaccountEmailSettings) GetEnablePrivateIPPurchase() bool {
	if o == nil || o.EnablePrivateIPPurchase == nil {
		var ret bool
		return ret
	}
	return *o.EnablePrivateIPPurchase
}

// GetEnablePrivateIPPurchaseOk returns a tuple with the EnablePrivateIPPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettings) GetEnablePrivateIPPurchaseOk() (*bool, bool) {
	if o == nil || o.EnablePrivateIPPurchase == nil {
		return nil, false
	}
	return o.EnablePrivateIPPurchase, true
}

// HasEnablePrivateIPPurchase returns a boolean if a field has been set.
func (o *SubaccountEmailSettings) HasEnablePrivateIPPurchase() bool {
	if o != nil && o.EnablePrivateIPPurchase != nil {
		return true
	}

	return false
}

// SetEnablePrivateIPPurchase gets a reference to the given bool and assigns it to the EnablePrivateIPPurchase field.
func (o *SubaccountEmailSettings) SetEnablePrivateIPPurchase(v bool) {
	o.EnablePrivateIPPurchase = &v
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *SubaccountEmailSettings) GetPoolName() string {
	if o == nil || o.PoolName == nil {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettings) GetPoolNameOk() (*string, bool) {
	if o == nil || o.PoolName == nil {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *SubaccountEmailSettings) HasPoolName() bool {
	if o != nil && o.PoolName != nil {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *SubaccountEmailSettings) SetPoolName(v string) {
	o.PoolName = &v
}

func (o SubaccountEmailSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MonthlyRefillCredits != nil {
		toSerialize["MonthlyRefillCredits"] = o.MonthlyRefillCredits
	}
	if o.RequiresEmailCredits != nil {
		toSerialize["RequiresEmailCredits"] = o.RequiresEmailCredits
	}
	if o.EmailSizeLimit != nil {
		toSerialize["EmailSizeLimit"] = o.EmailSizeLimit
	}
	if o.DailySendLimit != nil {
		toSerialize["DailySendLimit"] = o.DailySendLimit
	}
	if o.MaxContacts != nil {
		toSerialize["MaxContacts"] = o.MaxContacts
	}
	if o.EnablePrivateIPPurchase != nil {
		toSerialize["EnablePrivateIPPurchase"] = o.EnablePrivateIPPurchase
	}
	if o.PoolName != nil {
		toSerialize["PoolName"] = o.PoolName
	}
	return json.Marshal(toSerialize)
}

type NullableSubaccountEmailSettings struct {
	value *SubaccountEmailSettings
	isSet bool
}

func (v NullableSubaccountEmailSettings) Get() *SubaccountEmailSettings {
	return v.value
}

func (v *NullableSubaccountEmailSettings) Set(val *SubaccountEmailSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountEmailSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountEmailSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountEmailSettings(val *SubaccountEmailSettings) *NullableSubaccountEmailSettings {
	return &NullableSubaccountEmailSettings{value: val, isSet: true}
}

func (v NullableSubaccountEmailSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountEmailSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


