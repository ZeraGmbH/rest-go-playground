/*
 * SourceApi
 *
 * A Web API for controlling a source.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"
)

// VeinApiService is a service that implements the logic for the VeinApiServicer
// This service should implement the business logic for every endpoint for the VeinApi API.
// Include any external packages or services that will be required by this service.
type VeinApiService struct {
}

// NewVeinApiService creates a default api service
func NewVeinApiService() VeinApiServicer {
	return &VeinApiService{}
}

// ApiV1VeinGetInfoPost - Gets information from Vein.
func (s *VeinApiService) ApiV1VeinGetInfoPost(ctx context.Context, veinGet VeinGet) (ImplResponse, error) {
	// TODO - update ApiV1VeinGetInfoPost with the required logic for this service method.
	// Add api_vein_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	var ret VeinGetResponse
	ret.ReturnInformation = make(map[string]interface{})
	ret.ReturnInformation["ReceivedEntityID"] = veinGet.EntityID
	ret.ReturnInformation["ReceivedComponentName"] = veinGet.ComponentName
	ret.ReturnInformation["ReceivedMiscInfo"] = veinGet.MiscFieldForInfo

	//TODO: Uncomment the next line to return response Response(200, VeinGetResponse{}) or use other options such as http.Ok ...
	return Response(200, ret), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(422, {}) or use other options such as http.Ok ...
	//return Response(422, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	//return Response(http.StatusNotImplemented, nil), errors.New("ApiV1VeinGetInfoPost method not implemented")
}

// ApiV1VeinSetInfoPost - Sets information to Vein.
func (s *VeinApiService) ApiV1VeinSetInfoPost(ctx context.Context, veinSet VeinSet) (ImplResponse, error) {
	// TODO - update ApiV1VeinSetInfoPost with the required logic for this service method.
	// Add api_vein_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ProblemDetails{}) or use other options such as http.Ok ...
	//return Response(200, ProblemDetails{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(422, {}) or use other options such as http.Ok ...
	//return Response(422, nil),nil

	//TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	//return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ApiV1VeinSetInfoPost method not implemented")
}
