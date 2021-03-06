/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"encoding/json"
)
// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	// Form parameter enum test (string array)
	EnumFormStringArray *[]string `json:"enum_form_string_array,omitempty"`

	// Form parameter enum test (string)
	EnumFormString *string `json:"enum_form_string,omitempty"`

}

// GetEnumFormStringArray returns the EnumFormStringArray field if non-nil, zero value otherwise.
func (o *InlineObject2) GetEnumFormStringArray() []string {
	if o == nil || o.EnumFormStringArray == nil {
		var ret []string
		return ret
	}
	return *o.EnumFormStringArray
}

// GetEnumFormStringArrayOk returns a tuple with the EnumFormStringArray field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetEnumFormStringArrayOk() ([]string, bool) {
	if o == nil || o.EnumFormStringArray == nil {
		var ret []string
		return ret, false
	}
	return *o.EnumFormStringArray, true
}

// HasEnumFormStringArray returns a boolean if a field has been set.
func (o *InlineObject2) HasEnumFormStringArray() bool {
	if o != nil && o.EnumFormStringArray != nil {
		return true
	}

	return false
}

// SetEnumFormStringArray gets a reference to the given []string and assigns it to the EnumFormStringArray field.
func (o *InlineObject2) SetEnumFormStringArray(v []string) {
	o.EnumFormStringArray = &v
}

// GetEnumFormString returns the EnumFormString field if non-nil, zero value otherwise.
func (o *InlineObject2) GetEnumFormString() string {
	if o == nil || o.EnumFormString == nil {
		var ret string
		return ret
	}
	return *o.EnumFormString
}

// GetEnumFormStringOk returns a tuple with the EnumFormString field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetEnumFormStringOk() (string, bool) {
	if o == nil || o.EnumFormString == nil {
		var ret string
		return ret, false
	}
	return *o.EnumFormString, true
}

// HasEnumFormString returns a boolean if a field has been set.
func (o *InlineObject2) HasEnumFormString() bool {
	if o != nil && o.EnumFormString != nil {
		return true
	}

	return false
}

// SetEnumFormString gets a reference to the given string and assigns it to the EnumFormString field.
func (o *InlineObject2) SetEnumFormString(v string) {
	o.EnumFormString = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnumFormStringArray != nil {
		toSerialize["enum_form_string_array"] = o.EnumFormStringArray
	}
	if o.EnumFormString != nil {
		toSerialize["enum_form_string"] = o.EnumFormString
	}
	return json.Marshal(toSerialize)
}


