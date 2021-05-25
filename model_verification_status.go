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
	"fmt"
)

// VerificationStatus the model 'VerificationStatus'
type VerificationStatus string

// List of VerificationStatus
const (
	PROCESSING VerificationStatus = "Processing"
	READY VerificationStatus = "Ready"
	EXPIRED VerificationStatus = "Expired"
	VERIFIED VerificationStatus = "Verified"
	ERROR VerificationStatus = "Error"
)

func (v *VerificationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerificationStatus(value)
	for _, existing := range []VerificationStatus{ "Processing", "Ready", "Expired", "Verified", "Error",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VerificationStatus", value)
}

// Ptr returns reference to VerificationStatus value
func (v VerificationStatus) Ptr() *VerificationStatus {
	return &v
}

type NullableVerificationStatus struct {
	value *VerificationStatus
	isSet bool
}

func (v NullableVerificationStatus) Get() *VerificationStatus {
	return v.value
}

func (v *NullableVerificationStatus) Set(val *VerificationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationStatus(val *VerificationStatus) *NullableVerificationStatus {
	return &NullableVerificationStatus{value: val, isSet: true}
}

func (v NullableVerificationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

