package repository_test

import (
	"log"
	"opa-echo-test/domain/entity"
	"opa-echo-test/infrastructure/sqlite"
	"opa-echo-test/infrastructure/sqlite/repository"
	"opa-echo-test/internal/jsonutil"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s -run ^Test$ opa-echo-test/infrastructure/sqlite/repository

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("repository", func() {

		sqlite.Setup("./testdata", "./testdata/test.db")

		repo := repository.NewUserAuthInfoRepository(sqlite.GetDB())

		g.It("user_auth_info:create", func() {

			userAuthInfo := &entity.UserAuthInfo{
				UserID:       1,
				Name:         "test-name",
				RoleList:     []string{"test-role-1", "test-role-2"},
				ResourceList: []string{"test-resource-1", "test-resource-2"},
			}

			repo.Insert(userAuthInfo)
		})

		g.It("user_auth_info:find", func() {
			userAuthInfo := repo.Get(1)
			log.Println("userAuthInfo is ", jsonutil.Marshal(userAuthInfo))
		})

	})

}
