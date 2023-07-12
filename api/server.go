package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kierquebs/capcut-scraper/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	router     *gin.Engine
}

func NewServer(config util.Config) (*Server, error) {

	server := &Server{
		config:     config,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// Users
	router.GET("/capcut/template/:id", server.getRenderData)
	router.GET("/capcut/template-detail/:id", server.getRenderDetail)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
