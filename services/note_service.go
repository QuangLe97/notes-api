package services

import (
	"github.com/google/uuid"
	"log"
	"notes-api/dtos"
	"notes-api/models"
	"notes-api/repositories"
)

func CreateNote(note *models.Note, repository repositories.NoteRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	note.ID = uuidResult.String()

	operationResult := repository.Save(note)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Note)

	return dtos.Response{Success: true, Data: data}
}

func FindAllNotes(repository repositories.NoteRepository, pagination *dtos.Pagination) dtos.Response {
	operationResult := repository.FindAll(pagination)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}
	var data = operationResult.Result.(*dtos.Pagination)
	return dtos.Response{Success: true, Data: data}
}

func FindOneNoteById(id string, repository repositories.NoteRepository) dtos.Response {
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Note)

	return dtos.Response{Success: true, Data: data}
}

func UpdateNoteById(id string, note *models.Note, repository repositories.NoteRepository) dtos.Response {
	existingNoteResponse := FindOneNoteById(id, repository)

	if !existingNoteResponse.Success {
		return existingNoteResponse
	}

	existingNote := existingNoteResponse.Data.(*models.Note)

	existingNote.Name = note.Name
	existingNote.Content = note.Content

	operationResult := repository.Save(existingNote)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func DeleteOneNoteById(id string, repository repositories.NoteRepository) dtos.Response {
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}
