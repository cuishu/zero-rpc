package server

import (
	"context"
	"fmt"

	"{{.Package}}/config"
	"{{.Package}}/logic"
	"{{.Package}}/proto"
	"{{.Package}}/svc"
)

type {{.Service.Name}}Server struct {
	proto.SensitiveServer
	conf *config.Config
}

func New{{.Service.Name}}Server(conf *config.Config) *{{.Service.Name}}Server {
	return &SensitiveServer{
		conf: conf,
	}
}
{{range .Service.Methods}}
{{.Comment}}
func (server {{$.Service.Name}}Server) {{.Name}}(ctx context.Context, input *proto.{{.Param}}) (*proto.{{.Return}}, error) {
	return logic.{{.Name}}(&svc.Session{
		Config: server.conf,
		Ctx:    ctx,
	}, input)
}
{{end}}
