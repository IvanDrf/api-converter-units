package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IvanDrf/units/internal/convert"
	"github.com/IvanDrf/units/internal/database"
	"github.com/IvanDrf/units/internal/models"
	"github.com/labstack/echo/v4"
)

func (s *Server) PostHandler(ctx echo.Context) error {
	req := models.Request{}

	if err := ctx.Bind(&req); err != nil {
		s.Logger.Error("post: invalid body req")
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body request"})
	}

	result := models.Responce{
		UnitsType: req.UnitsType,
		Units:     req.Units,
		Value:     req.Value,

		NewUnits: req.NewUnits,
		NewValue: -1,
	}

	var err error
	result.NewValue, err = convert.Convert(&req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("post: %s", err.Error()))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := database.Database.Create(&result).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	s.Logger.Info("post: success")
	return ctx.JSON(http.StatusOK, result)
}

func (s *Server) GetHandler(ctx echo.Context) error {
	history := []models.Responce{}

	if err := database.Database.Find(&history).Error; err != nil {
		s.Logger.Error("get: db error")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	s.Logger.Info("get: success")
	return ctx.JSON(http.StatusOK, history)
}

func (s *Server) PatchHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		s.Logger.Error(fmt.Sprintf("patch: id is not a number %v", ctx.Param("id")))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "id is not a number"})
	}

	newConversion := models.Request{}
	if err := ctx.Bind(&newConversion); err != nil {
		s.Logger.Error("patch: invalid body req")
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body request"})
	}

	oldConversion := models.Responce{}
	if err := database.Database.First(&oldConversion, id).Error; err != nil {
		s.Logger.Info("patch: could not find conv")
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "could not find conversion"})
	}

	newValue, err := convert.Convert(&newConversion)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("patch: %s", err.Error()))
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	oldConversion.UnitsType = newConversion.UnitsType
	oldConversion.Units = newConversion.Units
	oldConversion.Value = newConversion.Value

	oldConversion.NewUnits = newConversion.NewUnits
	oldConversion.NewValue = newValue

	if err := database.Database.Save(&oldConversion).Error; err != nil {
		s.Logger.Error("patch: could not save update in db")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not save update"})
	}

	s.Logger.Info("patch: success")
	return ctx.JSON(http.StatusOK, oldConversion)
}

func (s *Server) DeleteHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		s.Logger.Error(fmt.Sprintf("delete: id is not a number %v", ctx.Param("id")))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "id is not a number"})
	}

	if err := database.Database.First(&models.Responce{}, id).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "incorrect id"})
	}

	if err := database.Database.Delete(&models.Responce{}, id).Error; err != nil {
		s.Logger.Error("delete: could not delete conv in db")
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "could not delete conversion"})
	}

	s.Logger.Info("delete: success")
	return ctx.NoContent(http.StatusNoContent)
}
