package legionella_dossier

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewVacancyLocationByIDGet() VacancyLocationByIDGet {
	r := VacancyLocationByIDGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VacancyLocationByIDGet struct {
	client      *Client
	queryParams *VacancyLocationByIDGetQueryParams
	pathParams  *VacancyLocationByIDGetPathParams
	method      string
	headers     http.Header
	requestBody VacancyLocationByIDGetBody
}

func (r VacancyLocationByIDGet) NewQueryParams() *VacancyLocationByIDGetQueryParams {
	return &VacancyLocationByIDGetQueryParams{}
}

type VacancyLocationByIDGetQueryParams struct {
}

func (p VacancyLocationByIDGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VacancyLocationByIDGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r VacancyLocationByIDGet) NewPathParams() *VacancyLocationByIDGetPathParams {
	return &VacancyLocationByIDGetPathParams{}
}

type VacancyLocationByIDGetPathParams struct {
	LocationID int `schema:"location"`
}

func (p *VacancyLocationByIDGetPathParams) Params() map[string]string {
	return map[string]string{
		"location": strconv.Itoa(p.LocationID),
	}
}

func (r *VacancyLocationByIDGet) PathParams() *VacancyLocationByIDGetPathParams {
	return r.pathParams
}

func (r *VacancyLocationByIDGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VacancyLocationByIDGet) SetMethod(method string) {
	r.method = method
}

func (r *VacancyLocationByIDGet) Method() string {
	return r.method
}

func (r VacancyLocationByIDGet) NewRequestBody() VacancyLocationByIDGetBody {
	return VacancyLocationByIDGetBody{}
}

type VacancyLocationByIDGetBody struct {
}

func (r *VacancyLocationByIDGet) RequestBody() *VacancyLocationByIDGetBody {
	return nil
}

func (r *VacancyLocationByIDGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *VacancyLocationByIDGet) SetRequestBody(body VacancyLocationByIDGetBody) {
	r.requestBody = body
}

func (r *VacancyLocationByIDGet) NewResponseBody() *VacancyLocationByIDGetResponseBody {
	return &VacancyLocationByIDGetResponseBody{}
}

type VacancyLocationByIDGetResponseBody VacancyLocation

func (r *VacancyLocationByIDGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/location/{{.location}}", r.PathParams())
	return &u
}

func (r *VacancyLocationByIDGet) Do() (VacancyLocationByIDGetResponseBody, error) {
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
