package legionella_dossier_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVacancyOfLocationGet(t *testing.T) {
	req := client.NewVacancyOfLocationGet()
	req.PathParams().LocationID = 5301
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
