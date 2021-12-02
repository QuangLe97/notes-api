package repositories

import (
	"gorm.io/gorm"
	"notes-api/models"
)

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{db: db}
}

func (r *LogRepository) Save(logData *models.LogMsg) RepositoryResult {
	err := r.db.Save(logData).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: logData}
}
