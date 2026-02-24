package repos

import (
	"URL-Shortener/internal/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type UrlRepo struct {
	conn *pgx.Conn
}

func NewUrlRepos(cfg *config.Config) *UrlRepo {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if err != nil {
		log.Fatalf("Cant connect to database: %v\n", err)
	}
	return &UrlRepo{conn: conn}
}

func (r *UrlRepo) Save(shortUrl, originalUrl string) {
	_, err := r.conn.Exec(context.Background(), "INSERT INTO urls (short_url, original_url) VALUES ($1, $2)", shortUrl, originalUrl)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func (r *UrlRepo) Get(shortUrl string) (string, bool) {
	var original_url string
	err := r.conn.QueryRow(context.Background(), "SELECT original_url FROM urls WHERE short_url=$1", shortUrl).Scan(&original_url)
	if err != nil {
		return "", false
	}
	return original_url, true
}
