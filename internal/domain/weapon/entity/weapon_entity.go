package entity

type Weapon struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"size:150"`
	RangerID    int    `json:"ranger_id"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func (Weapon) TableName() string {
	return "weapons"
}
