package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/burbokop/forum-app/server/forums"
)

type HttpPortNumber int

// ChatApiServer configures necessary handlers and starts listening on a configured port.
type ForumApiServer struct {
	Port HttpPortNumber

	ListForumsHandler forums.HttpListForumsHandlerFunc
	AddUserHandler    forums.HttpAddUserHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *ForumApiServer) Start() error {
	if s.ListForumsHandler == nil {
		return fmt.Errorf("HTTP ListVmsHandler is not defined - cannot start")
	}
	if s.AddUserHandler == nil {
		return fmt.Errorf("HTTP DiscConnectionHandler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/forums", s.ListForumsHandler)
	handler.HandleFunc("/add_user", s.AddUserHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *ForumApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
