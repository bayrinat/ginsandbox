package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var routeTable = []struct {
	in struct {
		method string
		path   string
		reader io.Reader
	}
	out int
}{
	{
		struct {
			method string
			path   string
			reader io.Reader
		}{method: "POST", path: checkTextPath, reader: okDataReader()},
		http.StatusOK,
	},
	{
		struct {
			method string
			path   string
			reader io.Reader
		}{method: "POST", path: checkTextPath, reader: noContentDataReader()},
		http.StatusNoContent,
	},
	{
		struct {
			method string
			path   string
			reader io.Reader
		}{method: "POST", path: "/checkSomethingElse", reader: okDataReader()},
		http.StatusNotFound,
	},
	{
		struct {
			method string
			path   string
			reader io.Reader
		}{method: "GET", path: checkTextPath, reader: okDataReader()},
		http.StatusNotFound,
	},
}

func okDataReader() io.Reader {
	request := NewRequest()

	request.Site = []string{"https://google.com", "https://yahoo.com"}
	request.SearchText = "Google"

	data, _ := json.Marshal(request)

	return bytes.NewReader(data)
}

func noContentDataReader() io.Reader {
	request := NewRequest()

	request.Site = []string{"https://google.com", "https://yahoo.com"}
	request.SearchText = "Yandex"

	data, _ := json.Marshal(request)

	return bytes.NewReader(data)
}

func TestRouter(t *testing.T) {
	for _, data := range routeTable {
		req, _ := http.NewRequest(data.in.method, data.in.path, data.in.reader)
		w := httptest.NewRecorder()

		router := NewRouter()
		router.Initialize()

		router.router.ServeHTTP(w, req)

		if w.Code != data.out {
			t.Errorf("Response code should be %v, was: %v", data.out, w.Code)
		}
	}
}
