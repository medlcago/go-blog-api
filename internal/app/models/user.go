package models

type User struct {
	Base
	Username string `gorm:"not null;unique;size:16" json:"username" binding:"ascii,required,min=5,max=20"`
	Password string `gorm:"not null;size:128" json:"-" binding:"min=6,max=64"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
	Posts    []Post `json:"posts,omitempty"`
}
