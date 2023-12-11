package entity

type SalesDetail struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Sales 	  Sales   `gorm:"foreignKey:SalesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"sales"`
	SalesID   uint64 `gorm:"not null" json:"sales_id"`

	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"product"`
	ProductID uint64 `gorm:"not null" json:"product_id"`

	Shop 	  Shops  `gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"shop"`
	ShopID    uint64 `gorm:"not null" json:"shop_id"`

	Name      string `gorm:"not null" json:"name"`
	Images	string `gorm:"not null" json:"images"`
	Quantity  int    `gorm:"not null" json:"quantity"`
	Category  string `gorm:"not null" json:"category"`
	Price     int    `gorm:"not null" json:"price"`
	Total     int    `gorm:"not null" json:"total"`
	Timestamp
}

func (SalesDetail) TableName() string {
	return "sales_detail"
}