package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// Не разобрался, как адекватно реализовать проверку методов внутри сервера
	// Реализовал внутри хендлера

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	http.ServeFile(w, r, "./index.html")

}

func Upload(w http.ResponseWriter, r *http.Request) {

	// Не разобрался, как адекватно реализовать проверку методов внутри сервера
	// Реализовал внутри хендлера

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusInternalServerError)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	convertedData := []byte(service.Convert(string(data)))

	locFileName := fmt.Sprintf("%s%s", time.Now().UTC().Format("20060102_150405"), filepath.Ext(header.Filename))

	if err = os.WriteFile(locFileName, convertedData, 0755); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if _, err = w.Write(convertedData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
