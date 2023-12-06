package repository

import (
	"context"
	"errors"
	"sosmed/features/comments"

	"gorm.io/gorm"
)

func NewCommentRepository(mysqlDB *gorm.DB) comments.Repository {
	return &commentRepository{
		mysqlDB: mysqlDB,
	}
}

type commentRepository struct {
	mysqlDB *gorm.DB
}

func (repo *commentRepository) Create(ctx context.Context, data comments.Comment) error {
	var mod = new(Comment)
	mod.FromEntity(data)

	if err := repo.mysqlDB.WithContext(ctx).Create(mod).Error; err != nil {
		return err
	}

	return nil
}

func (repo *commentRepository) Delete(ctx context.Context, commentId uint) error {
	qry := repo.mysqlDB.WithContext(ctx).Delete(&Comment{Id: commentId})
	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil

	return nil
}
