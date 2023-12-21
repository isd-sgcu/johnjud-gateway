package auth

import (
	"github.com/isd-sgcu/johnjud-gateway/src/app/handler/auth"
	"github.com/isd-sgcu/johnjud-gateway/src/app/router"
	"github.com/isd-sgcu/johnjud-gateway/src/config"
)

type Guard struct {
	service    auth.Service
	excludes   map[string]struct{}
	conf       config.App
	isValidate bool
}

func NewAuthGuard(s auth.Service, e map[string]struct{}, conf config.App) Guard {
	return Guard{
		service:    s,
		excludes:   e,
		conf:       conf,
		isValidate: true,
	}
}

func (m *Guard) Use(ctx *router.FiberCtx) {
	m.isValidate = true

	m.Validate(ctx)

	if !m.isValidate {
		return
	}

	ctx.Next()

}

func (m *Guard) Validate(ctx *router.FiberCtx) {
	ctx.Next()
}