package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	checkTextPath = "/checkText"
)

// Router wrap gin.Engine entity
type Router struct {
	router *gin.Engine
}

// NewRouter returns new Router instance
func NewRouter() *Router {
	return &Router{
		router: gin.Default(),
	}
}

// Initialize is initializing all allowed URL
func (r *Router) Initialize() {
	r.router.POST(checkTextPath, func(c *gin.Context) {
		var request Request
		if c.BindJSON(&request) == nil {
			searchText, ok := request.Search()
			if ok {
				response := NewResponse(searchText)
				c.JSON(http.StatusOK, response)
			} else {
				c.JSON(http.StatusNoContent, nil)
			}
		}
	})
}

// Run starts listening and serving HTTP requests
func (r *Router) Run(addr string) {
	r.router.Run(addr)
}
