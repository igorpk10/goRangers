package entity

type Cameo struct {
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	RangerID  int    `json:"ranger_id"`
	EpisodeID int    `json:"episode_id"`
	CameoType string `json:"cameo_type"`
}

func (Cameo) TableName() string {
	return "cameos"
}
