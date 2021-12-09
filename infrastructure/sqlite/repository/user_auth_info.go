package repository

import (
	"opa-echo-test/domain/entity"

	"gorm.io/gorm"
)

type UserAuthInfoRepositoryImpl struct {
	db *gorm.DB
}

func (impl UserAuthInfoRepositoryImpl) Get(uesrID uint64) *entity.UserAuthInfo {

	// db := sqlite.GetDB()

	impl.db.Get("")

	return nil
}
