package main

import (
	"context"
	"fmt"
	c "github.com/FKouhai/OpenCaaServer/src/pkg/config"
	ph "github.com/FKouhai/OpenCaaServer/src/pkg/handlers"
	l "github.com/FKouhai/OpenCaaServer/src/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.Index)
	//putRouter := sm.Methods(http.MethodPut).Subrouter()
	//putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)

	//	sm.Handle("/products", ph)
	config, err := c.NewConfig()
	if err != nil {
		l.Logger(err)
		return
	}

	s := &http.Server{
		Addr:         config.Port,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Hour,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Logger(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	fmt.Println("Received terminate, graceful shutdown", sig)
	to, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(to)
}
