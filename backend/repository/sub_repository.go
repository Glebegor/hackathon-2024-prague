package repository

import (
	"context"
	"fmt"
	"strings"

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
	query := fmt.Sprintf("DELETE FROM %s WHERE channel_id=$1 AND user_id=$2", tr.SubTable)
	_, err := tr.database.Exec(query, subId, userId)
	if err != nil {
		return err
	}
	return nil
}
func (tr *subRepository) Update(c context.Context, subId int, sub *domain.UpdateSub) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if sub.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name='%v'", sub.Name))
		args = append(args, fmt.Sprintf("$%d", argId))
		argId++
	}
	if sub.Price > 0 {
		setValues = append(setValues, fmt.Sprintf("price=%v", sub.Price))
		args = append(args, fmt.Sprintf("$%d", argId))
		argId++
	}
	if sub.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description='%v'", sub.Description))
		args = append(args, fmt.Sprintf("$%d", argId))
		argId++
	}
	if sub.Images != "" {
		setValues = append(setValues, fmt.Sprintf("images='%v'", sub.Images))
		args = append(args, fmt.Sprintf("$%d", argId))
		argId++
	}
	if sub.Tags != "" {
		setValues = append(setValues, fmt.Sprintf("tags='%v'", sub.Tags))
		args = append(args, fmt.Sprintf("$%d", argId))
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE channel_id=%d", tr.SubTable, setQuery, subId)
	fmt.Print(query)
	_, err := tr.database.Exec(query)
	return err
}
func (tr *subRepository) GetById(c context.Context, subId int) (domain.Sub, error) {
	var data domain.Sub
	query := fmt.Sprintf("SELECT * FROM %s WHERE channel_id=$1", tr.SubTable)
	err := tr.database.QueryRow(query, subId).Scan(&data.ChannelId, &data.Name, &data.Price, &data.Description, &data.UserId, &data.Link, &data.Images, &data.Tags)
	return data, err
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
