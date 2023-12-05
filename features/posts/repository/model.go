package repository

import (
	"time"

	cr "sosmed/features/comments/repository"
	ur "sosmed/features/users/repository"

	"gorm.io/gorm"
)

type Post struct {
	Id      uint   `gorm:"column:id; primaryKey;"`
	Caption string `gorm:"column:caption; type:text; index;"`

	UserId uint
	User   ur.User `gorm:"foreignKey:UserId"`

	Comment []cr.Comment

	Attachment []File `gorm:"many2many:attachment"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type File struct {
	Id   uint   `gorm:"column:id; primaryKey;"`
	Path string `gorm:"column:path; type:text;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
