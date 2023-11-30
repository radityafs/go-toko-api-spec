package entity

import "time"

type User struct {
	ID               uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Email            string `gorm:"not null;unique" json:"email"`
	EmailVerifiedAt *time.Time `gorm:"default:null" json:"-"`
	Password         string `gorm:"not null" json:"-"`
	LastSeen         *time.Time `gorm:"default:null" json:"last_seen"`
	IsSubscribe      bool   `gorm:"default:false" json:"is_subscribe"`
	RememberToken    string `gorm:"default:null" json:"-"`

	Role             UserRole `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"role"`
	RoleID           uint64 `gorm:"not null" json:"-"`

	Shop			 Shops `gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"shop"`
	ShopID           uint64 `gorm:"default:null" json:"-"`

	GoogleID         string `gorm:"default:null" json:"-"`
	Timestamp
}

func (User) TableName() string {
	return "users"
}
