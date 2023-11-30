package entity

type ProductsBrand struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Shops 	  	Shops   `gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"shop"`
	ShopID      uint64 `db:"shop_id" json:"shop_id"`
	Name        string `gorm:"not null;unique" json:"name"`
	Description string `gorm:"default:null" json:"description"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`
	IsParent    bool   `gorm:"default:false" json:"is_parent"`
	Timestamp
}

func (ProductsBrand) TableName() string {
	return "products_brand"
}
