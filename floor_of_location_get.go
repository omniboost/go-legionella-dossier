package legionella_dossier

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewFloorOfLocationGet() FloorOfLocationGet {
	r := FloorOfLocationGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type FloorOfLocationGet struct {
	client      *Client
	queryParams *FloorOfLocationGetQueryParams
	pathParams  *FloorOfLocationGetPathParams
	method      string
	headers     http.Header
	requestBody FloorOfLocationGetBody
}

func (r FloorOfLocationGet) NewQueryParams() *FloorOfLocationGetQueryParams {
	return &FloorOfLocationGetQueryParams{}
}

type FloorOfLocationGetQueryParams struct{}

func (p FloorOfLocationGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *FloorOfLocationGet) QueryParams() *FloorOfLocationGetQueryParams {
	return r.queryParams
}

func (r FloorOfLocationGet) NewPathParams() *FloorOfLocationGetPathParams {
	return &FloorOfLocationGetPathParams{}
}

type FloorOfLocationGetPathParams struct {
	LocationID int `schema:"location"`
}

func (p *FloorOfLocationGetPathParams) Params() map[string]string {
	return map[string]string{
		"location": strconv.Itoa(p.LocationID),
	}
}

func (r *FloorOfLocationGet) PathParams() *FloorOfLocationGetPathParams {
	return r.pathParams
}

func (r *FloorOfLocationGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *FloorOfLocationGet) SetMethod(method string) {
	r.method = method
}

func (r *FloorOfLocationGet) Method() string {
	return r.method
}

func (r FloorOfLocationGet) NewRequestBody() FloorOfLocationGetBody {
	return FloorOfLocationGetBody{}
}

type FloorOfLocationGetBody struct {
}

func (r *FloorOfLocationGet) RequestBody() *FloorOfLocationGetBody {
	return nil
}

func (r *FloorOfLocationGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *FloorOfLocationGet) SetRequestBody(body FloorOfLocationGetBody) {
	r.requestBody = body
}

func (r *FloorOfLocationGet) NewResponseBody() *FloorOfLocationGetResponseBody {
	return &FloorOfLocationGetResponseBody{}
}

type FloorOfLocationGetResponseBody Floors

func (r *FloorOfLocationGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/location/{{.location}}/floor", r.PathParams())
	return &u
}

func (r *FloorOfLocationGet) Do() (FloorOfLocationGetResponseBody, error) {
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
