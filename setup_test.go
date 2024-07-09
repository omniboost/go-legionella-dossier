package legionella_dossier_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	legionella_dossier "github.com/omniboost/go-legionella-dossier"
)

var (
	client *legionella_dossier.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	token := os.Getenv("LEGIONELLA_DOSSIER_TOKEN")
	debug := os.Getenv("DEBUG")

	client = legionella_dossier.NewClient(nil)
	client.SetToken(token)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	m.Run()
}
