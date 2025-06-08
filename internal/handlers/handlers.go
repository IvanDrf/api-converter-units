package handlers

import (
	"net/http"

	"github.com/IvanDrf/units/internal/convert"
	"github.com/IvanDrf/units/internal/database"
	"github.com/IvanDrf/units/internal/models"
	"github.com/labstack/echo/v4"
)

func PostHandler(ctx echo.Context) error {
	req := models.Request{}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid body request"})
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
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := database.Database.Create(&result).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}
