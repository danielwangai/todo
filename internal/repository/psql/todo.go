package psql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	repo "github.com/danielwangai/todo-app/internal/repository"
)

func (dao *dbClient) CreateTodoItem(ctx context.Context, item *repo.ItemModelType) (*repo.ItemModelType, error) {
	// check if user already has an item with the same name
	_, err := dao.findTodoByNameAndUserId(ctx, item.Name, item.UserId)
	// if err != sql.ErrNoRows && err != nil {
	if err == nil {
		dao.log.Infof("user of id: %d already has an item with name: %s", item.UserId, item.Name)
		return nil, errors.New("item matching user id and name not found")
	}
	var id int

	query := `
	INSERT INTO items (name, user_id, created_at)
	VALUES($1, $2, $3) RETURNING id`
	err = dao.db.QueryRow(query, item.Name, item.UserId, item.CreatedAt).Scan(&id)
	if err != nil {
		dao.log.WithError(err).Error("an error occurred when inserting item to the database")
		return nil, err
	}

	item.ID = id

	dao.log.Infof("DB-OP: item created successfully: %v", item)
	return item, nil
}

func (dao *dbClient) findTodoByNameAndUserId(ctx context.Context, name string, userId int) (*repo.ItemModelType, error) {
	var itemId int
	var createdAt time.Time
	var updatedAt sql.NullTime
	query := `SELECT id, created_at, updated_at FROM items WHERE name=$1 AND user_id=$2`
	row := dao.db.QueryRow(query, name, userId)
	err := row.Scan(&itemId, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	return &repo.ItemModelType{
		ID:        itemId,
		Name:      name,
		UserId:    userId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
