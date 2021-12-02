package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"notes-api/dtos"
	"notes-api/models"
	"strings"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

func (r *NoteRepository) Save(note *models.Note) RepositoryResult {
	err := r.db.Save(note).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: note}
}

func (r *NoteRepository) FindAll(pagination *dtos.Pagination) RepositoryResult {
	var notes models.Notes

	var totalRows int64 = 0
	offset := pagination.Page * pagination.Limit
	find := r.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	searchs := pagination.Searchs
	if searchs != nil {
		for _, value := range searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			case "contains":
				whereQuery := fmt.Sprintf("%s LIKE ?", column)
				find = find.Where(whereQuery, "%"+query+"%")
				break
			case "in":
				whereQuery := fmt.Sprintf("%s IN (?)", column)
				queryArray := strings.Split(query, ",")
				find = find.Where(whereQuery, queryArray)
				break
			default:
				whereQuery := fmt.Sprintf("%s = ?", column)
				find = find.Where(whereQuery, query)
				break
			}
		}
	}
	findErr := find.Find(&notes).Error

	if findErr != nil {
		return RepositoryResult{Error: findErr}
	}
	pagination.Rows = notes
	errCount := r.db.Model(&models.Note{}).Count(&totalRows).Error
	if errCount != nil {
		return RepositoryResult{Error: errCount}
	}
	pagination.TotalRows = totalRows

	return RepositoryResult{Result: pagination}
}

func (r *NoteRepository) FindOneById(id string) RepositoryResult {
	var note models.Note

	err := r.db.Where(&models.Note{ID: id}).Take(&note).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &note}
}

func (r *NoteRepository) DeleteOneById(id string) RepositoryResult {
	err := r.db.Delete(&models.Note{ID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
