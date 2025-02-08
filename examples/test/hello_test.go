package test

import (
	"golang-tutorial/examples/simple-http-server/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HelloHandler(t *testing.T) {
	
	req, err := http.NewRequest("GET", "/", nil) 
	if err != nil {
		t.Fatal(err)
	}
	
	res := httptest.NewRecorder()
	
	handlers.HelloHandler(res, req)
	
	if res.Body.String() == "Hello, World!" {
		t.Logf("Response contains Hello World")
	} else {
		t.Errorf("Expected response code to be 200, but got %s", res.Body.String())
	}

	if res.Code == 200 {
		t.Logf("Response status code 200 OK")
	} else {
		t.Errorf("Expected response code to be 200, but got %d", res.Code)
	}
}
