package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"go-rest-api/server/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port           string
	server         *gin.Engine
	readTimeout    time.Duration
	writeTimeout   time.Duration
	idleTimeout    time.Duration
	maxHeaderBytes int
}

func NewServer() Server {
	server := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	server.Use(cors.New(config))
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	return Server{
		port:           port,
		server:         server,
		readTimeout:    50 * time.Second,  // Adjust as necessary
		writeTimeout:   50 * time.Second,  // Adjust as necessary
		idleTimeout:    220 * time.Second, // Adjust as necessary
		maxHeaderBytes: 2 << 20,           // Example: 1 MB
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	httpServer := &http.Server{
		Addr:           "0.0.0.0:" + s.port, // Set the address and the port
		Handler:        router,       // Use the Gin router as the handler
		ReadTimeout:    s.readTimeout,
		WriteTimeout:   s.writeTimeout,
		IdleTimeout:    s.idleTimeout,
		MaxHeaderBytes: s.maxHeaderBytes,
	}

	log.Printf("Server running at port: %v", s.port)
	// Start the server with our custom configurations
	// log.Fatal here will exit the program if the server fails to start
	log.Fatal(httpServer.ListenAndServe())
}
