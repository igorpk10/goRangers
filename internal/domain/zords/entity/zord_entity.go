package entity

type Zord struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	SeasonID    int    `json:"season_id"`
	Name        string `json:"name" gorm:"size:150;not null"`
	Type        string `json:"type" gorm:"size:100"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func (Zord) TableName() string {
	return "zords"
}
