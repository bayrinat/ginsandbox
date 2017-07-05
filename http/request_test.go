package http

import (
	"testing"
)

var requestTable = []struct {
	in  Request
	out struct {
		string
		bool
	}
}{
	{
		Request{
			Site:       []string{"https://google.com", "https://yahoo.com"},
			SearchText: "Google",
		},
		struct {
			string
			bool
		}{string: "https://google.com", bool: true},
	},
	{
		Request{
			Site:       []string{"", "https://google.com", "https://yahoo.com"},
			SearchText: "Google",
		},
		struct {
			string
			bool
		}{string: "https://google.com", bool: true},
	},
	{
		Request{
			Site:       nil,
			SearchText: "Google",
		},
		struct {
			string
			bool
		}{string: "", bool: false},
	},
	{
		Request{
			Site:       []string{"https://google.com", "https://yahoo.com"},
			SearchText: "Yandex",
		},
		struct {
			string
			bool
		}{string: "", bool: false},
	},
}

func TestRequest_Search(t *testing.T) {
	for _, req := range requestTable {
		site, ok := req.in.Search()

		if ok != req.out.bool {
			t.Errorf("Ok flag is %v, want: %v", ok, req.out.bool)
		}
		if site != req.out.string {
			t.Errorf("Site is %v, want: %v", site, req.out.string)
		}
	}
}
