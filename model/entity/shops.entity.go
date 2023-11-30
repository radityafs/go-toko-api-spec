package entity

type Shops struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserID      uint64 `gorm:"not null" json:"-"`
	Product    []Product `gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"default:null" json:"description"`
	Province    string `gorm:"default:null" json:"province"`
	Regency     string `gorm:"default:null" json:"regency"`
	District    string `gorm:"default:null" json:"district"`
	Village     string `gorm:"default:null" json:"village"`
	Address     string `gorm:"default:null" json:"address"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`
	Timestamp
}

func (Shops) TableName() string {
	return "shops"
}