package entity

type RangerPhoto struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	RangerID int    `json:"ranger_id"`
	PhotoURL string `json:"photo_url" gorm:"not null"`
	Caption  string `json:"caption"`
}

func (RangerPhoto) TableName() string {
	return "ranger_photos"
}
