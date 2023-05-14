package activity

import (
	"go_todo/api/presenter"
	"go_todo/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertActivity(book *entities.Activity) (*entities.Activity, error)
	GetActivity(id uint) (*presenter.Activity, error)
	RemoveActivity(id uint) error
	GetActivities() (*[]entities.Activity, error)
}

type service struct {
	repository Repository
}

// NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// InsertBook is a service layer that helps insert book in BookShop
func (s *service) InsertActivity(activity *entities.Activity) (*entities.Activity, error) {
	return s.repository.CreateActivity(activity)
}

// FetchBooks is a service layer that helps fetch all books in BookShop
func (s *service) GetActivity(id uint) (*presenter.Activity, error) {
	return s.repository.GetActivity(id)
}
func (s *service) GetActivities() (*[]entities.Activity, error) {
	return s.repository.GetActivities()
}

// RemoveBook is a service layer that helps remove books from BookShop
func (s *service) RemoveActivity(id uint) error {
	return s.repository.DeleteActivity(id)
}
