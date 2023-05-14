package activity

import (
	"go_todo/pkg/entities"

	"gorm.io/gorm"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateActivity(activity *entities.Activity) (*entities.Activity, error)
	GetActivities() (*[]entities.Activity, error)
	GetActivity(id uint) (*entities.Activity, error)
	DeleteActivity(id uint) error
	UpdateActivity(activity *entities.Activity) (*entities.Activity, error)
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

func (r *repository) UpdateActivity(activity *entities.Activity) (*entities.Activity, error) {
	activityupdate := new(entities.Activity)
	activityupdate.ID = activity.ID
	err := r.db.First(activityupdate).Error
	if err != nil {
		return nil, err
	}
	activityupdate.Title = activity.Title
	activityupdate.Email = activity.Email
	id := r.db.Save(activityupdate)
	errupdate := id.Error
	if errupdate != nil {
		return nil, errupdate
	}
	return activity, nil
}

// ReadBook is a mongo repository that helps to fetch books
func (r *repository) GetActivity(id uint) (*entities.Activity, error) {
	var activity entities.Activity
	result := r.db.First(&activity, id)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *repository) GetActivities() (*[]entities.Activity, error) {
	var activity []entities.Activity
	result := r.db.Find(&activity)
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
	errcheck := r.db.First(&activity).Error
	if errcheck != nil {
		return errcheck
	}
	err := r.db.Delete(&activity, id).Error
	if err != nil {
		return err
	}
	return nil
}
