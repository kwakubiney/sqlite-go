package handlers

import (
	"github.com/sqlite-go/internal/engine"
)

type Handler struct {
	DB *engine.DB
}

func New(db *engine.DB) *Handler {
	return &Handler{
		DB: db,
	}
}
