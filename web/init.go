package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/NickLovera/weather-service/mgr"
)

type IWeatherServer interface {
	InitServer() *http.ServeMux
}

type weatherServer struct {
	weatherSvc mgr.IWeatherService
}

func NewWeatherServer(weatherSvc mgr.IWeatherService) IWeatherServer {
	return &weatherServer{
		weatherSvc: weatherSvc,
	}
}

func (ws *weatherServer) InitServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("resources"))))

	mux.HandleFunc("/swagger-ui", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger-ui/", http.StatusMovedPermanently)
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, "/swagger-ui/", http.StatusMovedPermanently)
	})

	//Our endpoint
	mux.HandleFunc("/currentweather/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		path := strings.TrimPrefix(r.URL.Path, "/currentweather/")
		parts := strings.Split(path, "/")
		if len(parts) != 2 {
			http.Error(w, "invalid path", http.StatusBadRequest)
			return
		}

		lat, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			http.Error(w, "invalid lat", http.StatusBadRequest)
			return
		}

		long, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			http.Error(w, "invalid long", http.StatusBadRequest)
			return
		}

		resp, err := ws.weatherSvc.GetCurrentWeatherByLatLong(r.Context(), lat, long)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get current weather Err: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	return mux
}
