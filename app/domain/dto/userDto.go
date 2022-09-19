package dto

type User struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint64 `gorm:"not null;" json:"userId"`
	Description string `gorm:"type:text ;" json:"shortContent"`
}
