package models

type Picture struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	RangerID    uint   `gorm:"not null"`
	URL         string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`

	// Relacionamento inverso (opcional)
	Ranger Ranger `gorm:"foreignKey:RangerID"`
}
