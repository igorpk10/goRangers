package entity

type Genre struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name" gorm:"size:100;not null"`
	ImageURL   string `json:"image_url"`
	ColorTheme string `json:"color_theme"`
}

func (Genre) TableName() string {
	return "genres"
}
