package response

type UserProfile struct {
	Picture   string `gorm:"column:picture" json:"picture"`
	Phone     string `gorm:"column:phone" json:"phone"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Nickname  string `gorm:"column:nickname" json:"nickname"`
	Gender    string `gorm:"column:gender" json:"gender"`
	Birthdate string `gorm:"column:birthdate" json:"birthdate"`
	Address   string `gorm:"column:address" json:"address"`
}