package db

import "weather_api/internal/entities"

type ListStorage interface {
	CreateList(userId int, list entities.UserList) (id int, err error)
	GetListById(userId, listId int) (list entities.UserList, err error)
	GetAllList(userId int) (lists []entities.UserList, err error)
	UpdateList(userId, listId int, input entities.UserList) error
	DeleteList(userId int, title string) error
	CreateItem(listId int, city string) (id int, err error)
	GetAllItems(listId int) (items []entities.Item, err error)
}
