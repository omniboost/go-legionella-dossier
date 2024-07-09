package legionella_dossier_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRelationsGet(t *testing.T) {
	req := client.NewRelationsGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
