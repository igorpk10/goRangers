package entity

type SeasonTrivia struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	SeasonID int    `json:"season_id"`
	Fact     string `json:"fact"`
}

func (SeasonTrivia) TableName() string {
	return "season_trivia"
}
