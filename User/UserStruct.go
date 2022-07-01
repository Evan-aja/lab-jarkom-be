package User

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"notNull" json:"name"`
	NIM      string `gorm:"uniquiIndex;notNull;size:30" json:"nim"`
	Password string `gorm:"notNull" json:"password"`
}
type UserLog struct {
	Name     string `gorm:"notNull" json:"name"`
	NIM      string `gorm:"uniquiIndex;notNull;size:30" json:"nim"`
	Password string `gorm:"notNull" json:"password"`
}
