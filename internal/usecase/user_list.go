package usecase

import (
	"weather_api/internal/adapters/db"
	"weather_api/internal/entities"
)

type listUseCase struct {
	storage db.ListStorage
}

func NewListUseCase(storage db.ListStorage) *listUseCase {
	return &listUseCase{
		storage: storage,
	}
}

func (l *listUseCase) CreateList(userId int, list entities.UserList) (id int, err error) {
	return l.storage.CreateList(userId, list)
}

func (l *listUseCase) GetListById(userId, listId int) (list entities.UserList, err error) {
	return l.storage.GetListById(userId, listId)
}

func (l *listUseCase) GetAllList(userId int) (lists []entities.UserList, err error) {
	return l.storage.GetAllList(userId)
}

func (l *listUseCase) UpdateList(userId, listId int, input entities.UserList) error {
	return l.storage.UpdateList(userId, listId, input)
}

func (l *listUseCase) DeleteList(userId int, title string) error {
	return l.storage.DeleteList(userId, title)
}
