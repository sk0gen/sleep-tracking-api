package api

import (
	"net/http"
	"os"
	"time"
)

func Run() {
	routersInit := initRouter()
	server := &http.Server{
		Addr:         os.Getenv("SERVER_PORT"),
		Handler:      routersInit,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
