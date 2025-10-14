package entity

type RangerGeneratedImage struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	RangerID    int    `json:"ranger_id"`
	ImageURL    string `json:"image_url" gorm:"not null"`
	Description string `json:"description"`
}

func (RangerGeneratedImage) TableName() string {
	return "ranger_generated_images"
}
