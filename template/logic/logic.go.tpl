package logic

import (
	"template/proto"
	"template/svc"
)
{{.Comment}}
func {{.Name}}(sess *svc.Session, input *proto.{{.Param}}) (*proto.{{.Return}}, error) {
	var resp proto.{{.Param}}
	return &resp, nil
}
