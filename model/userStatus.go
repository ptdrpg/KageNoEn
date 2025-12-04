package model

type UserStatus struct {
	Id        string    `gorm:"id;primaryKey" json:"id"`
	Label     string    `gorm:"label" json:"label"`
}

type UserStatusList struct {
	Data []UserStatus `json:"data"`
}

type UserStatusResponse struct {
	Data UserStatus `json:"data"`
}
