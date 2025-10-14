package entity

import (
	seasonEntity "igaopk.com/goPower/internal/domain/season/entity"
)

type Villain struct {
	ID          int                 `json:"id" gorm:"primaryKey;autoIncrement"`
	SeasonID    int                 `json:"season_id"`
	Season      seasonEntity.Season `json:"season" gorm:"foreignKey:SeasonID;references:ID"`
	Name        string              `json:"name" gorm:"size:150;not null"`
	Description string              `json:"description"`
	MainVillain bool                `json:"main_villain" gorm:"default:false"`
	ActorName   string              `json:"actor_name" gorm:"size:150"`
	ImageURL    string              `json:"image_url"`
}

func (Villain) TableName() string {
	return "villains"
}
