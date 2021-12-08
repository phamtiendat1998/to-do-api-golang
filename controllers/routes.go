package controllers

import (
	"to-do-api-golang/middlewares"
)

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	s.Router.HandleFunc("/todos", middlewares.SetMiddlewareJSON(s.CreateTodo)).Methods("POST")
	s.Router.HandleFunc("/todos", middlewares.SetMiddlewareJSON(s.GetTodos)).Methods("GET")
	s.Router.HandleFunc("/todos/{id}", middlewares.SetMiddlewareJSON(s.GetToDo)).Methods("GET")
	s.Router.HandleFunc("/todos/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateToDo))).Methods("PUT")
	s.Router.HandleFunc("/todos/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteTodo)).Methods("DELETE")
}
