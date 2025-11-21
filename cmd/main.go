package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "server: ", log.Ldate|log.LstdFlags|log.Lshortfile)
	serv := server.NewServer(logger)

	logger.Println("Starting server")
	if err := serv.Http.ListenAndServe(); err != nil {
		logger.Fatal("Ошибка запуска сервера:", err)
	}

}
