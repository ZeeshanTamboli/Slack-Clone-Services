package controllers

import (
	"github.com/ZeeshanTamboli/slack-clone-services/api/middlewares"
)

// InitializeRoutes : Initialize all the endpoints
func (s *Server) InitializeRoutes() {
	// User routes
	s.Router.HandleFunc("/api/user/createWorkspace", middlewares.SetMiddlewareJSON(s.CreateWorkspace)).Methods("POST")
}
