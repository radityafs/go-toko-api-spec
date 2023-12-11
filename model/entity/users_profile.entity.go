package entity

type Gender string
const (
	Female 	Gender = "female"
	Male   	Gender = "male"
)


type UserProfile struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserID    uint64 `gorm:"not null" json:"user_id"`
	User      User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
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

type UserProfileResponse struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	UserID    uint64 `gorm:"not null" json:"-"`
	Picture   string `json:"picture"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Gender Gender `json:"gender"`
	Birthdate string `json:"birthdate"`
	Address   string `json:"address"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}

func (UserProfileResponse) TableName() string {
	return "user_profile"
}