package models

type Book struct {
	Id     int64  `gorm:"primaryKey" json:"id"`
	Title  string `gorm:"type:varchar(300)" json:"title"`
	Author string `gorm:"type:varchar(300)" json:"author"`
	Desc   string `gorm:"type:text" json:"desc"`
}
