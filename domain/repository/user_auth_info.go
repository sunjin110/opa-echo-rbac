package repository

import "opa-echo-test/domain/entity"

// UserAuthInfoRepository .
type UserAuthInfoRepository interface {
	Get(userID uint64) *entity.UserAuthInfo
}
