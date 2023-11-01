package models

type User struct {
	BaseModel
	FirstName string `json:"first_name" gorm:"column:first_name; type:varchar(255); not null"`
	LastName  string `json:"last_name" gorm:"column:last_name; type:varchar(255); not null"`
	Phone     string `json:"phone" gorm:"column:phone; type:varchar(255); not null"`
	Address   string `json:"address" gorm:"column:address; type:varchar(255); not null"`
}

func (u User) TableName() string {
	return "users"
}

func NewUser() *User {
	return &User{}
}
