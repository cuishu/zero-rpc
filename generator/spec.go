package generator

type Method struct {
	Module string
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
	Main       string
	Session    string
	Svc        string
	Server     string
	Client     string
	Logic      string
	Config     string
	ConfigFile string
	BuildSH    string
	Makefile   string
	Dockerfile string
	GitIgnore  string
}

type Spec struct {
	Module   string
	Package  string
	Service  Service
	Template Template
}
