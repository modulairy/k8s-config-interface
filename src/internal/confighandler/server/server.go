package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/modulairy/k8s-configmap-api-server/internal/confighandler/middleware"
	"github.com/modulairy/k8s-configmap-api-server/internal/confighandler/model"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/health":
		healthCheck(w, r)
	case "/config":
		var ctx *model.Context
		var err error
		ctx, err = middleware.InvokeContext(w, r)
		if err != nil {
			return
		}
		ctx, err = middleware.InvokeConfigure(ctx)
		if err != nil {
			return
		}
		middleware.InvokeHandler(ctx)
	default:
		http.NotFound(w, r)
		return
	}
}

func ServerStart() {
	port := "0.0.0.0:8080"
	server := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(handler),
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		fmt.Println("Server is starting...")
		fmt.Println("Server listening on", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()
	<-stop
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	} else {
		fmt.Println("Server has been shut down")
	}
}
