package agro

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string `gorm:"unique;not null;size:50"`
	Email    string `gorm:"unique;not null;size:100"`
	Password string `gorm:"not null;size:255"`
}
