package domain

type Customer struct {
	Customer_id  int    `gorm:"column:customer_id;primaryKey;autoIncrement"`
	Name         string `gorm:"column:name"`
	Address      string `gorm:"column:address"`
	Username     string `gorm:"column:username"`
	Password     string `gorm:"column:password"`
	Phone_number string `gorm:"column:phone_number"`
	Email        string `gorm:"column:email"`
}

func (Customer) TableName() string {
	return "customer"
}
