package main

import (
	"database/sql"
	"log"

	"auth-system/internal/api"
	sqlc "auth-system/internal/db/sqlc"
	"auth-system/internal/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Konfigürasyon okunamadı: %v", err)
	}
	conn, err := sql.Open("postgres", config.DBSource)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatalf("Veritabanı bağlantısı doğrulanamadı: %v", err)
	}

	store := sqlc.New(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("API sunucusu oluşturulamadı: %v", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("API sunucusu başlatılamadı: %v", err)
	}
}
