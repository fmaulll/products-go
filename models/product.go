package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Product     string `gorm:"type:varchar(300)" json:"product"`
	Description string `gorm:"type:text" json:"description"`
}
