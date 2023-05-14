package activity

import (
	"go_todo/api/presenter"
	"go_todo/pkg/entities"

	"gorm.io/gorm"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateActivity(activity *entities.Activity) (*entities.Activity, error)
	// GetActivities() (*[]presenter.Book, error)
	GetActivity(id uint) (*presenter.Activity, error)
	DeleteActivity(id uint) error
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
func (r *repository) CreateActivity(activity *entities.Activity) (*entities.Activity, error) {
	id := r.db.Create(&activity)
	err := id.Error
	if err != nil {
		return nil, err
	}
	return activity, nil
}

// ReadBook is a mongo repository that helps to fetch books
func (r *repository) GetActivity(id uint) (*presenter.Activity, error) {
	var activity presenter.Activity
	result := r.db.First(&activity, id)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

// DeleteBook is a mongo repository that helps to delete books
func (r *repository) DeleteActivity(id uint) error {
	var activity entities.Activity
	activity.ID = id
	err := r.db.Delete(&activity).Error
	if err != nil {
		return err
	}
	return nil
}
