package logic

import (
	"{{.Module}}/proto"
	"{{.Module}}/svc"
)
{{.Comment}}
func {{.Name}}(sess *svc.Session, input *proto.{{.Param}}) (*proto.{{.Return}}, error) {
	var resp proto.{{.Return}}
	return &resp, nil
}
