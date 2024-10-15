package handlers

import (
	"github.com/Eagoker/url-shortener/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Handler struct {
	serverAddress string
	db            *pgxpool.Pool
	cfg           *config.Config 
}

// NewHandler создает новый хендлер с указанием адреса сервера и базы данных
func NewHandler(address string, db *pgxpool.Pool, cfg *config.Config) *Handler {
	return &Handler{serverAddress: address, db: db, cfg: cfg}
}