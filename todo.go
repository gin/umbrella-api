package main

import (
	"./handlers"
	"database/sql"

	"github.com/labstack/echo"
	//	"github.com/labstack/echo/engine/standard"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	//	e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
	//	e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
	//	e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

	// Start as a web server
	//e.Run(standard.New(":8000"))

	e.File("/", "index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	);
	`

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
