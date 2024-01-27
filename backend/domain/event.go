package domain

import (
	"context"


)

const (
	CollectionEvent = "events"
)

type Event struct {
	Id           int                  `json:"id"`
	Name         string               `json:"name"`
	Id_of_person string               `jsom:"id_of_person"`
	tags         string               `json:"tags"`
	Description  string               `json:"description"`
	Images       string               `json:"images"`
	Start-date   string               `json:"start-date"`
	End-date     string               `json:"end-date"`
	Start-time   string               `json:"start-time"`
}

type EventRepository interface {
	
}

type TaskUsercase interface {
	
}
