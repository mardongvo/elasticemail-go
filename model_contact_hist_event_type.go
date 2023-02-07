/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// ContactHistEventType the model 'ContactHistEventType'
type ContactHistEventType string

// List of ContactHistEventType
const (
	ContactHistEventTypeOPENED ContactHistEventType = "Opened"
	ContactHistEventTypeCLICKED ContactHistEventType = "Clicked"
	ContactHistEventTypeBOUNCED ContactHistEventType = "Bounced"
	ContactHistEventTypeUNSUBSCRIBED ContactHistEventType = "Unsubscribed"
	COMPLAINED ContactHistEventType = "Complained"
	ACTIVATED ContactHistEventType = "Activated"
	TRANSACTIONAL_UNSUBSCRIBED ContactHistEventType = "TransactionalUnsubscribed"
	MANUAL_STATUS_CHANGE ContactHistEventType = "ManualStatusChange"
	MANUAL_CONSENT_TRACKING_CHANGE ContactHistEventType = "ManualConsentTrackingChange"
	ACTIVATION_SENT ContactHistEventType = "ActivationSent"
	JOURNEY_STARTED ContactHistEventType = "JourneyStarted"
	JOURNEY_STEP_PROCESSED ContactHistEventType = "JourneyStepProcessed"
	JOURNEY_FINISHED ContactHistEventType = "JourneyFinished"
	DELETED ContactHistEventType = "Deleted"
)

// All allowed values of ContactHistEventType enum
var AllowedContactHistEventTypeEnumValues = []ContactHistEventType{
	"Opened",
	"Clicked",
	"Bounced",
	"Unsubscribed",
	"Complained",
	"Activated",
	"TransactionalUnsubscribed",
	"ManualStatusChange",
	"ManualConsentTrackingChange",
	"ActivationSent",
	"JourneyStarted",
	"JourneyStepProcessed",
	"JourneyFinished",
	"Deleted",
}

func (v *ContactHistEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContactHistEventType(value)
	for _, existing := range AllowedContactHistEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContactHistEventType", value)
}

// NewContactHistEventTypeFromValue returns a pointer to a valid ContactHistEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContactHistEventTypeFromValue(v string) (*ContactHistEventType, error) {
	ev := ContactHistEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContactHistEventType: valid values are %v", v, AllowedContactHistEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContactHistEventType) IsValid() bool {
	for _, existing := range AllowedContactHistEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContactHistEventType value
func (v ContactHistEventType) Ptr() *ContactHistEventType {
	return &v
}

type NullableContactHistEventType struct {
	value *ContactHistEventType
	isSet bool
}

func (v NullableContactHistEventType) Get() *ContactHistEventType {
	return v.value
}

func (v *NullableContactHistEventType) Set(val *ContactHistEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableContactHistEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableContactHistEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactHistEventType(val *ContactHistEventType) *NullableContactHistEventType {
	return &NullableContactHistEventType{value: val, isSet: true}
}

func (v NullableContactHistEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactHistEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

