package legionella_dossier_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTaskgroupsOfLocationsGet(t *testing.T) {
	req := client.NewTaskgroupsOfLocationsGet()
	req.PathParams().LocationID = 73164
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
