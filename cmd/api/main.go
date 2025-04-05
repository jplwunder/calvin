package api

import (
	"calvin/connections"
	"calvin/internal/data"
	"calvin/internal/jsonlog"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const version = "1.0.0"
const corsTrustedOrigins = "http://localhost:3000"

type config struct {
	port    int
	env     string
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	cors struct {
		trustedOrigins []string
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
}

func Run() {
	var cfg config

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	cfg.port, _ = strconv.Atoi(port)
	flag.IntVar(&cfg.port, "port", cfg.port, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")
	flag.Parse()

	cfg.cors.trustedOrigins = strings.Fields(corsTrustedOrigins)

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	// Initialize MongoDB connection
	connections.Connect()

	models := data.NewModels()

	app := &application{
		config: cfg,
		logger: logger,
		models: models,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     log.New(logger, "", 0),
	}

	logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  cfg.env,
	})

	err = srv.ListenAndServe()
	logger.PrintFatal(err, nil)
}
