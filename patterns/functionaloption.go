package main

import (
	"time"
)

type ServerOption func(*Server)

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
}

func NewServer(options ...ServerOption) *Server {
	server := &Server{
		Host:    "localhost",
		Port:    8080,
		Timeout: 30 * time.Second,
	}

	for _, option := range options {
		option(server)
	}

	return server
}

func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
