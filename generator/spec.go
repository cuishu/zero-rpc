package generator

import (
	"strings"
)

type Method struct {
	Module  string
	Comment string
	Name    string
	Param   string
	Return  string
}

type Service struct {
	Comment string
	Name    string
	Methods []Method
}

type Template struct {
	Main             string
	Session          string
	Svc              string
	Server           string
	Client           string
	Logic            string
	Config           string
	ConfigFile       string
	BuildSH          string
	Makefile         string
	Dockerfile       string
	GitIgnore        string
	ExampleProtoTmpl string
}

type Module struct {
	Name  string
	Short string
}

func (m *Module) Set(name string) {
	m.Name = name
	slice := strings.Split(name, "/")
	m.Short = slice[len(slice)-1]
}

type Spec struct {
	Module   Module
	Package  string
	Service  Service
	Template Template
}
