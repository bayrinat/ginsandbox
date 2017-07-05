package http

import "strings"

// Request body
type Request struct {
	Site       []string `json:"Site" binding:"required"`
	SearchText string   `json:"SearchText" binding:"required"`
}

// NewRequest returns new empty Request instance
func NewRequest() *Request {
	return &Request{}
}

// Search returns the first match of site in SearchText slice and true.
// If there is no match, returns empty string and false
func (r *Request) Search() (string, bool) {
	for _, site := range r.Site {
		// Compare lower case strings
		if strings.Contains(strings.ToLower(site), strings.ToLower(r.SearchText)) {
			return site, true
		}
	}

	return "", false
}
