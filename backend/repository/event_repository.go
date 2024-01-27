package repository

import (
	"context"

	"ton-event-bot/domain"
	"github.com/jmoiron/sqlx"
)

type eventRepository struct {
	database *sqlx.DB
	collection string
}

func NewEventRepository(db *sqlx.DB, collection string) domain.EventRepository {
	return &eventRepository{
		database: db,
		collection: collection,
	}
}


