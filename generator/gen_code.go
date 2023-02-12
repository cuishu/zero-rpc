package generator

import (
	"fmt"
	"os"
	"text/template"
)

func genMain(spec *Spec) {
	t, err := template.New("main.go").Parse(spec.Template.Main)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile("main.go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	if err := t.Execute(file, spec); err != nil {
		panic(err)
	}
}

func genSession(spec *Spec) {
	filename := "svc/session.go"
	t, err := template.New("session.go").Parse(spec.Template.Session)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	if err := t.Execute(file, spec); err != nil {
		panic(err)
	}
}

func genServer(spec *Spec) {
	filename := "server/server.go"
	t, err := template.New("server.go").Parse(spec.Template.Server)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	if err := t.Execute(file, spec); err != nil {
		panic(err)
	}
}

func genLogic(spec *Spec) {
	t, err := template.New("logic.go").Parse(spec.Template.Logic)
	if err != nil {
		panic(err)
	}
	for _, logic := range spec.Service.Methods {
		filename := fmt.Sprintf("logic/%s.go", logic.Name)
		if _, err := os.Stat(filename); err == nil {
			continue
		}
		logic.Package = spec.Package
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		if err := t.Execute(file, logic); err != nil {
			panic(err)
		}
	}
}

func genConfig(spec *Spec) {
	filename := "config/config.go"
	if _, err := os.Stat(filename); err == nil {
		return
	}
	t, err := template.New("config.go").Parse(spec.Template.Config)
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

func genConfigFile(spec *Spec) {
	filename := "config/config.yaml"
	t, err := template.New("config.yaml").Parse(spec.Template.ConfigFile)
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

func GenerateCode(spec *Spec) {
	genMain(spec)
	genSession(spec)
	genServer(spec)
	genLogic(spec)
	genConfig(spec)
	genConfigFile(spec)
}
