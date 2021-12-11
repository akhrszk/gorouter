package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/akhrszk/gorouter"
)

const port = 3000

func main() {
	r := gorouter.NewRouter()
	r.Get("/", Index)
	r.Get("/say", Say)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		log.Printf("This app listening on :%d", port)

		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		log.Print("shutting down...")
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal("Server forced shutdown:", err)
		}
	}()

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome!!")
}

func Say(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "にゃーん")
}
