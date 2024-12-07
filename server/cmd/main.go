package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/db"
	"todo/middleware"
	"todo/routes"
)

func main() {
	router := routes.InitRoutes()
	handler := middleware.CORS(router)
	log.Println("Server running on port 8080")
	if err := db.ConnectToMongo("mongodb+srv://mango26june:mango123@cluster0.ga9pq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func gracefulShutdown(server *http.Server) {
	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Wait for a signal
	<-quit
	log.Println("Shutting down server...")

	// Context with timeout to allow cleanup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Close the MongoDB connection
	if db.MongoClient != nil {
		if err := db.MongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect MongoDB client: %v", err)
		}
		log.Println("MongoDB connection closed.")
	}

	log.Println("Server exited gracefully.")
}
