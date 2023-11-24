package response

import "go-toko/model/entity"

type User struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []entity.User `json:"data"`
}

type DetailUser struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    entity.User `json:"data"`
}
