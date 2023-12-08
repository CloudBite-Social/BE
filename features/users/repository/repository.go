package repository

import (
	"context"
	"errors"
	"sosmed/features/users"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

func NewUserRepository(mysqlDB *gorm.DB, cloudinary *cloudinary.Cloudinary) users.Repository {
	return &userRepository{
		mysqlDB:    mysqlDB,
		cloudinary: cloudinary,
	}
}

type userRepository struct {
	mysqlDB    *gorm.DB
	cloudinary *cloudinary.Cloudinary
}

func (repo *userRepository) Register(ctx context.Context, data users.User) error {
	var mod = new(User)
	mod.FromEntity(data)

	if err := repo.mysqlDB.WithContext(ctx).Create(mod).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) Login(ctx context.Context, email string) (*users.User, error) {
	var mod = new(User)

	if err := repo.mysqlDB.WithContext(ctx).Where(&User{Email: email}).First(mod).Error; err != nil {
		return nil, err
	}

	return mod.ToEntity(), nil
}

func (repo *userRepository) GetById(ctx context.Context, id uint) (*users.User, error) {
	panic("unimplemented")
}

func (repo *userRepository) Update(ctx context.Context, id uint, data users.User) error {
	if data.RawImage != nil {
		UniqueFilename := true
		res, err := repo.cloudinary.Upload.Upload(ctx, data.RawImage, uploader.UploadParams{
			UniqueFilename: &UniqueFilename,
			Folder:         "users",
		})

		if err != nil {
			return err
		}

		data.Image = res.URL
	}

	var mod = new(User)
	mod.FromEntity(data)

	if err := repo.mysqlDB.WithContext(ctx).Where(&User{Id: id}).Updates(mod).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) Delete(ctx context.Context, id uint) error {
	qry := repo.mysqlDB.WithContext(ctx).Delete(&User{}, id)
	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}
