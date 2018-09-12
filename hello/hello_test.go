package main

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// $ go test workshop/hello -bench=.
func BenchmarkHello(b *testing.B) {

	req := httptest.NewRequest("GET", "/", nil)

	for i := 0; i < b.N; i++ {
		rec := httptest.NewRecorder()
		hello(rec, req)
	}

}

// $ go test workshop/hello
func TestHello(t *testing.T) {

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	hello(rec, req)
	res := rec.Result()

	if res.StatusCode != 200 {
		t.Fatalf("Invalid http code, expected 200, given %d", res.StatusCode)
	}

	body, _ := ioutil.ReadAll(res.Body)

	if bytes.Contains(body, []byte("Hello World!")) == false {
		t.Fatalf("Invalid http body, expected \"Hello World!\", given \"%s\"", body)
	}

}
