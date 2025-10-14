package entity

import (
	"time"

	seasonEntity "igaopk.com/goPower/internal/domain/season/entity"
)

type Episode struct {
	ID            int                 `json:"id" gorm:"primaryKey;autoIncrement"`
	SeasonID      int                 `json:"season_id"`
	Season        seasonEntity.Season `json:"season" gorm:"foreignKey:SeasonID;references:ID"`
	Title         string              `json:"title" gorm:"size:150;not null"`
	Synopsis      string              `json:"synopsis"`
	EpisodeNumber int                 `json:"episode_number"`
	AirDate       *time.Time          `json:"air_date"`
	IsKeyEpisode  bool                `json:"is_key_episode" gorm:"default:false"`
	ThumbnailURL  string              `json:"thumbnail_url"`

	// Relacionamentos via IDs
	RangerIDs []int `json:"ranger_ids,omitempty" gorm:"-"`
	CameoIDs  []int `json:"cameo_ids,omitempty" gorm:"-"`
}

func (Episode) TableName() string {
	return "episodes"
}
