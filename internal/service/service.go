package service

import (
	"github.com/IvanDrf/units/internal/convert"
	"github.com/IvanDrf/units/internal/models"
)

type ConvertService interface {
	CreateConversion(req *models.Request) (models.Responce, error)

	GetAllConversions() ([]models.Responce, error)
	GetConversionByID(id uint) (models.Responce, error)

	UpdateConversion(id uint, newConv *models.Request) (models.Responce, error)
	DeleteConversion(id uint) error
}

type convService struct {
	repo ConvertRepo
}

func NewConvertService(r ConvertRepo) ConvertService {
	return &convService{repo: r}
}

func (service *convService) CreateConversion(req *models.Request) (models.Responce, error) {
	result, err := convert.CreateConversion(req)
	if err != nil {
		return models.Responce{}, err
	}

	if err = service.repo.CreateConversion(&result); err != nil {
		return models.Responce{}, err
	}

	return result, err
}

func (service *convService) GetAllConversions() ([]models.Responce, error) {
	return service.repo.GetAllConversions()
}

func (service *convService) GetConversionByID(id uint) (models.Responce, error) {
	return service.repo.GetConversionByID(id)
}

func (service *convService) UpdateConversion(id uint, req *models.Request) (models.Responce, error) {
	conv, err := convert.CreateConversion(req)
	if err != nil {
		return models.Responce{}, err
	}

	oldConv, err := service.repo.GetConversionByID(id)
	if err != nil {
		return models.Responce{}, err
	}

	conv.ID = oldConv.ID

	if err = service.repo.UpdateConversion(&conv); err != nil {
		return models.Responce{}, err
	}

	return conv, err
}

func (service *convService) DeleteConversion(id uint) error {
	return service.repo.DeleteConversion(id)
}
