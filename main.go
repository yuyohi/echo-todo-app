package app

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := sql.Open("sqlite3", "todo_app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewTaskRepository(db)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1234"))
}
