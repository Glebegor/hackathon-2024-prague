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

func (tr *eventRepository) Create(c context.Context, event *domain.Event) error {
	return nil
}
func (tr *eventRepository) Delete(c context.Context) error {
	return nil
}
func (tr *eventRepository) Update(c *gin.Context, eventId string) error {
	return nil
}
func (tr *eventRepository) GetById(c *gin.Context, eventId string) () {
	var data domain.Event
	return data, nil 
}
func (tr *eventRepository) GetAll(c *gin.Context) ([]domain.Event, error){
	var data []domain.Event
	return data, nil
}
