package handler

import (
	"io"
	"sosmed/features/users"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *RegisterUserRequest) ToEntity() *users.User {
	var ent = new(users.User)

	if req.Name != "" {
		ent.Name = req.Name
	}

	if req.Email != "" {
		ent.Email = req.Email
	}

	if req.Password != "" {
		ent.Password = req.Password
	}

	return ent
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *LoginUserRequest) ToEntity() *users.User {
	var ent = new(users.User)

	if req.Email != "" {
		ent.Email = req.Email
	}

	if req.Password != "" {
		ent.Password = req.Password
	}

	return ent
}

type UpdateUserRequest struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Image    io.Reader
}

func (req *UpdateUserRequest) ToEntity() *users.User {
	var ent = new(users.User)

	if req.Name != "" {
		ent.Name = req.Name
	}

	if req.Email != "" {
		ent.Email = req.Email
	}

	if req.Password != "" {
		ent.Password = req.Password
	}

	if req.Image != nil {
		ent.RawImage = req.Image
	}

	return ent
}
