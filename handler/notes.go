package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	m "notes-api/models"
)

type H map[string]interface{}

// GetListNotes Get all notes endpoint
func GetListNotes(db *sql.DB) echo.HandlerFunc {

	return func(c echo.Context) error {
		queryString := c.QueryParam("query_str")
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page < 0 {
			return c.JSON(http.StatusUnprocessableEntity, m.BaseResponseErrors{
				Code: 1,
				Errors: map[string]string{
					"message": "page should be type int and greater than or equal to 0",
				},
			})
		}
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		if err != nil || limit < 0 || limit > 100 {
			limit = 50
		}

		res := m.GetNotes(db, queryString, page, limit)
		return c.JSON(http.StatusOK, m.BaseResponseSuccess{
			Code:   0,
			Result: res,
		})
	}
}

// GetDetailNotes Get detail notes endpoint
func GetDetailNotes(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		res, err := m.GetNote(db, id)
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusOK, m.BaseResponseSuccess{
			Code:   0,
			Result: res,
		})
	}
}

// PutNote endpoint
func PutNote(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new note
		var note m.Note
		id, _ := strconv.Atoi(c.Param("id"))
		// Map imcoming JSON body to the new Note
		_ = c.Bind(&note)

		log.Println(note)

		// Add a note using our new model
		rowsAffected, err := m.PutNote(db, id, note.Name, note.Content)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusOK, m.BaseResponseSuccess{
				Code: 0,
				Result: map[string]interface{}{
					"updated": rowsAffected},
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

// CreateNote endpoint
func CreateNote(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new note
		var note m.Note
		// Map incoming JSON body to the new Note
		_ = c.Bind(&note)

		log.Println(note)

		// Add a note using our new model
		id, err := m.CreateNote(db, note.Name, note.Content)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusOK, m.BaseResponseSuccess{
				Code: 0,
				Result: map[string]interface{}{
					"created": id},
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

// DeleteNote endpoint
func DeleteNote(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a note
		numDeleted, err := m.DeleteNote(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, m.BaseResponseSuccess{
				Code: 0,
				Result: map[string]interface{}{
					"deleted": numDeleted},
			})
			// Handle errors
		} else {
			return err
		}
	}
}
