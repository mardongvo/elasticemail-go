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

// SplitOptions Optional A/X split campaign options
type SplitOptions struct {
	// Type of results by which to determine the winner template (content)
	OptimizeFor *SplitOptimizationType `json:"OptimizeFor,omitempty"`
	// For how long should the results be measured until determining the winner template (content)
	OptimizePeriodMinutes *int32 `json:"OptimizePeriodMinutes,omitempty"`
}

// NewSplitOptions instantiates a new SplitOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitOptions() *SplitOptions {
	this := SplitOptions{}
	return &this
}

// NewSplitOptionsWithDefaults instantiates a new SplitOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitOptionsWithDefaults() *SplitOptions {
	this := SplitOptions{}
	return &this
}

// GetOptimizeFor returns the OptimizeFor field value if set, zero value otherwise.
func (o *SplitOptions) GetOptimizeFor() SplitOptimizationType {
	if o == nil || o.OptimizeFor == nil {
		var ret SplitOptimizationType
		return ret
	}
	return *o.OptimizeFor
}

// GetOptimizeForOk returns a tuple with the OptimizeFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitOptions) GetOptimizeForOk() (*SplitOptimizationType, bool) {
	if o == nil || o.OptimizeFor == nil {
		return nil, false
	}
	return o.OptimizeFor, true
}

// HasOptimizeFor returns a boolean if a field has been set.
func (o *SplitOptions) HasOptimizeFor() bool {
	if o != nil && o.OptimizeFor != nil {
		return true
	}

	return false
}

// SetOptimizeFor gets a reference to the given SplitOptimizationType and assigns it to the OptimizeFor field.
func (o *SplitOptions) SetOptimizeFor(v SplitOptimizationType) {
	o.OptimizeFor = &v
}

// GetOptimizePeriodMinutes returns the OptimizePeriodMinutes field value if set, zero value otherwise.
func (o *SplitOptions) GetOptimizePeriodMinutes() int32 {
	if o == nil || o.OptimizePeriodMinutes == nil {
		var ret int32
		return ret
	}
	return *o.OptimizePeriodMinutes
}

// GetOptimizePeriodMinutesOk returns a tuple with the OptimizePeriodMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitOptions) GetOptimizePeriodMinutesOk() (*int32, bool) {
	if o == nil || o.OptimizePeriodMinutes == nil {
		return nil, false
	}
	return o.OptimizePeriodMinutes, true
}

// HasOptimizePeriodMinutes returns a boolean if a field has been set.
func (o *SplitOptions) HasOptimizePeriodMinutes() bool {
	if o != nil && o.OptimizePeriodMinutes != nil {
		return true
	}

	return false
}

// SetOptimizePeriodMinutes gets a reference to the given int32 and assigns it to the OptimizePeriodMinutes field.
func (o *SplitOptions) SetOptimizePeriodMinutes(v int32) {
	o.OptimizePeriodMinutes = &v
}

func (o SplitOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OptimizeFor != nil {
		toSerialize["OptimizeFor"] = o.OptimizeFor
	}
	if o.OptimizePeriodMinutes != nil {
		toSerialize["OptimizePeriodMinutes"] = o.OptimizePeriodMinutes
	}
	return json.Marshal(toSerialize)
}

type NullableSplitOptions struct {
	value *SplitOptions
	isSet bool
}

func (v NullableSplitOptions) Get() *SplitOptions {
	return v.value
}

func (v *NullableSplitOptions) Set(val *SplitOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitOptions(val *SplitOptions) *NullableSplitOptions {
	return &NullableSplitOptions{value: val, isSet: true}
}

func (v NullableSplitOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


