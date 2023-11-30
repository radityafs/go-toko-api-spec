package entity

type UserRole struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name 	string `gorm:"not null;unique" json:"name"`
	Timestamp
}

func (UserRole) TableName() string {
	return "roles"
}