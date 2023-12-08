package repository

import (
	"context"
	"errors"
	"sosmed/features/posts"
	"sosmed/helpers/filters"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

func NewPostRepository(mysqlDB *gorm.DB, cloudinary *cloudinary.Cloudinary) posts.Repository {
	return &postRepository{
		mysqlDB:    mysqlDB,
		cloudinary: cloudinary,
	}
}

type postRepository struct {
	mysqlDB    *gorm.DB
	cloudinary *cloudinary.Cloudinary
}

func (repo *postRepository) Create(ctx context.Context, data posts.Post) error {
	for i := 0; i < len(data.Attachment); i++ {
		UniqueFilename := true

		res, err := repo.cloudinary.Upload.Upload(ctx, data.Attachment[i].Raw, uploader.UploadParams{
			UniqueFilename: &UniqueFilename,
			Folder:         "posts",
		})

		if err != nil {
			return err
		}

		data.Attachment[i].URL = res.URL
	}

	var mod = new(Post)
	mod.FromEntity(data)

	if err := repo.mysqlDB.WithContext(ctx).Create(mod).Error; err != nil {
		return err
	}

	return nil
}

func (repo *postRepository) GetById(ctx context.Context, postId uint) (*posts.Post, error) {
	var result = new(Post)

	if err := repo.mysqlDB.WithContext(ctx).Preload("User").Preload("Comment").Preload("Comment.User").Preload("Attachment").First(result, postId).Error; err != nil {
		return nil, err
	}

	return result.ToEntity(), nil
}

func (repo *postRepository) GetList(ctx context.Context, filter filters.Filter, userId *uint) ([]posts.Post, int, error) {
	var data []Post
	var totalData = new(int64)

	qry := repo.mysqlDB.WithContext(ctx).Model(&Post{})

	if filter.Search.Keyword != "" {
		qry = qry.Where("caption LIKE ?", "%"+filter.Search.Keyword+"%")
	}

	if userId != nil {
		qry = qry.Where(&Post{UserId: *userId})
	}

	qry.Count(totalData)

	qry = qry.Preload("User").Preload("Comment").Preload("Attachment")
	qry = qry.Limit(filter.Pagination.Limit).Offset(filter.Pagination.Start)

	if err := qry.Order("id DESC").Find(&data).Error; err != nil {
		return nil, 0, err
	}

	var result []posts.Post
	for _, mod := range data {
		result = append(result, *mod.ToEntity())
	}

	return result, int(*totalData), nil
}

func (repo *postRepository) Update(ctx context.Context, postId uint, data posts.Post) error {
	for i := 0; i < len(data.Attachment); i++ {
		UniqueFilename := true

		res, err := repo.cloudinary.Upload.Upload(ctx, data.Attachment[i].Raw, uploader.UploadParams{
			UniqueFilename: &UniqueFilename,
			Folder:         "posts",
		})

		if err != nil {
			return err
		}

		data.Attachment[i].URL = res.URL
	}

	var updateData = new(Post)
	updateData.FromEntity(data)

	qry := repo.mysqlDB.Model(&Post{Id: postId}).Updates(updateData)
	if qry.Error != nil {
		return qry.Error
	}

	if len(data.Attachment) != 0 {
		if err := repo.mysqlDB.Model(&Post{Id: postId}).Association("Attachment").Replace(updateData.Attachment); err != nil {
			return err
		}
	}

	return nil
}

func (repo *postRepository) Delete(ctx context.Context, postId uint) error {
	qry := repo.mysqlDB.WithContext(ctx).Delete(&Post{Id: postId})
	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

func (repo *postRepository) DeleteByUserId(ctx context.Context, userId uint) error {
	qry := repo.mysqlDB.WithContext(ctx).Where(&Post{UserId: userId}).Delete(&Post{})
	if qry.Error != nil {
		return qry.Error
	}

	return nil
}
