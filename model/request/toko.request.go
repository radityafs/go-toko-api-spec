package request

type CreateToko struct {
	Nama    string `json:"name" validate:"required"`
	Alamat  string `json:"address"`
	Telepon string `json:"phone"`
}

type UpdateToko struct {
	Nama    string `json:"name"`
	Alamat  string `json:"address"`
	Telepon string `json:"phone"`
}
