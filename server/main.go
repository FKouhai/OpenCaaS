package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	c "github.com/FKouhai/OpenCaaServer/src/pkg/config"
	ph "github.com/FKouhai/OpenCaaServer/src/pkg/handlers"
	l "github.com/FKouhai/OpenCaaServer/src/pkg/logger"

	"github.com/gorilla/mux"
)

func main() {
	sm := mux.NewRouter()
	newLog := l.NewLogger()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.Index)
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/add", ph.Add)
	postRouter.HandleFunc("/remove", ph.Remove)
	getRouter.HandleFunc("/list", ph.List)
	config, err := c.NewConfig()
	if err != nil {
		l.LoggErr(newLog, err)
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
			l.LoggErr(newLog, err)
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
