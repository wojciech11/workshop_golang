package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHelloMsgHandler(t *testing.T) {

	app := App{}
	// Create a request
	r, err := http.NewRequest("GET",
		"http://test.com?lang=pl&user=wojtek", nil)
	if err != nil {
		t.Fatal("failed to create request")
	}

	// Handle request and store result in w
	w := httptest.NewRecorder()
	app.GetHelloMsgHandler(w, r)

	if w.Code != http.StatusOK {
		t.Fatal(w.Code, w.Body.String())
	}
}
