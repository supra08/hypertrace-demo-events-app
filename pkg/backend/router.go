package backend

import "github.com/gin-gonic/gin"

// NewRouter creates a router to serve the backend's REST API.
func NewRouter() (*gin.Engine, error) {
	router := gin.Default()

	// TODO: register routes here

	return router, nil
}
