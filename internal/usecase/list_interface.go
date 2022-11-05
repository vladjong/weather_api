package usecase

import "weather_api/internal/entities"

type List interface {
	CreateList(userId int, list entities.UserList) (id int, err error)
	GetListById(userId, listId int) (list entities.UserList, err error)
	GetAllList(userId int) (lists []entities.UserList, err error)
	UpdateList(userId, listId int, input entities.UserList) error
	DeleteList(userId int, title string) error
	CreateItem(userId, listId int, city string) (id int, err error)
	GetAllItems(userId, listId int) (items []entities.Item, err error)
}
