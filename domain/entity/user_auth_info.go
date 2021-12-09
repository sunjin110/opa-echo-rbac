package entity

// UserAuthInfo ユーザーの認可情報のやつ
type UserAuthInfo struct {
	UserID       uint64
	Name         string
	RoleList     []string // ユーザーが所持しているRole
	ResourceList []string // ユーザーがアクセスを許可されているResource
}
