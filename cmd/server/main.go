package main

import (
	"github.com/labstack/echo/middleware"

	"notes-api/database"
	"notes-api/models"
	rb "notes-api/rabbit_mq"
	"notes-api/repositories"
	"os"
)
import (
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"

	h "notes-api/handler"
)

func main() {

	db := database.ConnectToDB(os.Getenv("DB_PATH"))

	// migration
	errMrg := db.AutoMigrate(&models.Note{})
	if errMrg != nil {
		panic("migrate error")
	}
	noteRepository := repositories.NewNoteRepository(db)
	e := echo.New()
	e.Use(middleware.Logger())
	// RabbitMQ
	channel, queue := rb.ConnectChannel()
	defer channel.Close()

	e.GET("/notes", h.GetAllNotes(noteRepository, channel, queue))
	e.GET("/notes/:id", h.GetDetailNotes(noteRepository, channel, queue))
	e.PUT("/notes/:id", h.PutNote(noteRepository, channel, queue))
	e.POST("/notes", h.CreateNote(noteRepository, channel, queue))
	e.DELETE("/notes/:id", h.DeleteNote(noteRepository, channel, queue))

	// Start as a web server
	e.Logger.Fatal(e.Start(":8787"))
}
