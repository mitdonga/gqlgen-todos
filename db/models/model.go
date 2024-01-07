package models

import "time"

type Model struct {
	ID        int        `gorm:"primary_key;type:serial" json:"id" deep:"-"`
	CreatedAt time.Time  `json:"created_at" deep:"-"`
	UpdatedAt time.Time  `json:"updated_at" deep:"-"`
	DeletedAt *time.Time `json:"deleted_at" deep:"-"`
}

type Todo struct {
	Model
	Text string `json:"text"`
	Done bool   `json:"done"`
}
