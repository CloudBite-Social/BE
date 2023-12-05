package repository

import (
	"time"

	ur "sosmed/features/users/repository"

	"gorm.io/gorm"
)

type Comment struct {
	Id   uint   `gorm:"column:id; primaryKey;"`
	Text string `gorm:"column:text; type:text;"`

	UserId uint
	User   ur.User `gorm:"foreignKey:UserId"`

	PostId uint

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
