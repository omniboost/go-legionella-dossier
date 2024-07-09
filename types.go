package legionella_dossier

import (
	"strings"
)

type CommaSeparatedQueryParam []string

func (t CommaSeparatedQueryParam) MarshalSchema() string {
	return strings.Join(t, ",")
}

type Branches []Branch

type Branch struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	CertificateNumber string `json:"certificate_number"`
	Address           struct {
		Street       string `json:"street"`
		StreetNumber string `json:"street_number"`
		ZipCode      string `json:"zip_code"`
		City         string `json:"city"`
		Country      string `json:"country"`
	} `json:"address"`
	Image      string `json:"image"`
	HeadBranch bool   `json:"head_branch"`
}

type Relations []Relation

type Relation struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Locations []Location

type Location struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Relation struct {
		ID int `json:"id"`
	} `json:"relation"`
	Branch struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"branch"`
}

type Taskgroups []Taskgroup

type Taskgroup struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Relation struct {
		ID int `json:"id"`
	} `json:"relation"`
	Branch struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"branch"`
}
