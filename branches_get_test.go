package legionella_dossier_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBranchesGet(t *testing.T) {
	req := client.NewBranchesGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
