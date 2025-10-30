package main

import (
	"iot/internal/env"
	"iot/internal/store"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type dbconfig struct {
	user string
	pswd string
	addr string
	db   string
}

type config struct {
	addr        string
	db          dbconfig
	frontendURL string
}

type application struct {
	config config
	store  store.Store
}

func (app *application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{env.GetString("CORS_ALLOWED_ORIGIN", "http://localhost:3000")}, // FE URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/v1", func(r chi.Router) {

		r.Route("/cpu", func(r chi.Router) {
			r.Get("/statistics", app.getCPUStatistics)
			r.Route("/aggregation", func(r chi.Router) {
				r.Get("/location", app.getCPUAggregationPerLocation)
				r.Get("/model", app.getCPUAggregationPerModel)
			})
			r.Route("/dailyaggregation", func(r chi.Router) {
				r.Get("/location", app.getCPUDailyAggregationPerLocation)
			})
		})

		r.Route("/gps", func(r chi.Router) {
			r.Get("/statistics", app.getGPSStatistics)
			r.Route("/aggregation", func(r chi.Router) {
				r.Get("/location", app.getGPSAggregationPerLocation)
				r.Get("/model", app.getGPSAggregationPerModel)
			})
			r.Route("/dailyaggregation", func(r chi.Router) {
				r.Get("/location", app.getGPSDailyAggregationPerLocation)
			})
		})

		r.Route("/humidity", func(r chi.Router) {
			r.Get("/statistics", app.getHumidityStatistics)
			r.Route("/aggregation", func(r chi.Router) {
				r.Get("/location", app.getHumidityAggregationPerLocation)
				r.Get("/model", app.getHumidityAggregationPerModel)
			})
			r.Route("/dailyaggregation", func(r chi.Router) {
				r.Get("/location", app.getHumidityDailyAggregationPerLocation)
			})
		})

		r.Route("/pressure", func(r chi.Router) {
			r.Get("/statistics", app.getPressureStatistics)
			r.Route("/aggregation", func(r chi.Router) {
				r.Get("/location", app.getPressureAggregationPerLocation)
				r.Get("/model", app.getPressureAggregationPerModel)
			})
			r.Route("/dailyaggregation", func(r chi.Router) {
				r.Get("/location", app.getPressureDailyAggregationPerLocation)
			})
		})

		r.Route("/temperature", func(r chi.Router) {
			r.Get("/statistics", app.getTemperatureStatistics)
			r.Route("/aggregation", func(r chi.Router) {
				r.Get("/location", app.getTemperatureAggregationPerLocation)
				r.Get("/model", app.getTemperatureAggregationPerModel)
			})
			r.Route("/dailyaggregation", func(r chi.Router) {
				r.Get("/location", app.getTemperatureDailyAggregationPerLocation)
			})
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Millisecond,
	}
	log.Printf("Started server on %v\n", app.config.addr)
	err := srv.ListenAndServe()
	log.Printf("ListenAndServe is a blocking call and wint be executed unless it throws ant err %v", err)
	if err != nil {
		return err
	}
	return nil
}
