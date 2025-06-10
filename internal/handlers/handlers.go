package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IvanDrf/units/internal/models"
	"github.com/labstack/echo/v4"
)

func (s *Server) PostHandler(ctx echo.Context) error {
	req := models.Request{}
	if err := ctx.Bind(&req); err != nil {
		s.Logger.Error("post: invalid body req")
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body request"})
	}

	result, err := s.Service.CreateConversion(&req)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("post: %s", err.Error()))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	s.Logger.Info("post: success")
	return ctx.JSON(http.StatusOK, result)
}

func (s *Server) GetHandler(ctx echo.Context) error {
	history, err := s.Service.GetAllConversions()
	if err != nil {
		s.Logger.Error(fmt.Sprintf("get: %s", err.Error()))
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

	newReq := models.Request{}
	if err := ctx.Bind(&newReq); err != nil {
		s.Logger.Error("patch: invalid body req")
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body request"})
	}

	conv, err := s.Service.UpdateConversion(uint(id), &newReq)
	if err != nil {
		s.Logger.Info(fmt.Sprintf("patch: %s", err.Error()))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	s.Logger.Info("patch: success")
	return ctx.JSON(http.StatusOK, conv)
}

func (s *Server) DeleteHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		s.Logger.Error(fmt.Sprintf("delete: id is not a number %v", ctx.Param("id")))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "id is not a number"})
	}

	if err = s.Service.DeleteConversion(uint(id)); err != nil {
		s.Logger.Error(fmt.Sprintf("delete: %s", err.Error()))
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	s.Logger.Info("delete: success")
	return ctx.NoContent(http.StatusNoContent)
}
