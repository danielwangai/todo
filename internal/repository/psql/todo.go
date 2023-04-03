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

func (dao *dbClient) GetAllTodoItems(ctx context.Context, id int) ([]*repo.ItemModelType, error) {
	items := []*repo.ItemModelType{}
	rows, err := dao.db.Query("SELECT id, name, user_id, is_deleted, created_at, updated_at from items WHERE is_deleted=FALSE")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		i := repo.ItemModelType{}
		err = rows.Scan(&i.ID, &i.Name, &i.UserId, &i.IsDeleted, &i.CreatedAt, &i.UpdatedAt)
		if err != nil {
			dao.log.WithError(err).Error("error scanning get items result")
			return nil, err
		}
		items = append(items, &i)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (dao *dbClient) FindTodoItemById(ctx context.Context, id int) (*repo.ItemModelType, error) {
	var i repo.ItemModelType
	query := `SELECT id, name, user_id, is_deleted, created_at, updated_at from items WHERE id=$1`
	err := dao.db.QueryRow(query, id).
		Scan(&i.ID, &i.Name, &i.UserId, &i.IsDeleted, &i.CreatedAt, &i.UpdatedAt)
	if err != nil {
		return nil, err
	}

	dao.log.Infof("DB-OP: successfully fetched todo item: %v by ID: %v", i, id)
	return &i, nil
}

func (dao *dbClient) DeleteTodoItemById(ctx context.Context, id int) error {
	_, err := dao.FindTodoItemById(ctx, id)
	if err != nil {
		return err
	}

	_, err = dao.db.Exec(`UPDATE items SET is_deleted=TRUE WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}
