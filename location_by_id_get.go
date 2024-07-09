package legionella_dossier

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewLocationByIDGet() LocationByIDGet {
	r := LocationByIDGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type LocationByIDGet struct {
	client      *Client
	queryParams *LocationByIDGetQueryParams
	pathParams  *LocationByIDGetPathParams
	method      string
	headers     http.Header
	requestBody LocationByIDGetBody
}

func (r LocationByIDGet) NewQueryParams() *LocationByIDGetQueryParams {
	return &LocationByIDGetQueryParams{}
}

type LocationByIDGetQueryParams struct {
}

func (p LocationByIDGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *LocationByIDGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r LocationByIDGet) NewPathParams() *LocationByIDGetPathParams {
	return &LocationByIDGetPathParams{}
}

type LocationByIDGetPathParams struct {
	LocationID int `schema:"location"`
}

func (p *LocationByIDGetPathParams) Params() map[string]string {
	return map[string]string{
		"location": strconv.Itoa(p.LocationID),
	}
}

func (r *LocationByIDGet) PathParams() *LocationByIDGetPathParams {
	return r.pathParams
}

func (r *LocationByIDGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *LocationByIDGet) SetMethod(method string) {
	r.method = method
}

func (r *LocationByIDGet) Method() string {
	return r.method
}

func (r LocationByIDGet) NewRequestBody() LocationByIDGetBody {
	return LocationByIDGetBody{}
}

type LocationByIDGetBody struct {
}

func (r *LocationByIDGet) RequestBody() *LocationByIDGetBody {
	return nil
}

func (r *LocationByIDGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *LocationByIDGet) SetRequestBody(body LocationByIDGetBody) {
	r.requestBody = body
}

func (r *LocationByIDGet) NewResponseBody() *LocationByIDGetResponseBody {
	return &LocationByIDGetResponseBody{}
}

type LocationByIDGetResponseBody Location

func (r *LocationByIDGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/locations/{{.location}}", r.PathParams())
	return &u
}

func (r *LocationByIDGet) Do() (LocationByIDGetResponseBody, error) {
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
