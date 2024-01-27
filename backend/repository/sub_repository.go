package repository

import (
	"context"
	
	_ "github.com/lib/pq"
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
func (tr *eventRepository) Delete(c context.Context, eventId int, userId int) error {
	return nil
}
func (tr *eventRepository) Update(c context.Context, eventId int, userId int, event *domain.Event) error {
	return nil
}
func (tr *eventRepository) GetById(c context.Context, eventId int) (domain.Event, error) {
	var data domain.Event
	return data, nil 
}
func (tr *eventRepository) GetAll(c context.Context) ([]domain.Event, error){
	var data []domain.Event
	
	return data, nil
}
