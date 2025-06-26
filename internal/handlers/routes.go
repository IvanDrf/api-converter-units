package handlers

func (s *Server) RegisterRoutes() {
	s.Server.POST("/api", s.PostHandler)
	s.Server.GET("/api", s.GetHandler)
	s.Server.PATCH("/api/conversions/:id", s.PatchHandler)
	s.Server.DELETE("/api/conversions/:id", s.DeleteHandler)
}
