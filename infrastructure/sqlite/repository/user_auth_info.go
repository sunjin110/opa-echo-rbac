package repository

import (
	"opa-echo-test/domain/entity"

	"gorm.io/gorm"
)

type UserAuthInfoRepositoryImpl struct {
	db *gorm.DB
}

// NewUserAuthInfoRepository .
func NewUserAuthInfoRepository(db *gorm.DB) *UserAuthInfoRepositoryImpl {
	return &UserAuthInfoRepositoryImpl{
		db: db,
	}
}

// User ユーザー情報
type User struct {
	UserID uint64 `gorm:"primaryKey; column:user_id; AUTO_INCREMENT;"`
	Name   string `gorm:"column:name"`
}

// UserRole ユーザーアカウント情報
type UserRole struct {
	UserRoleID uint64 `gorm:"primaryKey;clumn:user_role_id; AUTO_INCREMENT;"`
	UserID     uint64 `gorm:"column:user_id"`
	Role       string `gorm:"column:role"`
}

// Get TODO 後でjoinのqueyrにする
func (impl UserAuthInfoRepositoryImpl) Get(userID uint64) *entity.UserAuthInfo {

	user := &User{}
	impl.db.Take(user, "user_id = ?", userID)

	var userRoleList []UserRole
	impl.db.Where("user_id = ?", userID).Find(&userRoleList)

	var entityUserRoleList []*entity.UserRole
	for _, userRole := range userRoleList {

		entityUserRole := &entity.UserRole{
			UserRoleID: userRole.UserRoleID,
			Role:       userRole.Role,
		}
		entityUserRoleList = append(entityUserRoleList, entityUserRole)
	}

	return &entity.UserAuthInfo{
		UserID:   user.UserID,
		Name:     user.Name,
		RoleList: entityUserRoleList,
	}
}

func (impl UserAuthInfoRepositoryImpl) Insert(userAuthInfo *entity.UserAuthInfo) bool {

	if userAuthInfo == nil {
		return false
	}

	user := &User{
		UserID: userAuthInfo.UserID,
		Name:   userAuthInfo.Name,
	}

	impl.db.Create(user)

	var userRoleList []UserRole
	for _, role := range userAuthInfo.RoleList {

		userRole := UserRole{
			UserID: userAuthInfo.UserID,
			Role:   role.Role,
		}
		userRoleList = append(userRoleList, userRole)
	}

	impl.db.Create(&userRoleList)

	return true
}
