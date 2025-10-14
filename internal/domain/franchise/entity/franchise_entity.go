package entity

import (
	genreEntity "igaopk.com/goPower/internal/domain/genres/entity"
)

type Franchise struct {
	ID             int                `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string             `json:"name" gorm:"size:150;not null"`
	GenreID        *int               `json:"genre_id"`
	Genre          *genreEntity.Genre `json:"genre" gorm:"foreignKey:GenreID;references:ID"`
	YoutubeChannel string             `json:"youtube_channel"`
	Description    string             `json:"description"`
	LogoURL        string             `json:"logo_url"`
}

func (Franchise) TableName() string {
	return "franchises"
}
