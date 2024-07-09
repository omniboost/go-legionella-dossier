package legionella_dossier_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFloorRoomsOfLocationGet(t *testing.T) {
	req := client.NewFloorRoomsOfLocationGet()
	req.PathParams().LocationID = 5301
	req.PathParams().FloorID = 30536
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
