package svc

import (
	"context"

	"{{.Package}}/config"
)

type Session struct {
	Ctx    context.Context
	Config *config.Config
}
