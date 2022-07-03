package seeder

import (
	"context"

	"github.com/nakoding-community/test-practical-devsecops/database"
	"github.com/nakoding-community/test-practical-devsecops/internal/abstraction"
	"github.com/nakoding-community/test-practical-devsecops/pkg/constant"
	"github.com/nakoding-community/test-practical-devsecops/pkg/util/env"

	"github.com/google/uuid"
)

var Context = &abstraction.Context{
	Context: context.Background(),
	Auth: &abstraction.AuthContext{
		ID:    uuid.New(),
		Name:  "system",
		Email: "system@system.sys",
	},
}

func Init() {
	if !env.NewEnv().GetBool(constant.IS_RUN_SEEDER) {
		return
	}

	conn, err := database.GetConnection(constant.DB_GOBOIL_CLEAN)
	if err != nil {
		panic(err)
	}

	UserTableSeeder(conn)
}
