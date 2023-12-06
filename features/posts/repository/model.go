package repository

import (
	"time"

	cr "sosmed/features/comments/repository"
	"sosmed/features/posts"
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

func (mod *Post) FromEntity(ent posts.Post) {
	if ent.Caption != "" {
		mod.Caption = ent.Caption
	}

	if ent.User.Id != 0 {
		mod.UserId = ent.User.Id
	}

	if len(ent.Attachment) != 0 {
		for _, file := range ent.Attachment {
			var modFile = new(File)
			modFile.FromEntity(file)
			mod.Attachment = append(mod.Attachment, *modFile)
		}
	}
}

type File struct {
	Id  uint   `gorm:"column:id; primaryKey;"`
	URL string `gorm:"column:url; type:text;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (mod *File) FromEntity(ent posts.File) {
	if ent.URL != "" {
		mod.URL = ent.URL
	}
}
