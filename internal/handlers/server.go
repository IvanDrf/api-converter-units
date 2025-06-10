package handlers

import (
	"log/slog"
	"os"

	"github.com/IvanDrf/units/internal/database"
	"github.com/IvanDrf/units/internal/logger"
	"github.com/IvanDrf/units/internal/service"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Server *echo.Echo
	Logger *slog.Logger

	Service service.ConvertService
}

func (s *Server) Start() {
	if err := s.Server.Start(":8080"); err != nil {
		s.Logger.Error(err.Error())
		os.Exit(1)
	}
}

func InitServer() *Server {
	return &Server{
		Server:  echo.New(),
		Logger:  logger.InitLogger(),
		Service: service.NewConvertService(service.NewConvertRepo(database.InitDB())),
	}
}
