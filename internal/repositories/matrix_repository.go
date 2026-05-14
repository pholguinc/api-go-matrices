package repositories

import (
	"github.com/pholguinc/api-go-matrices/internal/models"
	"gorm.io/gorm"
)

type MatrixRepository interface {
	SaveRecord(record *models.MatrixRecord) error
	GetHistoryByUserID(userID string) ([]models.MatrixRecord, error)
}

type matrixRepository struct {
	db *gorm.DB
}

func NewMatrixRepository(db *gorm.DB) MatrixRepository {
	return &matrixRepository{db: db}
}

func (r *matrixRepository) SaveRecord(record *models.MatrixRecord) error {
	return r.db.Create(record).Error
}

func (r *matrixRepository) GetHistoryByUserID(userID string) ([]models.MatrixRecord, error) {
	var records []models.MatrixRecord
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&records).Error
	return records, err
}
