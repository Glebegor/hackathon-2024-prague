package domain

import (
	"context"
)

const (
	CollectionEvent = "events"
)

type Event struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	IdOfPerson string `jsom:"id_of_person"`
	tags         string `json:"tags"`
	Description  string `json:"description"`
	Images       string `json:"images"`
	StartDate   string `json:"start-date"`
}

type EventRepository interface {
	Create (c context.Context, event *Event) error
	GetById (c context.Context, eventId int) (Event, error)
	GetAll (c context.Context) ([]Event, error)
	Update (c context.Context, eventId int, userId int, event *Event) error
	Delete (c context.Context, eventId int, userId int) error
}

type EventUsecase interface {
	Create (c context.Context, event *Event) error
       	GetById (c context.Context, eventId int) (Event, error)
       	GetAll (c context.Context) ([]Event, error)
       	Update (c context.Context, eventId int, userId int, event *Event) error
       	Delete (c context.Context, eventId int, userId int) error
}
