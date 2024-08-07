package models

type Post struct {
	Base
	Title   string `gorm:"type:varchar(128);not null" json:"title"`
	Content string `gorm:"type:varchar(1024);not null" json:"content"`
	UserID  uint64 `gorm:"not null" json:"-"`
	User    *User  `json:"author,omitempty"`
}
