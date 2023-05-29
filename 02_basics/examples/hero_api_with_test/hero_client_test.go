package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeroClint(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, req.URL.String(), "/some/path")
			rw.Write([]byte(`OK`))
		}),
	)

	defer server.Close()

	api := HeroClient{Client: server.Client(),
		Url: server.URL}
	r, err := api.GetThem()

	assert.Nil(t, err)

	fmt.Printf("%v", r)
}
