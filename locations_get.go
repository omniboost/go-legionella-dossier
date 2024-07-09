package legionella_dossier

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewLocationsGet() LocationsGet {
	r := LocationsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type LocationsGet struct {
	client      *Client
	queryParams *LocationsGetQueryParams
	pathParams  *LocationsGetPathParams
	method      string
	headers     http.Header
	requestBody LocationsGetBody
}

func (r LocationsGet) NewQueryParams() *LocationsGetQueryParams {
	return &LocationsGetQueryParams{}
}

type LocationsGetQueryParams struct {
}

func (p LocationsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LocationsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r LocationsGet) NewPathParams() *LocationsGetPathParams {
	return &LocationsGetPathParams{}
}

type LocationsGetPathParams struct{}

func (p *LocationsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LocationsGet) PathParams() *LocationsGetPathParams {
	return r.pathParams
}

func (r *LocationsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *LocationsGet) SetMethod(method string) {
	r.method = method
}

func (r *LocationsGet) Method() string {
	return r.method
}

func (r LocationsGet) NewRequestBody() LocationsGetBody {
	return LocationsGetBody{}
}

type LocationsGetBody struct {
}

func (r *LocationsGet) RequestBody() *LocationsGetBody {
	return nil
}

func (r *LocationsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *LocationsGet) SetRequestBody(body LocationsGetBody) {
	r.requestBody = body
}

func (r *LocationsGet) NewResponseBody() *LocationsGetResponseBody {
	return &LocationsGetResponseBody{}
}

type LocationsGetResponseBody Locations

func (r *LocationsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/locations", r.PathParams())
	return &u
}

func (r *LocationsGet) Do() (LocationsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
