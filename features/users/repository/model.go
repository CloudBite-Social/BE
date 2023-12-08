package repository

import (
	"sosmed/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"column:id; primaryKey;"`
	Name     string `gorm:"column:name; type:varchar(200);"`
	Email    string `gorm:"column:email; type:varchar(255); unique;"`
	Password string `gorm:"column:password; type:varchar(72); not null;"`
	Image    string `gorm:"column:image; type:text;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (mod *User) FromEntity(ent users.User) {
	if ent.Name != "" {
		mod.Name = ent.Name
	}

	if ent.Email != "" {
		mod.Email = ent.Email
	}

	if ent.Password != "" {
		mod.Password = ent.Password
	}

	if ent.Image != "" {
		mod.Image = ent.Image
	}
}
