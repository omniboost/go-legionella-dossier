package legionella_dossier

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewVacancyOfLocationGet() VacancyOfLocationGet {
	r := VacancyOfLocationGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VacancyOfLocationGet struct {
	client      *Client
	queryParams *VacancyOfLocationGetQueryParams
	pathParams  *VacancyOfLocationGetPathParams
	method      string
	headers     http.Header
	requestBody VacancyOfLocationGetBody
}

func (r VacancyOfLocationGet) NewQueryParams() *VacancyOfLocationGetQueryParams {
	return &VacancyOfLocationGetQueryParams{}
}

type VacancyOfLocationGetQueryParams struct {
}

func (p VacancyOfLocationGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VacancyOfLocationGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r VacancyOfLocationGet) NewPathParams() *VacancyOfLocationGetPathParams {
	return &VacancyOfLocationGetPathParams{}
}

type VacancyOfLocationGetPathParams struct {
	LocationID int `schema:"location"`
}

func (p *VacancyOfLocationGetPathParams) Params() map[string]string {
	return map[string]string{
		"location": strconv.Itoa(p.LocationID),
	}
}

func (r *VacancyOfLocationGet) PathParams() *VacancyOfLocationGetPathParams {
	return r.pathParams
}

func (r *VacancyOfLocationGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VacancyOfLocationGet) SetMethod(method string) {
	r.method = method
}

func (r *VacancyOfLocationGet) Method() string {
	return r.method
}

func (r VacancyOfLocationGet) NewRequestBody() VacancyOfLocationGetBody {
	return VacancyOfLocationGetBody{}
}

type VacancyOfLocationGetBody struct {
}

func (r *VacancyOfLocationGet) RequestBody() *VacancyOfLocationGetBody {
	return nil
}

func (r *VacancyOfLocationGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *VacancyOfLocationGet) SetRequestBody(body VacancyOfLocationGetBody) {
	r.requestBody = body
}

func (r *VacancyOfLocationGet) NewResponseBody() *VacancyOfLocationGetResponseBody {
	return &VacancyOfLocationGetResponseBody{}
}

type VacancyOfLocationGetResponseBody Vacancies

func (r *VacancyOfLocationGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/location/{{.location}}/vacancy", r.PathParams())
	return &u
}

func (r *VacancyOfLocationGet) Do() (VacancyOfLocationGetResponseBody, error) {
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
