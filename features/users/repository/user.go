package repository

import (
	"time"

	users "command-line-argumentsD:\\ALTA-GO\\BE\\features\\users\\entity.go"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Id       uint   `gorm:"column:id; primaryKey;"`
	Name     string `gorm:"column:name; type:varchar(200);"`
	Email    string `gorm:"column:email; type:varchar(255); unique;"`
	Password string `gorm:"column:password; type:varchar(72); not null;"`
	Image    string `gorm:"column:image; type:text;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type userquery struct {
	db *gorm.DB
}

func new(db *gorm.DB) users.Repository  {
	return &userquery {
		db : db
	}
}

func (uq *UserQuery) Register(newUser UserModel) bool  {
	if err := uq.DB.Create(&newUser).Error; err != nill {
		return false
	}
}

func (uq *userquery) Insert(newUser users.User) (users.user, error) {
	var inputDB = new(UserModel)
	inputDB.email = newUser.email
	inputDB.Nama = newUser.Nama
	inputDB.Password = newUser.Password
	
	newUser.ID = inputDB.ID
	if err := uq.db.Create(&inputDB).Error; err != nill {
		return users.user{}, err
	}
	
	newUser.ID = inputDB.ID

	return newUser, nil
}
