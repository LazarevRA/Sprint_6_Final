package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type ServerStr struct {
	Log  *log.Logger
	Http *http.Server
}

func Server(logger *log.Logger) *ServerStr {

	r := http.NewServeMux()
	r.HandleFunc("/", handlers.GetIndex)
	r.HandleFunc("/upload", handlers.Upload)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &ServerStr{
		Log:  logger,
		Http: server,
	}
}
