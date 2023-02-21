package generator

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

func GenerateProto(filename string) {
	cmd := exec.Command("protoc", "--go_out=plugins=grpc:.", filename)
	cmd.Run()
}

func GenerateProtoTemplate(spec *Spec) {
	filename := fmt.Sprintf("%s.proto", spec.ShortModule)
	if _, err := os.Stat(filename); err == nil {
		return
	}
	t, err := template.New(filename).Parse(spec.Template.ExampleProtoTmpl)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	if err := t.Execute(file, spec); err != nil {
		panic(err)
	}
}
