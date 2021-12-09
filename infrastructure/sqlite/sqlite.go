package sqlite

import (
	"opa-echo-test/infrastructure/sqlite/repository"
	"opa-echo-test/internal/chk"
	"os"

	// _ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Setup .
func Setup(dbDir string, dbPath string) {

	// refresh db
	os.Remove(dbPath)

	// make dir
	os.Mkdir(dbDir, 0777)

	gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	chk.SE(err)

	db = gormDB

	// migration
	chk.SE(db.AutoMigrate(&repository.User{}))
	chk.SE(db.AutoMigrate(&repository.UserRole{}))
}

func GetDB() *gorm.DB {
	return db
}
