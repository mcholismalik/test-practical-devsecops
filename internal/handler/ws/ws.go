package ws

import (
	"github.com/nakoding-community/test-practical-devsecops/internal/factory"
	"github.com/nakoding-community/test-practical-devsecops/internal/handler/ws/course"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, f factory.Factory) {
	prefix := "ws"
	course.NewHandler(f).Route(e.Group(prefix + "/course"))
}
