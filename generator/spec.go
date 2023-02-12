package generator

type Method struct {
	Package string
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
	Server     string
	Logic      string
	Config     string
	ConfigFile string
}

type Spec struct {
	Package  string
	Service  Service
	Template Template
}
