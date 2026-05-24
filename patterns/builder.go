package main

import (
	"errors"
	"time"
)

type ServerBuilder struct {
	server Server
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (sb *ServerBuilder) SetHost(host string) *ServerBuilder {
	sb.server.Host = host
	return sb
}

func (sb *ServerBuilder) SetPort(port int) *ServerBuilder {
	sb.server.Port = port
	return sb
}

func (sb *ServerBuilder) SetTimeout(timeout time.Duration) *ServerBuilder {
	sb.server.Timeout = timeout
	return sb
}

func (sb *ServerBuilder) Build() (Server, error) {
	if sb.server.Host == "" {
		return Server{}, errors.New("host is required")
	}
	return sb.server, nil
}
