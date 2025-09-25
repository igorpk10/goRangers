package models

type Picture struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	RangerID    uint   `gorm:"not null"`
	URL         string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`

	Ranger Ranger `gorm:"foreignKey:RangerID" json:"-"`
}

func (Picture) TableName() string {
	return "picture"
}
