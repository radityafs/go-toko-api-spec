package request

import "github.com/google/uuid"

type CreateUser struct {
	Email    string    `json:"email" validate:"required,email"`
	Nama     string    `json:"name" validate:"required"`
	TokoID   uuid.UUID `json:"shop_id" validate:"required,uuid"`
	Role     string    `json:"role" validate:"required"`
	Password string    `json:"password" validate:"required,min=8"`
}
