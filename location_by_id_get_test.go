package legionella_dossier_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLocationByIDGet(t *testing.T) {
	req := client.NewLocationByIDGet()
	req.PathParams().LocationID = 73164
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
