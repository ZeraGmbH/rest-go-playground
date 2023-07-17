/*
 * SourceApi
 *
 * A Web API for controlling a source.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VeinGetResponse struct {

	// The desired return information formated as json
	ReturnInformation map[string]interface{} `json:"ReturnInformation,omitempty"`
}

// AssertVeinGetResponseRequired checks if the required fields are not zero-ed
func AssertVeinGetResponseRequired(obj VeinGetResponse) error {
	return nil
}

// AssertRecurseVeinGetResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of VeinGetResponse (e.g. [][]VeinGetResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseVeinGetResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aVeinGetResponse, ok := obj.(VeinGetResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertVeinGetResponseRequired(aVeinGetResponse)
	})
}
