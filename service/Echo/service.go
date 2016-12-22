package Echo

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type Service struct {
	engine *echo.Echo
	port   int
	debug  bool
}

func NewServer(port int, debug bool) (x *Service) {
	x = new(Service)
	x.engine = echo.New()
	x.port = port
	x.debug = debug
	return
}

func (e *Service) Init() {
	// Middleware
	e.engine.Use(middleware.Logger())
	e.engine.Use(middleware.Recover())

	e.engine.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Working....")
	})
	e.engine.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "pong",
		})
	})
}

func (e *Service) Run() {
	go func() {
		if e.debug {
			fmt.Println("Debug Mode...")
			e.engine.Logger.Fatal(e.engine.Start(fmt.Sprintf(":%d", e.port)))
		} else {
			e.engine.Start(fmt.Sprintf(":%d", e.port))
		}
	}()
}
