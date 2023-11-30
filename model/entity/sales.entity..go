package entity

type PaymentStatus string

const (
	Unpaid 	PaymentStatus = "unpaid"
	Paid   	PaymentStatus = "paid"
	Void  	PaymentStatus = "void"
)

type PaymentType string
const (
	Cash 	PaymentType = "cash"
	Qris   	PaymentType = "qris"
)

type Sales struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	OrderID     string `gorm:"not null;unique" json:"order_id"`

	Shop 	  	Shops   `gorm:"foreignKey:ShopID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"shop"`
	ShopID      uint64 `db:"shop_id" json:"shop_id"`

	Cashier     User   `gorm:"foreignKey:CashierID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"cashier"`
	CashierID   uint64 `gorm:"not null" json:"cashier_id"`

	Status      PaymentStatus `gorm:"default:unpaid" json:"status"`
	PaymentType PaymentType `gorm:"default:cash" json:"payment_type"`
	TotalBill   int    `gorm:"default:0" json:"total_bill"`
	TotalPaid   int    `gorm:"default:0" json:"total_paid"`
	TotalItem   int    `gorm:"default:0" json:"total_item"`
	Timestamp
}

func (Sales) TableName() string {
	return "sales"
}