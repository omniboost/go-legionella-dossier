package legionella_dossier

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewRelationsGet() RelationsGet {
	r := RelationsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type RelationsGet struct {
	client      *Client
	queryParams *RelationsGetQueryParams
	pathParams  *RelationsGetPathParams
	method      string
	headers     http.Header
	requestBody RelationsGetBody
}

func (r RelationsGet) NewQueryParams() *RelationsGetQueryParams {
	return &RelationsGetQueryParams{}
}

type RelationsGetQueryParams struct {
}

func (p RelationsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *RelationsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r RelationsGet) NewPathParams() *RelationsGetPathParams {
	return &RelationsGetPathParams{}
}

type RelationsGetPathParams struct{}

func (p *RelationsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *RelationsGet) PathParams() *RelationsGetPathParams {
	return r.pathParams
}

func (r *RelationsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *RelationsGet) SetMethod(method string) {
	r.method = method
}

func (r *RelationsGet) Method() string {
	return r.method
}

func (r RelationsGet) NewRequestBody() RelationsGetBody {
	return RelationsGetBody{}
}

type RelationsGetBody struct {
}

func (r *RelationsGet) RequestBody() *RelationsGetBody {
	return nil
}

func (r *RelationsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *RelationsGet) SetRequestBody(body RelationsGetBody) {
	r.requestBody = body
}

func (r *RelationsGet) NewResponseBody() *RelationsGetResponseBody {
	return &RelationsGetResponseBody{}
}

type RelationsGetResponseBody Relations

func (r *RelationsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/relations", r.PathParams())
	return &u
}

func (r *RelationsGet) Do() (RelationsGetResponseBody, error) {
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
