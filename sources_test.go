package newsapi

import (
	"testing"
)

func TestGetSources(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("")
}
