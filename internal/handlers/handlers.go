package handlers

import (
	"net/http"
	"strconv"

	"github.com/IvanDrf/units/internal/models"
	"github.com/labstack/echo/v4"
)

const (
	post   = "POST"
	get    = "GET"
	patch  = "PATCH"
	delete = "DELETE"
)

func (s *Server) PostHandler(ctx echo.Context) error {
	req := models.Request{}
	if err := ctx.Bind(&req); err != nil {
		s.Logger.InfoLevel(post, "invalid body req")

		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body req"})
	}

	result, err := s.Service.CreateConversion(&req)
	if err != nil {
		s.Logger.ErrorLevel(post, err.Error())

		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	s.Logger.InfoLevel(post, "success")

	return ctx.JSON(http.StatusOK, result)
}

func (s *Server) GetHandler(ctx echo.Context) error {
	history, err := s.Service.GetAllConversions()
	if err != nil {
		s.Logger.ErrorLevel(get, err.Error())

		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	s.Logger.InfoLevel(get, "success")

	return ctx.JSON(http.StatusOK, history)
}

func (s *Server) PatchHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		s.Logger.ErrorLevel(patch, "id is not a number")

		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "id is not a number"})
	}

	newReq := models.Request{}
	if err := ctx.Bind(&newReq); err != nil {
		s.Logger.ErrorLevel(patch, "invalid body req")

		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body req"})
	}

	conv, err := s.Service.UpdateConversion(uint(id), &newReq)
	if err != nil {
		s.Logger.InfoLevel(patch, err.Error())

		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	s.Logger.InfoLevel(patch, "success")

	return ctx.JSON(http.StatusOK, conv)
}

func (s *Server) DeleteHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		s.Logger.ErrorLevel(delete, "id is not a number")

		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "id is not a number"})
	}

	if err = s.Service.DeleteConversion(uint(id)); err != nil {
		s.Logger.ErrorLevel(delete, err.Error())

		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	s.Logger.InfoLevel(delete, "success")

	return ctx.NoContent(http.StatusNoContent)
}
