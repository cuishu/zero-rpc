package svc

import (
	"context"

	"template/config"
)

type Session struct {
	Ctx    context.Context
	Config *config.Config
}
