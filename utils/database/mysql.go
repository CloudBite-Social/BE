package database

import (
	"fmt"
	"sosmed/config"
	cr "sosmed/features/comments/repository"
	pr "sosmed/features/posts/repository"
	ur "sosmed/features/users/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlInit(cfg config.Database) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MysqlMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&ur.User{},
		&pr.Post{},
		&cr.Comment{},
	)

	if err != nil {
		return err
	}

	return nil
}
