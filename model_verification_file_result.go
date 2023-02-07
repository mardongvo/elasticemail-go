/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"time"
)

// VerificationFileResult Simple verification file result info
type VerificationFileResult struct {
	// Identifier of this verification result
	VerificationID *string `json:"VerificationID,omitempty"`
	// Origin file name
	Filename *string `json:"Filename,omitempty"`
	VerificationStatus *VerificationStatus `json:"VerificationStatus,omitempty"`
	FileUploadResult *FileUploadResult `json:"FileUploadResult,omitempty"`
	// Date of creation in YYYY-MM-DDThh:ii:ss format
	DateAdded *time.Time `json:"DateAdded,omitempty"`
	// Origin file extension
	Source *string `json:"Source,omitempty"`
}

// NewVerificationFileResult instantiates a new VerificationFileResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationFileResult() *VerificationFileResult {
	this := VerificationFileResult{}
	var verificationStatus VerificationStatus = VerificationStatusPROCESSING
	this.VerificationStatus = &verificationStatus
	return &this
}

// NewVerificationFileResultWithDefaults instantiates a new VerificationFileResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationFileResultWithDefaults() *VerificationFileResult {
	this := VerificationFileResult{}
	var verificationStatus VerificationStatus = VerificationStatusPROCESSING
	this.VerificationStatus = &verificationStatus
	return &this
}

// GetVerificationID returns the VerificationID field value if set, zero value otherwise.
func (o *VerificationFileResult) GetVerificationID() string {
	if o == nil || isNil(o.VerificationID) {
		var ret string
		return ret
	}
	return *o.VerificationID
}

// GetVerificationIDOk returns a tuple with the VerificationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFileResult) GetVerificationIDOk() (*string, bool) {
	if o == nil || isNil(o.VerificationID) {
    return nil, false
	}
	return o.VerificationID, true
}

// HasVerificationID returns a boolean if a field has been set.
func (o *VerificationFileResult) HasVerificationID() bool {
	if o != nil && !isNil(o.VerificationID) {
		return true
	}

	return false
}

// SetVerificationID gets a reference to the given string and assigns it to the VerificationID field.
func (o *VerificationFileResult) SetVerificationID(v string) {
	o.VerificationID = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *VerificationFileResult) GetFilename() string {
	if o == nil || isNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFileResult) GetFilenameOk() (*string, bool) {
	if o == nil || isNil(o.Filename) {
    return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *VerificationFileResult) HasFilename() bool {
	if o != nil && !isNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *VerificationFileResult) SetFilename(v string) {
	o.Filename = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *VerificationFileResult) GetVerificationStatus() VerificationStatus {
	if o == nil || isNil(o.VerificationStatus) {
		var ret VerificationStatus
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFileResult) GetVerificationStatusOk() (*VerificationStatus, bool) {
	if o == nil || isNil(o.VerificationStatus) {
    return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *VerificationFileResult) HasVerificationStatus() bool {
	if o != nil && !isNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given VerificationStatus and assigns it to the VerificationStatus field.
func (o *VerificationFileResult) SetVerificationStatus(v VerificationStatus) {
	o.VerificationStatus = &v
}

// GetFileUploadResult returns the FileUploadResult field value if set, zero value otherwise.
func (o *VerificationFileResult) GetFileUploadResult() FileUploadResult {
	if o == nil || isNil(o.FileUploadResult) {
		var ret FileUploadResult
		return ret
	}
	return *o.FileUploadResult
}

// GetFileUploadResultOk returns a tuple with the FileUploadResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFileResult) GetFileUploadResultOk() (*FileUploadResult, bool) {
	if o == nil || isNil(o.FileUploadResult) {
    return nil, false
	}
	return o.FileUploadResult, true
}

// HasFileUploadResult returns a boolean if a field has been set.
func (o *VerificationFileResult) HasFileUploadResult() bool {
	if o != nil && !isNil(o.FileUploadResult) {
		return true
	}

	return false
}

// SetFileUploadResult gets a reference to the given FileUploadResult and assigns it to the FileUploadResult field.
func (o *VerificationFileResult) SetFileUploadResult(v FileUploadResult) {
	o.FileUploadResult = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *VerificationFileResult) GetDateAdded() time.Time {
	if o == nil || isNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFileResult) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DateAdded) {
    return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *VerificationFileResult) HasDateAdded() bool {
	if o != nil && !isNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *VerificationFileResult) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *VerificationFileResult) GetSource() string {
	if o == nil || isNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationFileResult) GetSourceOk() (*string, bool) {
	if o == nil || isNil(o.Source) {
    return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *VerificationFileResult) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *VerificationFileResult) SetSource(v string) {
	o.Source = &v
}

func (o VerificationFileResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.VerificationID) {
		toSerialize["VerificationID"] = o.VerificationID
	}
	if !isNil(o.Filename) {
		toSerialize["Filename"] = o.Filename
	}
	if !isNil(o.VerificationStatus) {
		toSerialize["VerificationStatus"] = o.VerificationStatus
	}
	if !isNil(o.FileUploadResult) {
		toSerialize["FileUploadResult"] = o.FileUploadResult
	}
	if !isNil(o.DateAdded) {
		toSerialize["DateAdded"] = o.DateAdded
	}
	if !isNil(o.Source) {
		toSerialize["Source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableVerificationFileResult struct {
	value *VerificationFileResult
	isSet bool
}

func (v NullableVerificationFileResult) Get() *VerificationFileResult {
	return v.value
}

func (v *NullableVerificationFileResult) Set(val *VerificationFileResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationFileResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationFileResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationFileResult(val *VerificationFileResult) *NullableVerificationFileResult {
	return &NullableVerificationFileResult{value: val, isSet: true}
}

func (v NullableVerificationFileResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationFileResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


