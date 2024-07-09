package legionella_dossier_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVacancyOfLocationByIDGet(t *testing.T) {
	req := client.NewVacancyOfLocationByIDGet()
	req.PathParams().LocationID = 5301
	req.PathParams().VacancyID = 1
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
