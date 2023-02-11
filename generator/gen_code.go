package generator

import (
	"io"
	"os"
	"text/template"
)

func genMain(spec *Spec) {
	t, err := template.New("main.go").Parse(spec.Template.Main)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout
	if err := t.Execute(w, spec); err != nil {
		panic(err)
	}
}

func genSession(spec *Spec) {
	t, err := template.New("session.go").Parse(spec.Template.Session)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout
	if err := t.Execute(w, spec); err != nil {
		panic(err)
	}
}

func genServer(spec *Spec) {
	t, err := template.New("server.go").Parse(spec.Template.Server)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout
	if err := t.Execute(w, spec); err != nil {
		panic(err)
	}
}

func genLogic(spec *Spec) {
	t, err := template.New("logic.go").Parse(spec.Template.Logic)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout
	for _, logic := range spec.Service.Methods {
		if err := t.Execute(w, logic); err != nil {
			panic(err)
		}
	}
}

func genConfig(spec *Spec) {
	t, err := template.New("config.go").Parse(spec.Template.Config)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout
	if err := t.Execute(w, spec); err != nil {
		panic(err)
	}
}

func genConfigFile(spec *Spec) {
	t, err := template.New("config.yaml").Parse(spec.Template.ConfigFile)
	if err != nil {
		panic(err)
	}
	var w io.Writer = os.Stdout
	if err := t.Execute(w, spec); err != nil {
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
