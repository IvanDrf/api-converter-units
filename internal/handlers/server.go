package handlers

import (
	"github.com/IvanDrf/units/internal/database"
	"github.com/IvanDrf/units/internal/logger"
	"github.com/IvanDrf/units/internal/service"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Server *echo.Echo
	Logger logger.Logger

	Service service.ConvertService
}

func (s *Server) Start() {
	if err := s.Server.Start(":8080"); err != nil {
		s.Logger.Fatal(err.Error())
	}
}

func InitServer() *Server {
	return &Server{
		Server:  echo.New(),
		Logger:  logger.InitLogger(),
		Service: service.NewConvertService(service.NewConvertRepo(database.InitDB())),
	}
}
