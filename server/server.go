package server

import (
	"log"

	"go-rest-api/server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true                                                               
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}          
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"} 

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
