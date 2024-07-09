package legionella_dossier

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewVacancyOfLocationPost() VacancyOfLocationPost {
	r := VacancyOfLocationPost{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VacancyOfLocationPost struct {
	client      *Client
	queryParams *VacancyOfLocationPostQueryParams
	pathParams  *VacancyOfLocationPostPathParams
	method      string
	headers     http.Header
	requestBody VacancyOfLocationPostBody
}

func (r VacancyOfLocationPost) NewQueryParams() *VacancyOfLocationPostQueryParams {
	return &VacancyOfLocationPostQueryParams{}
}

type VacancyOfLocationPostQueryParams struct{}

func (p VacancyOfLocationPostQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VacancyOfLocationPost) QueryParams() *VacancyOfLocationPostQueryParams {
	return r.queryParams
}

func (r VacancyOfLocationPost) NewPathParams() *VacancyOfLocationPostPathParams {
	return &VacancyOfLocationPostPathParams{}
}

type VacancyOfLocationPostPathParams struct {
	LocationID int `schema:"location"`
}

func (p *VacancyOfLocationPostPathParams) Params() map[string]string {
	return map[string]string{
		"location": strconv.Itoa(p.LocationID),
	}
}

func (r *VacancyOfLocationPost) PathParams() *VacancyOfLocationPostPathParams {
	return r.pathParams
}

func (r *VacancyOfLocationPost) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VacancyOfLocationPost) SetMethod(method string) {
	r.method = method
}

func (r *VacancyOfLocationPost) Method() string {
	return r.method
}

func (r VacancyOfLocationPost) NewRequestBody() VacancyOfLocationPostBody {
	return VacancyOfLocationPostBody{}
}

type VacancyOfLocationPostBody struct {
	Description string   `json:"description"`
	Finished    bool     `json:"finished"`
	Creation    DateTime `json:"creation"`
}

func (r *VacancyOfLocationPost) RequestBody() *VacancyOfLocationPostBody {
	return &r.requestBody
}

func (r *VacancyOfLocationPost) RequestBodyInterface() interface{} {
	return r.RequestBody()
}

func (r *VacancyOfLocationPost) SetRequestBody(body VacancyOfLocationPostBody) {
	r.requestBody = body
}

func (r *VacancyOfLocationPost) NewResponseBody() *VacancyOfLocationPostResponseBody {
	return &VacancyOfLocationPostResponseBody{}
}

type VacancyOfLocationPostResponseBody Vacancy

func (r *VacancyOfLocationPost) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/location/{{.location}}/vacancy", r.PathParams())
	return &u
}

func (r *VacancyOfLocationPost) Do() (VacancyOfLocationPostResponseBody, error) {
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
