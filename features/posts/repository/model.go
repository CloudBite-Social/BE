package repository

import (
	"reflect"
	"time"

	"sosmed/features/comments"
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

func (mod *Post) ToEntity() *posts.Post {
	var ent = new(posts.Post)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Caption != "" {
		ent.Caption = mod.Caption
	}

	if !reflect.ValueOf(mod.User).IsZero() {
		if mod.User.Id != 0 {
			ent.User.Id = mod.User.Id
		}

		if mod.User.Name != "" {
			ent.User.Name = mod.User.Name
		}

		if mod.User.Image != "" {
			ent.User.Image = mod.User.Image
		}

		if !mod.User.CreatedAt.IsZero() {
			ent.User.CreatedAt = mod.User.CreatedAt
		}

		if !mod.User.UpdatedAt.IsZero() {
			ent.User.UpdatedAt = mod.User.UpdatedAt
		}

		if !mod.User.DeletedAt.Time.IsZero() {
			ent.User.DeletedAt = mod.User.DeletedAt.Time
		}
	}

	if len(mod.Comment) != 0 {
		for _, comment := range mod.Comment {
			if !reflect.ValueOf(comment).IsZero() {
				tempComment := new(comments.Comment)

				if comment.Id != 0 {
					tempComment.Id = comment.Id
				}

				if comment.Text != "" {
					tempComment.Text = comment.Text
				}

				if !reflect.ValueOf(comment.User).IsZero() {
					if comment.User.Id != 0 {
						tempComment.User.Id = comment.User.Id
					}

					if comment.User.Name != "" {
						tempComment.User.Name = comment.User.Name
					}

					if comment.User.Image != "" {
						tempComment.User.Image = comment.User.Image
					}

					if !comment.User.CreatedAt.IsZero() {
						tempComment.User.CreatedAt = comment.User.CreatedAt
					}

					if !comment.User.UpdatedAt.IsZero() {
						tempComment.User.UpdatedAt = comment.User.UpdatedAt
					}

					if !comment.User.DeletedAt.Time.IsZero() {
						tempComment.User.DeletedAt = comment.User.DeletedAt.Time
					}
				}

				if !comment.CreatedAt.IsZero() {
					tempComment.CreatedAt = comment.CreatedAt
				}

				if !comment.UpdatedAt.IsZero() {
					tempComment.UpdatedAt = comment.UpdatedAt
				}

				if !comment.DeletedAt.Time.IsZero() {
					tempComment.DeletedAt = comment.DeletedAt.Time
				}

				ent.Comments = append(ent.Comments, *tempComment)
			}
		}
	}

	if len(mod.Attachment) != 0 {
		for _, file := range mod.Attachment {
			ent.Attachment = append(ent.Attachment, file.ToEntity())
		}
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	if !mod.UpdatedAt.IsZero() {
		ent.UpdatedAt = mod.UpdatedAt
	}

	if !mod.DeletedAt.Time.IsZero() {
		ent.DeletedAt = mod.DeletedAt.Time
	}

	return ent
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

func (mod *File) ToEntity() posts.File {
	var ent = new(posts.File)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.URL != "" {
		ent.URL = mod.URL
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	if !mod.UpdatedAt.IsZero() {
		ent.UpdatedAt = mod.UpdatedAt
	}

	if !mod.DeletedAt.Time.IsZero() {
		ent.DeletedAt = mod.DeletedAt.Time
	}

	return *ent
}
