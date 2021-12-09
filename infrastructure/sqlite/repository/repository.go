package repository

// User ユーザー情報
type User struct {
	UserID uint64 `gorm:"primaryKey column:user_id"`
	Name   string `gorm:"column:name"`
}

// UserRole ユーザーアカウント情報
type UserRole struct {
	UserRoleID uint64 `gorm:"primaryKey clumn:user_role_id"`
	UserID     uint64 `gorm:"column:user_id"`
	Role       string `gorm:"column:role"`
}
