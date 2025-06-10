package service

import (
	"github.com/IvanDrf/units/internal/models"
	"gorm.io/gorm"
)

type ConvertRepo interface {
	CreateConversion(conv *models.Responce) error

	GetAllConversions() ([]models.Responce, error)
	GetConversionByID(id uint) (models.Responce, error)

	UpdateConversion(conv *models.Responce) error
	DeleteConversion(id uint) error
}

type convRepo struct {
	database *gorm.DB
}

func NewConvertRepo(database *gorm.DB) ConvertRepo {
	return &convRepo{database: database}
}

func (r *convRepo) CreateConversion(conv *models.Responce) error {
	return r.database.Create(&conv).Error
}

func (r *convRepo) GetAllConversions() ([]models.Responce, error) {
	history := []models.Responce{}
	err := r.database.Find(&history).Error

	return history, err
}

func (r *convRepo) GetConversionByID(id uint) (models.Responce, error) {
	conv := models.Responce{}
	err := r.database.Find(&conv, id).Error

	return conv, err
}

func (r *convRepo) UpdateConversion(conv *models.Responce) error {
	return r.database.Save(&conv).Error
}

func (r *convRepo) DeleteConversion(id uint) error {
	return r.database.Delete(&models.Responce{}, id).Error
}
