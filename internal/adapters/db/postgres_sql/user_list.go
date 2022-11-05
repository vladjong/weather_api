package postgressql

import (
	"fmt"
	"weather_api/config"
	"weather_api/internal/entities"

	"github.com/jmoiron/sqlx"
)

type listStorage struct {
	db *sqlx.DB
}

func NewListStorage(db *sqlx.DB) *listStorage {
	return &listStorage{
		db: db,
	}
}

func (r *listStorage) CreateList(userId int, list entities.UserList) (id int, err error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, title) VALUES ($1, $2) RETURNING id", config.UserListTable)
	row := r.db.QueryRow(query, userId, list.Title)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *listStorage) GetListById(userId, listId int) (list entities.UserList, err error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND id = $2", config.UserListTable)
	if err := r.db.Get(&list, query, userId, listId); err != nil {
		return list, err
	}
	return list, nil
}

func (r *listStorage) GetAllList(userId int) (lists []entities.UserList, err error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", config.UserListTable)
	err = r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *listStorage) UpdateList(userId, listId int, input entities.UserList) error {
	query := fmt.Sprintf("UPDATE %s SET title = $1 WHERE user_id = $2 AND id = $3", config.UserListTable)
	_, err := r.db.Exec(query, input.Title, userId, listId)
	return err
}

func (r *listStorage) DeleteList(userId int, title string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND title = $2", config.UserListTable)
	_, err := r.db.Exec(query, userId, title)
	return err
}

func (r *listStorage) CreateItem(listId int, city string) (id int, err error) {
	var cityId int
	queryCity := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", config.CitiesTable)
	if err := r.db.Get(&cityId, queryCity, city); err != nil {
		return 0, err
	}
	queryIten := fmt.Sprintf("INSERT INTO %s (city_id, list_id) VALUES ($1, $2) RETURNING id", config.ListItemsTable)
	row := r.db.QueryRow(queryIten, cityId, listId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *listStorage) GetAllItems(listId int) (items []entities.Item, err error) {
	query := fmt.Sprintf(`SELECT c.name, AVG(w.temp) AS av_temp
							FROM %s AS l
							JOIN %s AS c ON l.city_id = c.id
							JOIN %s AS w ON w.city_id = c.id
							WHERE l.list_id = $1
							GROUP BY c.name`, config.ListItemsTable, config.CitiesTable, config.WeathersTable)
	err = r.db.Select(&items, query, listId)
	return items, err
}
