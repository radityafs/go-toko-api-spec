package response


type Pagination struct {
	TotalData int64 `gorm:"column:count" json:"total_data"`
	TotalPage int64 `json:"total_page"`
	CurrentPage int64 `json:"current_page"`
	PerPage int64 `json:"per_page"`
}