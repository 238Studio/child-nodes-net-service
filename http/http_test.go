package http_test

import (
	"net/url"
	"testing"

	"github.com/238Studio/child-nodes-net-service/http"
)

func TestHTTP(t *testing.T) {
	var header = make(map[string]string)
	header["test"] = "test"

	body := url.Values{}
	body.Set("test", "114514")

	responseBody, responseHeader, err := http.Get("http://127.0.0.1:8080/test", header, body)
	if err != nil {
		t.Error(err)
	}
	t.Log(responseBody)
	t.Log(responseHeader)
}
