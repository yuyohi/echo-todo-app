package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"

	"github.com/yuyohi/echo-todo-app/app"
)

func main() {
	db, err := sql.Open("sqlite3", "./db/todo_app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := app.NewTaskRepository(db)
	s := app.NewService(r)
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
			return nil
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/tasks", func(c echo.Context) error {
		req := new(app.TaskRequest)
		if err := c.Bind(req); err != nil {
			return err
		}
		id, err := s.CreateTask(req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, id)
	})

	e.GET("/tasks", func(c echo.Context) error {
		tasks, err := s.GetTasks()
		fmt.Println("tasks: ", tasks)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, tasks)
	})

	e.POST("/tasks/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		req := new(app.TaskRequest)
		if err := c.Bind(req); err != nil {
			return err
		}
		err = s.UpdateTask(req, id)
		if err != nil {
			return err
		}
		return c.NoContent(http.StatusNoContent)
	})

	e.DELETE("/tasks/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		err = s.DeleteTask(id)
		if err != nil {
			return err
		}
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Fatal(e.Start(":1234"))
}
