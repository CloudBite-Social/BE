package repository

import (
	"context"
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
	panic("unimplemented")
}

func (repo *postRepository) GetList(ctx context.Context, filter filters.Filter, userId *uint) ([]posts.Post, error) {
	panic("unimplemented")
}

func (repo *postRepository) Update(ctx context.Context, postId uint, data posts.Post) error {
	panic("unimplemented")
}

func (repo *postRepository) Delete(ctx context.Context, postId uint) error {
	panic("unimplemented")
}
