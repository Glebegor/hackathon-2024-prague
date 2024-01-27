package repository

import (
	"context"
	"fmt"

	"ton-event-bot/domain"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type subRepository struct {
	database   *sqlx.DB
	collection string
	SubTable   string
	SubUser    string
}

func NewSubRepository(db *sqlx.DB, collection string) domain.SubRepository {
	return &subRepository{
		database:   db,
		collection: collection,
		SubTable:   "sub",
		SubUser:    "sub_users",
	}
}

func (tr *subRepository) Create(c context.Context, sub *domain.Sub) error {
	query := fmt.Sprintf("INSERT INTO %s (channel_id, price, name, description, user_id, link, images, tags) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", tr.SubTable)
	_, err := tr.database.Exec(query, sub.ChannelId, sub.Price, sub.Name, sub.Description, sub.UserId, sub.Link, sub.Images, sub.Tags)
	if err != nil {
		return err
	}
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

	query := fmt.Sprintf("SELECT * from %s", tr.SubTable)
	rows, err := tr.database.Query(query)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var dataEl domain.Sub
		err := rows.Scan(&dataEl.ChannelId, &dataEl.Name, &dataEl.Price, &dataEl.Description, &dataEl.UserId, &dataEl.Link, &dataEl.Images, &dataEl.Tags)
		if err != nil {
			return data, err
		}
		data = append(data, dataEl)
	}
	return data, nil
}
