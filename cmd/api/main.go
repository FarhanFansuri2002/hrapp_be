package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"server/internal/config"
	"server/internal/handler"
	"server/internal/repository"
	"server/internal/router"
	"server/internal/service"
)

func main() {
	cfg := config.Load()
	db, err := sql.Open("mysql", cfg.MySQLDSN)
	if err != nil {
		log.Fatalf("open mysql connection: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("connect to mysql: %v", err)
	}

	store := repository.NewMySQLStore(db)
	handlers := handler.New(service.NewHRService(store))
	server := &http.Server{
		Addr:    cfg.Address,
		Handler: router.New(handlers),
	}

	log.Printf("HR API listening on %s", cfg.Address)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("server stopped: %v", err)
		os.Exit(1)
	}
}
