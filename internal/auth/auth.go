package auth

import "context"

type Auth struct {
	user  string `json:"user"`
	Token string `json:"token"`
	ctx   *context.Context
}

type authentication interface {
}
