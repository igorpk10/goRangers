package models

import (
	"time"
)

type Ranger struct {
	ID              uint       `gorm:"primaryKey;autoIncrement"`
	Name            string     `gorm:"type:varchar(100);not null"`
	MorphedName     string     `gorm:"type:varchar(100)"`
	Color           string     `gorm:"type:varchar(50)"`
	Role            string     `gorm:"type:varchar(100)"`
	Team            string     `gorm:"type:varchar(100)"`
	Season          int        `gorm:"type:varchar(100)"`
	Zord            string     `gorm:"type:varchar(100)"`
	Weapon          string     `gorm:"type:varchar(100)"`
	Morpher         string     `gorm:"type:varchar(100)"`
	ActorName       string     `gorm:"type:varchar(100)"`
	PlanetOrigin    string     `gorm:"type:varchar(100)"`
	FirstAppearance *time.Time `gorm:"type:date"` 
	Age             *int       `gorm:"type:int"`  
	IsLeader        bool       `gorm:"default:false"`
	Status          string     `gorm:"type:varchar(50);default:'Ativo'"`
	Bio             string     `gorm:"type:text"`

	// Relacionamento 1:N com Picture
	Pictures []Picture `gorm:"foreignKey:RangerID;constraint:OnDelete:CASCADE"`
}

func (Ranger) TableName() string {
	return "ranger"
}
