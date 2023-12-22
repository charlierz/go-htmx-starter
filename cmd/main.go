package main

import (
	"fmt"
	"strings"

	"github.com/charlierz/go-htmx-starter/internal/db"
	"github.com/charlierz/go-htmx-starter/internal/features/home"
	"github.com/charlierz/go-htmx-starter/internal/features/item"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

const DB_FILE string = "db.sqlite"
const PORT int = 8080

func main() {
	d := db.InitDb(DB_FILE)
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${error}\n",
		Skipper: func(c echo.Context) bool {
			return c.Request().RequestURI == "/favicon.ico" ||
				strings.HasPrefix(c.Request().RequestURI, "/static")
		},
	}))

	e.Static("/static", "web/static")

	itemService := item.NewService(d)
	homeService := home.NewService(d)

	item.Register(e, item.NewHandler(itemService))
	home.Register(e, home.NewHandler(homeService, itemService))

	addr := fmt.Sprintf(":%d", PORT)
	e.Logger.Fatal(e.Start(addr))
}
