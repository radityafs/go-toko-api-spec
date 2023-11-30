package entity

type ProductsCategory struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Shops 	  	Shops   `gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	ShopID      uint64 `gorm:"not null" json:"-"`
	Name        string `gorm:"not null" json:"name"`
	Code        string `gorm:"size:255;not null" json:"-"`
	Description string `gorm:"default:null" json:"-"`
	Images      string `gorm:"default:null" json:"-"`
	IsActive    bool   `gorm:"default:true" json:"-"`
	IsParent    bool   `gorm:"default:false" json:"-"`
}

func (ProductsCategory) TableName() string {
	return "products_category"
}