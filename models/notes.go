package models

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//base response model
type BaseResponseSuccess struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}
type BaseResponseErrors struct {
	Code   int         `json:"code"`
	Errors interface{} `json:"errors"`
}
type PaginationData struct {
	TaskCollection
	Totals   int `json:"totals"`
	PageSize int `json:"page_size"`
	Page     int `json:"page"`
	Size     int `json:"size"`
}

// Note is a struct containing Note data
type Note struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Content     string    `json:"content"`
	TimeCreated time.Time `json:"time_created"`
	TimeUpdated time.Time `json:"time_updated"`
}

// TaskCollection is collection of Notes
type TaskCollection struct {
	Notes []Note `json:"items"`
}

func GetNotes(db *sql.DB, queryStr string, page int, limit int) PaginationData {
	sqlRaw := "SELECT * FROM notes where name like ? LIMIT ? OFFSET ?"
	sqlRawCount := "SELECT COUNT(*) FROM notes where name like ?"
	rows, err := db.Query(sqlRaw, "%"+queryStr+"%", limit, page*limit)
	if err != nil {
		panic(err)
	}
	// make sure to clean up when the program exits
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	result := TaskCollection{}

	for rows.Next() {
		note := Note{}
		err2 := rows.Scan(&note.ID, &note.Name, &note.Content, &note.TimeCreated, &note.TimeUpdated)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Notes = append(result.Notes, note)
	}
	var count int
	err2 := db.QueryRow(sqlRawCount, "%"+queryStr+"%").Scan(&count)
	if err2 != nil {
		panic(err2)
	}
	return PaginationData{
		result,
		count,
		limit,
		page,
		len(result.Notes),
	}
}

func GetNote(db *sql.DB, id int) (*Note, error) {
	sqlRaw := "SELECT * FROM notes WHERE id=?"
	rows, err := db.Query(sqlRaw, id)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	note := &Note{}
	for rows.Next() {

		err2 := rows.Scan(&note.ID, &note.Name, &note.Content, &note.TimeCreated, &note.TimeUpdated)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}

		return note, nil
	}

	return nil, errors.New("Note not found")
}

func PutNote(db *sql.DB, id int, name string, content string) (int64, error) {
	sqlRaw := "UPDATE notes SET name = ?,content = ?,time_updated=? WHERE id=?"
	//sqlRaw := "INSERT INTO notes(name,content,time_updated) VALUES(?,?,?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sqlRaw)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name, content, time.Now(), id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func CreateNote(db *sql.DB, name string, content string) (int64, error) {
	sqlRaw := "INSERT INTO notes(name,content,time_created,time_updated) VALUES(?,?,?,?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sqlRaw)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name, content, time.Now(), time.Now())
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func DeleteNote(db *sql.DB, id int) (int64, error) {
	sqlRaw := "DELETE FROM notes WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sqlRaw)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
