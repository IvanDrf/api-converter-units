package handlers

func (s *Server) RegisterRoutes() {
	s.Server.POST("/", s.PostHandler)
	s.Server.GET("/", s.GetHandler)
	s.Server.PATCH("/conversions/:id", s.PatchHandler)
	s.Server.DELETE("/conversions/:id", s.DeleteHandler)
}
