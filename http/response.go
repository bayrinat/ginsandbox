package http

// Response body
type Response struct {
	FoundAtSite string `json:"FoundAtSite"`
}

// NewResponse returns new Response instance
func NewResponse(site string) *Response {
	return &Response{
		FoundAtSite: site,
	}
}
