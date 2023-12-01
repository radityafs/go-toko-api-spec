package entity

type Gender string
const (
	Female 	Gender = "female"
	Male   	Gender = "male"
)


type UserProfile struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	User      User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	UserID    uint64 `gorm:"not null" json:"user_id"`
	Picture   string `gorm:"default:null" json:"picture"`
	Phone     string `gorm:"default:null" json:"phone"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Nickname  string `gorm:"default:null" json:"nickname"`
	Gender Gender `gorm:"not null" json:"gender"`
	Birthdate string `gorm:"default:null" json:"birthdate"`
	Address   string `gorm:"default:null" json:"address"`
	Timestamp
}

func (UserProfile) TableName() string {
	return "users_profile"
}