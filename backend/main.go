package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync/atomic"

	pb "v5crimson/backend/proto" 
	"google.golang.org/grpc"
	"github.com/rs/cors" 
)

// Global counter variable
var counter int64

// gRPC Server struct
type server struct {
	pb.UnimplementedCounterServer
}

// gRPC Increment method, listens for gRPC requests to increment the counter
func (s *server) Increment(ctx context.Context, req *pb.IncrementRequest) (*pb.IncrementResponse, error) {
	// Increment the counter atomically
	counterValue := atomic.AddInt64(&counter, 1)

	// Return the updated counter value
	return &pb.IncrementResponse{
		Value: counterValue,
	}, nil
}

// HTTP Handler to get counter value
func valueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"value": atomic.LoadInt64(&counter)})
}

// HTTP Handler to increment counter
func incrementHandler(w http.ResponseWriter, r *http.Request) {
	// Increment counter atomically
	counterValue := atomic.AddInt64(&counter, 1)

	// Respond with the updated value
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"value": counterValue})
}

// Start the HTTP server with CORS
func startHTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/value", valueHandler)
	mux.HandleFunc("/increment", incrementHandler)

	// Set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, 
		AllowedMethods: []string{"GET", "POST"},          
		AllowedHeaders: []string{"Content-Type"},         
	})

	handler := c.Handler(mux) 

	fmt.Println("HTTP Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// Start the gRPC server
func startGRPCServer() {
	// Set up the TCP listener for gRPC
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the gRPC server
	pb.RegisterCounterServer(s, &server{})

	fmt.Println("gRPC Server listening on port 50051...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Main function to run both servers
func main() {
	// Run HTTP and gRPC servers in parallel
	go startHTTPServer()
	startGRPCServer()
}
