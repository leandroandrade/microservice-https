package main

import (
	"github.com/leandroandrade/microservice-https/homepage"
	"github.com/leandroandrade/microservice-https/server"
	"log"
	"net/http"
	"os"
)

var (
	MSCertFile    = os.Getenv("MS_CERT_FILE")
	MSKeyFile     = os.Getenv("MS_KEY_FILE")
	MSServiceAddr = os.Getenv("MS_SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "MS ", log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, MSServiceAddr)

	logger.Println("server starting")
	err := srv.ListenAndServeTLS(MSCertFile, MSKeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
