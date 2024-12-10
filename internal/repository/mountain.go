package repository

import (
	"errors"
	"yamato/internal/domain"

	"gorm.io/gorm"
)

type MountainRepository interface {
	GetMountainByID(id int) (*domain.Mountain, error)
	GetAllMountains() ([]*domain.Mountain, error)
}

type mountainRepository struct {
	db *gorm.DB
}

func NewMountainRepository(db *gorm.DB) MountainRepository {
	return &mountainRepository{db: db}
}

func (r *mountainRepository) GetMountainByID(id int) (*domain.Mountain, error) {
	var mountain domain.Mountain
	result := r.db.First(&mountain, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("mountain not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &mountain, nil
}

func (r *mountainRepository) GetAllMountains() ([]*domain.Mountain, error) {
	var mountains []*domain.Mountain
	result := r.db.Find(&mountains)
	if result.Error != nil {
		return nil, result.Error
	}
	return mountains, nil
}
