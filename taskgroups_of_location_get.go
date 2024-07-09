package legionella_dossier

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-legionella-dossier/utils"
)

func (c *Client) NewTaskgroupsOfLocationsGet() TaskgroupsOfLocationsGet {
	r := TaskgroupsOfLocationsGet{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaskgroupsOfLocationsGet struct {
	client      *Client
	queryParams *TaskgroupsOfLocationsGetQueryParams
	pathParams  *TaskgroupsOfLocationsGetPathParams
	method      string
	headers     http.Header
	requestBody TaskgroupsOfLocationsGetBody
}

func (r TaskgroupsOfLocationsGet) NewQueryParams() *TaskgroupsOfLocationsGetQueryParams {
	return &TaskgroupsOfLocationsGetQueryParams{}
}

type TaskgroupsOfLocationsGetQueryParams struct {
}

func (p TaskgroupsOfLocationsGetQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaskgroupsOfLocationsGet) QueryParams() QueryParams {
	return r.queryParams
}

func (r TaskgroupsOfLocationsGet) NewPathParams() *TaskgroupsOfLocationsGetPathParams {
	return &TaskgroupsOfLocationsGetPathParams{}
}

type TaskgroupsOfLocationsGetPathParams struct {
	LocationID int `schema:"location"`
}

func (p *TaskgroupsOfLocationsGetPathParams) Params() map[string]string {
	return map[string]string{
		"location": strconv.Itoa(p.LocationID),
	}
}

func (r *TaskgroupsOfLocationsGet) PathParams() *TaskgroupsOfLocationsGetPathParams {
	return r.pathParams
}

func (r *TaskgroupsOfLocationsGet) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaskgroupsOfLocationsGet) SetMethod(method string) {
	r.method = method
}

func (r *TaskgroupsOfLocationsGet) Method() string {
	return r.method
}

func (r TaskgroupsOfLocationsGet) NewRequestBody() TaskgroupsOfLocationsGetBody {
	return TaskgroupsOfLocationsGetBody{}
}

type TaskgroupsOfLocationsGetBody struct {
}

func (r *TaskgroupsOfLocationsGet) RequestBody() *TaskgroupsOfLocationsGetBody {
	return nil
}

func (r *TaskgroupsOfLocationsGet) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaskgroupsOfLocationsGet) SetRequestBody(body TaskgroupsOfLocationsGetBody) {
	r.requestBody = body
}

func (r *TaskgroupsOfLocationsGet) NewResponseBody() *TaskgroupsOfLocationsGetResponseBody {
	return &TaskgroupsOfLocationsGetResponseBody{}
}

type TaskgroupsOfLocationsGetResponseBody Taskgroups

func (r *TaskgroupsOfLocationsGet) URL() *url.URL {
	u := r.client.GetEndpointURL("/api3/locations/{{.location}}/taskgroups", r.PathParams())
	return &u
}

func (r *TaskgroupsOfLocationsGet) Do() (TaskgroupsOfLocationsGetResponseBody, error) {
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
