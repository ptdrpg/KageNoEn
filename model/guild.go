package model

import "time"

type Guild struct {
	Id        string    `gorm:"id;primarykey" json:"id"`
	Name      string    `gorm:"name;unique" json:"name"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
