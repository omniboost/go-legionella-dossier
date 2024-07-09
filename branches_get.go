package legionella_dossier

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewBranchesGet() BranchesGet {
	r := BranchesGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type BranchesGet struct {
	client      *Client
	queryParams *BranchesGetQueryParams
	pathParams  *BranchesGetPathParams
	method      string
	headers     http.Header
	requestBody BranchesGetBody
}

func (r BranchesGet) NewQueryParams() *BranchesGetQueryParams {
	return &BranchesGetQueryParams{}
}

type BranchesGetQueryParams struct {
}

func (p BranchesGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *BranchesGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r BranchesGet) NewPathParams() *BranchesGetPathParams {
	return &BranchesGetPathParams{}
}

type BranchesGetPathParams struct{}

func (p *BranchesGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *BranchesGet) PathParams() *BranchesGetPathParams {
	return r.pathParams
}

func (r *BranchesGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *BranchesGet) SetMethod(method string) {
	r.method = method
}

func (r *BranchesGet) Method() string {
	return r.method
}

func (r BranchesGet) NewRequestBody() BranchesGetBody {
	return BranchesGetBody{}
}

type BranchesGetBody struct {
}

func (r *BranchesGet) RequestBody() *BranchesGetBody {
	return nil
}

func (r *BranchesGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *BranchesGet) SetRequestBody(body BranchesGetBody) {
	r.requestBody = body
}

func (r *BranchesGet) NewResponseBody() *BranchesGetResponseBody {
	return &BranchesGetResponseBody{}
}

type BranchesGetResponseBody Branches

func (r *BranchesGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/branches", r.PathParams())
	return &u
}

func (r *BranchesGet) Do() (BranchesGetResponseBody, error) {
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
