package todo

import (
	"go_todo/pkg/entities"
)

// Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertTodo(book *entities.Todo) (*entities.Todo, error)
	GetTodo(id uint) (*entities.Todo, error)
	RemoveTodo(id uint) error
	GetTodos() (*[]entities.Todo, error)
	UpdateTodo(Todo *entities.Todo) (*entities.Todo, error)
	GetTodosByActivity(activity_id uint) (*[]entities.Todo, error)
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
func (s *service) InsertTodo(Todo *entities.Todo) (*entities.Todo, error) {
	return s.repository.CreateTodo(Todo)
}
func (s *service) UpdateTodo(Todo *entities.Todo) (*entities.Todo, error) {
	return s.repository.UpdateTodo(Todo)
}

// FetchBooks is a service layer that helps fetch all books in BookShop
func (s *service) GetTodo(id uint) (*entities.Todo, error) {
	return s.repository.GetTodo(id)
}

func (s *service) GetTodos() (*[]entities.Todo, error) {
	return s.repository.GetTodos()
}

func (s *service) GetTodosByActivity(activity_id uint) (*[]entities.Todo, error) {
	return s.repository.GetTodosByactivity(activity_id)
}

// RemoveBook is a service layer that helps remove books from BookShop
func (s *service) RemoveTodo(id uint) error {
	return s.repository.DeleteTodo(id)
}
