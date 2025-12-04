package model

type Rank struct {
	Id     string `gorm:"id;primarykey" json:"id"`
	Label  string `gorm:"label" json:"label"`
	EloMin int    `gorm:"elo_min" json:"elo_min"`
	EloMax int    `gorm:"elo_max" json:"elo_max"`
}

type RankList struct {
	Data []Rank `json:"data"`
}

type RankResponse struct {
	Data Rank `json:"data"`
}
