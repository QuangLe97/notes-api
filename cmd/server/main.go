package main

import (
	"database/sql"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"

	h "notes-api/handler"
)

func main() {
	db := initDB("storage.db")

	e := echo.New()

	e.GET("/notes", h.GetListNotes(db))
	e.GET("/notes/:id", h.GetDetailNotes(db))
	e.PUT("/notes/:id", h.PutNote(db))
	e.POST("/notes", h.CreateNote(db))
	e.DELETE("/notes/:id", h.DeleteNote(db))

	// Start as a web server
	e.Logger.Fatal(e.Start(":8787"))
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
