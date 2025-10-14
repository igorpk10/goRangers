package entity

import (
	cameos "igaopk.com/goPower/internal/domain/cameo/entity"
	seasonEntity "igaopk.com/goPower/internal/domain/season/entity"
	weaponEntity "igaopk.com/goPower/internal/domain/weapon/entity"
)

type Ranger struct {
	ID             int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string `json:"name" gorm:"size:150;not null"`
	ActorName      string `json:"actor_name" gorm:"size:150"`
	Color          string `json:"color" gorm:"size:50"`
	Bio            string `json:"bio"`
	Trivia         string `json:"trivia"`
	DebutEpisodeID *int   `json:"debut_episode_id"` // apenas ID
	Status         string `json:"status" gorm:"size:50"`

	Photos          []RangerPhoto          `json:"photos,omitempty" gorm:"foreignKey:RangerID"`
	GeneratedImages []RangerGeneratedImage `json:"generated_images,omitempty" gorm:"foreignKey:RangerID"`
	Seasons         []seasonEntity.Season  `json:"seasons,omitempty" gorm:"many2many:season_rangers"`
	Weapons         []weaponEntity.Weapon  `json:"weapons,omitempty" gorm:"foreignKey:RangerID"`
	Cameos          []cameos.Cameo         `json:"cameos,omitempty" gorm:"foreignKey:RangerID"`
}

func (Ranger) TableName() string {
	return "rangers"
}
