package repository

import (
	"time"

	"sosmed/features/comments"
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

func (mod *Comment) FromEntity(ent comments.Comment) {
	if ent.Text != "" {
		mod.Text = ent.Text
	}

	if ent.PostId != 0 {
		mod.PostId = ent.PostId
	}

	if ent.User.Id != 0 {
		mod.UserId = ent.User.Id
	}
}
