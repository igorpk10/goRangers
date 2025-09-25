package repository

import (
	"gorm.io/gorm"
	"igaopk.com/goPower/internal/models"
)

type RangerRepository interface {
	FindAll() ([]models.Ranger, error)
}

type rangerGormRepository struct {
	db *gorm.DB
}

func NewRangerGormRepository(db *gorm.DB) RangerRepository {
	return &rangerGormRepository{db: db}
}

func (r *rangerGormRepository) FindAll() ([]models.Ranger, error) {
	var rangers []models.Ranger
	err := r.db.Preload("Pictures").Find(&rangers).Error
	return rangers, err
}
