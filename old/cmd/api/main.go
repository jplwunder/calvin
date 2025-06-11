package api

import (
	"calvin/internal/data"
	"calvin/internal/jsonlog"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5/pgxpool"
)

const version = "1.0.0"
const corsTrustedOrigins = "http://localhost:3000"

type config struct {
	port        int
	environment string
	limiter     struct {
		rps     float64
		burst   int
		enabled bool
	}
	db struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	cors struct {
		trustedOrigins []string
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
	db     *pgxpool.Pool
}

func Run() {
	var cfg config

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "DEVELOPMENT"
	}
	cfg.environment = environment

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.environment, "environment", cfg.environment, "Environment (DEVELOPMENT|STAGING|PRODUCTION)")
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")

	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		dbDSN = "postgres://user:pass@localhost:5432/dbname?sslmode=disable"
	}
	flag.StringVar(&cfg.db.dsn, "db-dsn", dbDSN, "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	var enable_limiter bool
	if environment == "PRODUCTION" {
		enable_limiter = true
	} else {
		enable_limiter = false
	}
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", enable_limiter, "Enable rate limiter")
	flag.Parse()

	cfg.cors.trustedOrigins = strings.Fields(corsTrustedOrigins)

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()
	logger.PrintInfo("database connection pool established", nil)

	models := data.NewModels(db) // Pass the db connection pool

	app := &application{
		config: cfg,
		logger: logger,
		models: models, // Store the models
		db:     db,
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
		"env":  cfg.environment,
	})

	err = srv.ListenAndServe()
	logger.PrintFatal(err, nil)
}

func openDB(cfg config) (*pgxpool.Pool, error) {
	pgxCfg, err := pgxpool.ParseConfig(cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	pgxCfg.MaxConns = int32(cfg.db.maxOpenConns)

	maxIdleTime, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	pgxCfg.MaxConnIdleTime = maxIdleTime

	db, err := pgxpool.NewWithConfig(context.Background(), pgxCfg)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}
