package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
)

func TestAll (t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:9090/", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	ReplyHello(res, req)

	exp := "Hello world version 6!"
	act := res.Body.String()
	actCode := res.Code

	if 200 != actCode {
		t.Fatalf("Expected %d got %d", 200, actCode)
	}

	if exp != act {
		t.Fatalf("Expected %s got %s", exp, act)
	}

	req, err = http.NewRequest("GET", "http://localhost:9090/test/", nil)

	if err != nil {
		t.Fatal(err)
	}

	res = httptest.NewRecorder()
	ReplyHello(res, req)

	actCode = res.Code

	if 404 != actCode {
		fmt.Printf("%v", res.Body.String())
		t.Fatalf("Expected %d got %d", 404, actCode)
	}


}