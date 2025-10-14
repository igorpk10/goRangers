package repository

import (
	"gorm.io/gorm"
	"igaopk.com/goPower/internal/domain/ranger/entity"
)

type RangerRepository struct {
	db *gorm.DB
}

func NewRangerRepository(db *gorm.DB) *RangerRepository {
	return &RangerRepository{db: db}
}

func (r *RangerRepository) getRangersBySeasonID(id int32) ([]entity.Ranger, error) {
	var rangers []entity.Ranger

	err := r.db.
		Preload("Photos").
		Preload("GeneratedImages").
		Preload("Weapons").
		Preload("Cameos").
		Order("rangers.id ASC").
		Joins("JOIN season_rangers ON season_rangers.ranger_id = rangers.id").
		Where("season_rangers.season_id = ?", id).
		Find(&rangers).Error

	if err != nil {
		panic(err)
	}

	return rangers, nil
}
