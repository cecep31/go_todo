package todo

import (
	"fmt"
	"go_todo/api/presenter"
	"go_todo/pkg/entities"

	"gorm.io/gorm"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateTodo(Todo *entities.Todo) (*entities.Todo, error)
	GetTodos() (*[]entities.Todo, error)
	GetTodo(id uint) (*presenter.Todo, error)
	DeleteTodo(id uint) error
	UpdateTodo(Todo *entities.Todo) (*entities.Todo, error)
	GetTodosByactivity(activity_id uint) (*[]entities.Todo, error)
}
type repository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewRepo(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// CreateBook is a mongo repository that helps to create books
func (r *repository) CreateTodo(Todo *entities.Todo) (*entities.Todo, error) {
	id := r.db.Create(&Todo)
	err := id.Error
	if err != nil {
		return nil, err
	}
	return Todo, nil
}

func (r *repository) UpdateTodo(todo *entities.Todo) (*entities.Todo, error) {
	todoupdate := new(entities.Todo)
	todoupdate.ID = todo.ID
	err := r.db.First(todoupdate).Error
	if err != nil {
		return nil, err
	}
	todoupdate.Title = todo.Title
	todoupdate.Is_active = todo.Is_active
	todoupdate.Priority = todo.Priority
	id := r.db.Save(todoupdate)
	errupdate := id.Error
	if errupdate != nil {
		return nil, errupdate
	}
	return todo, nil
}

// ReadBook is a mongo repository that helps to fetch books
func (r *repository) GetTodo(id uint) (*presenter.Todo, error) {
	var todo presenter.Todo
	result := r.db.First(&todo, id)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *repository) GetTodos() (*[]entities.Todo, error) {
	fmt.Println("tidak masuk filtering")
	var todo []entities.Todo
	result := r.db.Find(&todo)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *repository) GetTodosByactivity(activity_id uint) (*[]entities.Todo, error) {
	fmt.Println("masuk filtering")
	var todo []entities.Todo
	result := r.db.Where("activity_group_id = ?", activity_id).Find(&todo)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// DeleteBook is a mongo repository that helps to delete books
func (r *repository) DeleteTodo(id uint) error {
	var Todo entities.Todo
	Todo.ID = id
	err := r.db.Delete(&Todo).Error
	if err != nil {
		return err
	}
	return nil
}
