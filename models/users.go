package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:user"`
	IsBanned bool   `gorm:"default:false"` // нақты атауы
	Active   bool   `gorm:"default:true"`
}
