package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/sasaki-q/test/config"
)

func main() {
	e := echo.New()
	db := config.New()
	config.Migrate(db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
