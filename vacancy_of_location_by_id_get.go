package legionella_dossier

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewVacancyOfLocationByIDGet() VacancyOfLocationByIDGet {
	r := VacancyOfLocationByIDGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VacancyOfLocationByIDGet struct {
	client      *Client
	queryParams *VacancyOfLocationByIDGetQueryParams
	pathParams  *VacancyOfLocationByIDGetPathParams
	method      string
	headers     http.Header
	requestBody VacancyOfLocationByIDGetBody
}

func (r VacancyOfLocationByIDGet) NewQueryParams() *VacancyOfLocationByIDGetQueryParams {
	return &VacancyOfLocationByIDGetQueryParams{}
}

type VacancyOfLocationByIDGetQueryParams struct{}

func (p VacancyOfLocationByIDGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VacancyOfLocationByIDGet) QueryParams() *VacancyOfLocationByIDGetQueryParams {
	return r.queryParams
}

func (r VacancyOfLocationByIDGet) NewPathParams() *VacancyOfLocationByIDGetPathParams {
	return &VacancyOfLocationByIDGetPathParams{}
}

type VacancyOfLocationByIDGetPathParams struct {
	LocationID int `schema:"location"`
	VacancyID  int `schema:"vacancy"`
}

func (p *VacancyOfLocationByIDGetPathParams) Params() map[string]string {
	return map[string]string{
		"location": strconv.Itoa(p.LocationID),
		"vacancy":  strconv.Itoa(p.VacancyID),
	}
}

func (r *VacancyOfLocationByIDGet) PathParams() *VacancyOfLocationByIDGetPathParams {
	return r.pathParams
}

func (r *VacancyOfLocationByIDGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VacancyOfLocationByIDGet) SetMethod(method string) {
	r.method = method
}

func (r *VacancyOfLocationByIDGet) Method() string {
	return r.method
}

func (r VacancyOfLocationByIDGet) NewRequestBody() VacancyOfLocationByIDGetBody {
	return VacancyOfLocationByIDGetBody{}
}

type VacancyOfLocationByIDGetBody struct {
}

func (r *VacancyOfLocationByIDGet) RequestBody() *VacancyOfLocationByIDGetBody {
	return nil
}

func (r *VacancyOfLocationByIDGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *VacancyOfLocationByIDGet) SetRequestBody(body VacancyOfLocationByIDGetBody) {
	r.requestBody = body
}

func (r *VacancyOfLocationByIDGet) NewResponseBody() *VacancyOfLocationByIDGetResponseBody {
	return &VacancyOfLocationByIDGetResponseBody{}
}

type VacancyOfLocationByIDGetResponseBody Vacancy

func (r *VacancyOfLocationByIDGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/location/{{.location}}/vacancy/{{.vacancy}}", r.PathParams())
	return &u
}

func (r *VacancyOfLocationByIDGet) Do() (VacancyOfLocationByIDGetResponseBody, error) {
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
