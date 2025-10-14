package repository

import (
	"gorm.io/gorm"
	"igaopk.com/goPower/internal/domain/season/entity"
)

type SeasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) *SeasonRepository {
	return &SeasonRepository{db: db}
}

func (r *SeasonRepository) getAllSeasons() ([]entity.Season, error) {
	var seasons []entity.Season

	err := r.db.Preload("Rangers").Find(&seasons).Error

	if err != nil {
		panic(err)
	}

	return seasons, nil
}
