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

// UserAllowedResource ユーザーのアクセスが許可されているresource
type UserAllowedResource struct {
	UserAllowedResourceID uint64 `gorm:"primaryKey; column:user_allowed_resource_id; AUTO_INCREMENT;"`
	UserID                uint64 `gorm:"column:user_id"`
	Resource              string `gorm:"column:resource"`
}

// Get TODO 後でjoinのqueyrにする
func (impl UserAuthInfoRepositoryImpl) Get(userID uint64) *entity.UserAuthInfo {

	user := &User{}
	impl.db.Take(user, "user_id = ?", userID)

	var userRoleList []UserRole
	impl.db.Where("user_id = ?", userID).Find(&userRoleList)

	var userAllowedResourceList []UserAllowedResource
	impl.db.Where("user_id = ?", userID).Find(&userAllowedResourceList)

	var roleList []string
	for _, userRole := range userRoleList {
		roleList = append(roleList, userRole.Role)
	}

	var resourceList []string
	for _, userResource := range userAllowedResourceList {
		resourceList = append(resourceList, userResource.Resource)
	}

	return &entity.UserAuthInfo{
		UserID:       user.UserID,
		Name:         user.Name,
		RoleList:     roleList,
		ResourceList: resourceList,
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
			Role:   role,
		}
		userRoleList = append(userRoleList, userRole)
	}
	if len(userRoleList) > 0 {
		impl.db.Create(&userRoleList)
	}

	var userAllowedResourceList []UserAllowedResource
	for _, resource := range userAuthInfo.ResourceList {
		userAllowedResource := UserAllowedResource{
			UserID:   userAuthInfo.UserID,
			Resource: resource,
		}
		userAllowedResourceList = append(userAllowedResourceList, userAllowedResource)
	}

	if len(userAllowedResourceList) > 0 {
		impl.db.Create(&userAllowedResourceList)
	}

	return true
}
