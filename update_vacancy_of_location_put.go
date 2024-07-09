package legionella_dossier

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewUpdateVacancyOfLocationPut() UpdateVacancyOfLocationPut {
	r := UpdateVacancyOfLocationPut{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type UpdateVacancyOfLocationPut struct {
	client      *Client
	queryParams *UpdateVacancyOfLocationPutQueryParams
	pathParams  *UpdateVacancyOfLocationPutPathParams
	method      string
	headers     http.Header
	requestBody UpdateVacancyOfLocationPutBody
}

func (r UpdateVacancyOfLocationPut) NewQueryParams() *UpdateVacancyOfLocationPutQueryParams {
	return &UpdateVacancyOfLocationPutQueryParams{}
}

type UpdateVacancyOfLocationPutQueryParams struct{}

func (p UpdateVacancyOfLocationPutQueryParams) ToURLValues() (url.Values, error) {
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

func (r *UpdateVacancyOfLocationPut) QueryParams() *UpdateVacancyOfLocationPutQueryParams {
	return r.queryParams
}

func (r UpdateVacancyOfLocationPut) NewPathParams() *UpdateVacancyOfLocationPutPathParams {
	return &UpdateVacancyOfLocationPutPathParams{}
}

type UpdateVacancyOfLocationPutPathParams struct {
	LocationID int `schema:"location"`
	VacancyID  int `schema:"vacancy"`
}

func (p *UpdateVacancyOfLocationPutPathParams) Params() map[string]string {
	return map[string]string{
		"location": strconv.Itoa(p.LocationID),
		"vacancy":  strconv.Itoa(p.VacancyID),
	}
}

func (r *UpdateVacancyOfLocationPut) PathParams() *UpdateVacancyOfLocationPutPathParams {
	return r.pathParams
}

func (r *UpdateVacancyOfLocationPut) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *UpdateVacancyOfLocationPut) SetMethod(method string) {
	r.method = method
}

func (r *UpdateVacancyOfLocationPut) Method() string {
	return r.method
}

func (r UpdateVacancyOfLocationPut) NewRequestBody() UpdateVacancyOfLocationPutBody {
	return UpdateVacancyOfLocationPutBody{}
}

type UpdateVacancyOfLocationPutBody struct {
	Finished bool `json:"finished"`
}

func (r *UpdateVacancyOfLocationPut) RequestBody() *UpdateVacancyOfLocationPutBody {
	return &r.requestBody
}

func (r *UpdateVacancyOfLocationPut) RequestBodyInterface() interface{} {
	return r.RequestBody()
}

func (r *UpdateVacancyOfLocationPut) SetRequestBody(body UpdateVacancyOfLocationPutBody) {
	r.requestBody = body
}

func (r *UpdateVacancyOfLocationPut) NewResponseBody() *UpdateVacancyOfLocationPutResponseBody {
	return &UpdateVacancyOfLocationPutResponseBody{}
}

type UpdateVacancyOfLocationPutResponseBody Vacancy

func (r *UpdateVacancyOfLocationPut) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/location/{{.location}}/vacancy/{{.vacancy}}", r.PathParams())
	return &u
}

func (r *UpdateVacancyOfLocationPut) Do() (UpdateVacancyOfLocationPutResponseBody, error) {
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
