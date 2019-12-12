package main

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/mrombout/solorpg/backend/rollfun"
	"log"
	"os"
)

func main() {
	funcframework.RegisterHTTPFunction("/", rollfun.Roll)
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
