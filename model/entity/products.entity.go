package entity

type Product struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	ProductsCategory 	ProductsCategory   `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"category"`
	CategoryID  uint64 `gorm:"not null" json:"category_id"`

	ProductsBrand 	  	ProductsBrand   `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"brand"`
	BrandID     uint64 `gorm:"not null" json:"brand_id"`

	Shops 	  	Shops   `gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	ShopID      uint64 `gorm:"not null" json:"shop_id"`
	
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"default:null" json:"-"`
	SKU         string `gorm:"default:null" json:"sku"`
	Quantity    int    `gorm:"default:0" json:"quantity"`
	PriceBuy    int    `gorm:"default:0" json:"-"`
	PriceSell   int    `gorm:"default:0" json:"price"`
	Images      string `gorm:"default:null" json:"images"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`
	Timestamp
}

func (Product) TableName() string {
	return "products"
}