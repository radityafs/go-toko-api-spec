package response

type Analytics struct {
	Status  int         `json:"-"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    []AnalyticsData `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type ProductsResponse struct {
	Photo string `gorm:"column:images" json:"photo"`
}


type AnalyticsData struct {
	Name string `gorm:"column:name" json:"name"`
	Images string `gorm:"column:images" json:"images"`
	TotalSales int64 `gorm:"column:total_sales" json:"total_sales"`
	TotalRevenue int64 `gorm:"column:total_revenue" json:"total_revenue"`
}

func (ProductsResponse) TableName() string {
	return "products"
}