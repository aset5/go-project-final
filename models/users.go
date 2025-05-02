package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:user"`
	Active   bool   `gorm:"default:true"` // 👈 осы жол қосылуы керек
}
