package handler

import "sosmed/features/users"

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
