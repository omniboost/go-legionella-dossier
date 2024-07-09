package legionella_dossier_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFloorOfLocationGet(t *testing.T) {
	req := client.NewFloorOfLocationGet()
	req.PathParams().LocationID = 5301
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
