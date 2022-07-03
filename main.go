package main

import (
	"os"

	db "github.com/nakoding-community/test-practical-devsecops/database"
	"github.com/nakoding-community/test-practical-devsecops/database/migration"
	"github.com/nakoding-community/test-practical-devsecops/database/seeder"
	"github.com/nakoding-community/test-practical-devsecops/internal/factory"
	"github.com/nakoding-community/test-practical-devsecops/internal/handler/web"
	"github.com/nakoding-community/test-practical-devsecops/internal/handler/ws"
	"github.com/nakoding-community/test-practical-devsecops/internal/middleware"
	"github.com/nakoding-community/test-practical-devsecops/pkg/constant"
	"github.com/nakoding-community/test-practical-devsecops/pkg/cron"
	"github.com/nakoding-community/test-practical-devsecops/pkg/util/env"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func init() {
	ENV := os.Getenv(constant.ENV)
	env := env.NewEnv()
	env.Load(ENV)

	logrus.Info("Choosen environment " + ENV)
}

// @title test-practical-devsecops
// @version 0.0.1
// @description This is a doc for test-practical-devsecops.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:3030
// @BasePath /
func main() {
	var PORT = os.Getenv(constant.PORT)

	// dependency
	db.Init()

	// hook
	migration.Init()
	seeder.Init()

	// lib
	cron.Init()

	e := echo.New()
	middleware.Init(e)

	// factory
	f := factory.Init()

	// handler
	// rest.Init(e, f)
	web.Init(e, f)
	ws.Init(e, f)

	e.Logger.Fatal(e.Start(":" + PORT))
}
