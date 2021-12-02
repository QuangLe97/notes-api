package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	m "notes-api/models"
	rb "notes-api/rabit_mq"
	r "notes-api/repositories"
	s "notes-api/services"
	"notes-api/utils"
)

type H map[string]interface{}

// GetAllNotes Get all notes endpoint
func GetAllNotes(noteRepository *r.NoteRepository, ch *amqp.Channel, q amqp.Queue) echo.HandlerFunc {
	return func(c echo.Context) error {
		pagination := utils.GetPaginationInfo(c)
		response := s.FindAllNotes(*noteRepository, pagination)
		code := http.StatusOK
		if !response.Success {
			code = http.StatusBadRequest
		}
		rb.PublicMessage(ch, q, fmt.Sprintf("GET | GetAllNotes | status: %d", code))
		return c.JSON(code, response)
	}
}

// GetDetailNotes Get detail notes endpoint
func GetDetailNotes(noteRepository *r.NoteRepository, ch *amqp.Channel, q amqp.Queue) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		res := s.FindOneNoteById(id, *noteRepository)
		rb.PublicMessage(ch, q, fmt.Sprintf("GET | GetDetailNotes | status: %d", 200))
		return c.JSON(http.StatusOK, res)
	}
}

// PutNote endpoint
func PutNote(noteRepository *r.NoteRepository, ch *amqp.Channel, q amqp.Queue) echo.HandlerFunc {
	return func(c echo.Context) error {
		var note m.Note
		id := c.Param("id")
		_ = c.Bind(&note)
		log.Println(note)
		code := http.StatusOK
		response := s.UpdateNoteById(id, &note, *noteRepository)
		if !response.Success {
			code = http.StatusBadRequest
		}
		rb.PublicMessage(ch, q, fmt.Sprintf("PUT | PutNote | status: %d", code))
		return c.JSON(code, response)
	}
}

// CreateNote endpoint
func CreateNote(noteRepository *r.NoteRepository, ch *amqp.Channel, q amqp.Queue) echo.HandlerFunc {
	return func(c echo.Context) error {
		var note m.Note
		_ = c.Bind(&note)
		log.Println(note)
		code := http.StatusOK
		response := s.CreateNote(&note, *noteRepository)
		if !response.Success {
			code = http.StatusBadRequest
		}
		rb.PublicMessage(ch, q, fmt.Sprintf("POST | CreateNote | status: %d", code))
		return c.JSON(code, response)
	}
}

// DeleteNote endpoint
func DeleteNote(noteRepository *r.NoteRepository, ch *amqp.Channel, q amqp.Queue) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		code := http.StatusOK
		response := s.DeleteOneNoteById(id, *noteRepository)
		if !response.Success {
			code = http.StatusBadRequest
		}
		rb.PublicMessage(ch, q, fmt.Sprintf("DELETE | GetAllNotes | status: %d", code))
		return c.JSON(code, response)
	}
}
