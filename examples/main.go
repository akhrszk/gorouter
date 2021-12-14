package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/akhrszk/gorouter"
)

const port = 3000

func main() {
	r := gorouter.NewRouter()
	r.Get("/", Index)
	r.Get("/sum/:num1(\\d)/:num2(\\d)", Sum)
	r.Get("/:name(\\w+)/hello", Hello)

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

func Index(w http.ResponseWriter, r *http.Request, _ gorouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome!!")
}

func Sum(w http.ResponseWriter, r *http.Request, params gorouter.Params) {
	w.WriteHeader(http.StatusOK)
	num1, _ := strconv.Atoi(params["num1"])
	num2, _ := strconv.Atoi(params["num2"])
	s := fmt.Sprintf("%d+%d=%d", num1, num2, num1+num2)
	fmt.Fprintf(w, s)
}

func Hello(w http.ResponseWriter, r *http.Request, params gorouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, params["name"]+", Hello!!")
}
