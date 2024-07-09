package legionella_dossier

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewFloorRoomsOfLocationGet() FloorRoomsOfLocationGet {
	r := FloorRoomsOfLocationGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type FloorRoomsOfLocationGet struct {
	client      *Client
	queryParams *FloorRoomsOfLocationGetQueryParams
	pathParams  *FloorRoomsOfLocationGetPathParams
	method      string
	headers     http.Header
	requestBody FloorRoomsOfLocationGetBody
}

func (r FloorRoomsOfLocationGet) NewQueryParams() *FloorRoomsOfLocationGetQueryParams {
	return &FloorRoomsOfLocationGetQueryParams{}
}

type FloorRoomsOfLocationGetQueryParams struct{}

func (p FloorRoomsOfLocationGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *FloorRoomsOfLocationGet) QueryParams() *FloorRoomsOfLocationGetQueryParams {
	return r.queryParams
}

func (r FloorRoomsOfLocationGet) NewPathParams() *FloorRoomsOfLocationGetPathParams {
	return &FloorRoomsOfLocationGetPathParams{}
}

type FloorRoomsOfLocationGetPathParams struct {
	LocationID int `schema:"location"`
	FloorID    int `schema:"floor"`
}

func (p *FloorRoomsOfLocationGetPathParams) Params() map[string]string {
	return map[string]string{
		"location": strconv.Itoa(p.LocationID),
		"floor":    strconv.Itoa(p.FloorID),
	}
}

func (r *FloorRoomsOfLocationGet) PathParams() *FloorRoomsOfLocationGetPathParams {
	return r.pathParams
}

func (r *FloorRoomsOfLocationGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *FloorRoomsOfLocationGet) SetMethod(method string) {
	r.method = method
}

func (r *FloorRoomsOfLocationGet) Method() string {
	return r.method
}

func (r FloorRoomsOfLocationGet) NewRequestBody() FloorRoomsOfLocationGetBody {
	return FloorRoomsOfLocationGetBody{}
}

type FloorRoomsOfLocationGetBody struct {
}

func (r *FloorRoomsOfLocationGet) RequestBody() *FloorRoomsOfLocationGetBody {
	return nil
}

func (r *FloorRoomsOfLocationGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *FloorRoomsOfLocationGet) SetRequestBody(body FloorRoomsOfLocationGetBody) {
	r.requestBody = body
}

func (r *FloorRoomsOfLocationGet) NewResponseBody() *FloorRoomsOfLocationGetResponseBody {
	return &FloorRoomsOfLocationGetResponseBody{}
}

type FloorRoomsOfLocationGetResponseBody Rooms

func (r *FloorRoomsOfLocationGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/location/{{.location}}/floor/{{.floor}}/room", r.PathParams())
	return &u
}

func (r *FloorRoomsOfLocationGet) Do() (FloorRoomsOfLocationGetResponseBody, error) {
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
