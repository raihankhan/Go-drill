package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	go startServer(&server)
	keepProcessAlive()
	shutDownServerGracefully(&server)
}

func keepProcessAlive() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		_, err := fmt.Fprintf(w, "hello")
		if err != nil {
			log.Println("Failed to write response to webserver")
			return
		}
	}
}

func startServer(server *http.Server) {
	fmt.Println("Starting server at port 8080")
	http.HandleFunc("/hello", helloHandler)
	err := server.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server closed")
	}
}

func shutDownServerGracefully(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("Server getting gracefully stopped")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
