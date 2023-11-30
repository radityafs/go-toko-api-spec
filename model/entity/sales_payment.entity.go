package entity


type SalesPayment struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Sales 	  	Sales   `gorm:"foreignKey:SalesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"sales"`
	SalesID     uint64 `gorm:"not null" json:"sales_id"`
	
	PaymentRef  string `gorm:"not null;unique" json:"payment_ref"`
	PaymentType PaymentType `gorm:"default:cash" json:"payment_type"`
	Amount      int    `gorm:"not null" json:"amount"`
	Status      PaymentStatus `gorm:"default:unpaid" json:"status"`
	Timestamp
}

func (SalesPayment) TableName() string {
	return "sales_payment"
}