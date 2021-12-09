package entity

// UserAuthInfo ユーザーの認可情報のやつ
type UserAuthInfo struct {
	UserID   uint64
	Name     string
	RoleList []*UserRole // ユーザーが所持しているRole
}

// UserRole ユーザーが所持できるRole情報
type UserRole struct {
	UserRoleID uint64 // id
	Role       string // role名
}
