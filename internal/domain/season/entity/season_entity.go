package entity

import (
	franchiseEntity "igaopk.com/goPower/internal/domain/franchise/entity"
)

type Season struct {
	ID                int                        `json:"id" gorm:"primaryKey;autoIncrement"`
	FranchiseID       *int                       `json:"franchise_id"`
	Franchise         *franchiseEntity.Franchise `json:"franchise" gorm:"foreignKey:FranchiseID;references:ID"`
	Title             string                     `json:"title" gorm:"size:150;not null"`
	Subtitle          string                     `json:"subtitle"`
	Description       string                     `json:"description"`
	ImageURL          string                     `json:"image_url"`
	ReleaseYear       *int                       `json:"release_year"`
	NumEpisodes       *int                       `json:"num_episodes"`
	ThemeSongURL      string                     `json:"theme_song_url"`
	OpeningCreditsURL string                     `json:"opening_credits_url"`

	// Relacionamentos via IDs
	RangerIDs  []int `json:"ranger_ids,omitempty" gorm:"-"`
	ZordIDs    []int `json:"zord_ids,omitempty" gorm:"-"`
	VillainIDs []int `json:"villain_ids,omitempty" gorm:"-"`
	TriviaIDs  []int `json:"trivia_ids,omitempty" gorm:"-"`
}

func (Season) TableName() string {
	return "seasons"
}
