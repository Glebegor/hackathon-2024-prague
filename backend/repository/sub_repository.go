package repository

import (
	"context"

	"ton-event-bot/domain"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type subRepository struct {
	database   *sqlx.DB
	collection string
}

func NewSubRepository(db *sqlx.DB, collection string) domain.SubRepository {
	return &subRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *subRepository) Create(c context.Context, sub *domain.Sub) error {
	return nil
}
func (tr *subRepository) Delete(c context.Context, subId int, userId int) error {
	return nil
}
func (tr *subRepository) Update(c context.Context, subId int, userId int, sub *domain.Sub) error {
	return nil
}
func (tr *subRepository) GetById(c context.Context, subId int) (domain.Sub, error) {
	var data domain.Sub
	return data, nil
}
func (tr *subRepository) GetAll(c context.Context) ([]domain.Sub, error) {
	var data []domain.Sub

	return data, nil
}
