package entity

type Gender string
const (
	Female 	Gender = "female"
	Male   	Gender = "male"
)


// | Column     | Type         | Modifiers                    |
// | ---------- | ------------ | ---------------------------- |
// | id         | integer      | not null default primary key |
// | user_id    | integer      | foreign key: `users.id`      |
// | picture    | varchar(255) | nullable                     |
// | phone      | varchar(255) | nullable                     |
// | first_name | varchar(255) |                              |
// | last_name  | varchar(255) |                              |
// | nickname   | varchar(255) | nullable                     |
// | gender     | enum         | enum: `male`, `female`       |
// | birthdate  | date         | nullable                     |
// | address    | varchar(255) | nullable                     |
// | created_at | timestamp    | nullable                     |
// | updated_at | timestamp    | nullable                     |

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