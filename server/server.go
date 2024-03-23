package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-rest-api/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	server := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                               // This allows all origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}          // Specify HTTP methods allowed
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"} // Specify headers allowed

	// Apply CORS middleware to the server
	server.Use(cors.New(config))

	return Server{
		port:   "5000",
		server: server,
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
