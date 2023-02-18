package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	_ "embed"

	"github.com/cuishu/functools"
	"github.com/cuishu/zero-rpc/generator"
	"github.com/emicklei/proto"
)

var (
	protoFile string
	spec      generator.Spec
	//go:embed template/main.go.tpl
	mainTmpl string
	//go:embed template/svc/session.go.tpl
	sessionTmpl string
	//go:embed template/svc/svc.go.tpl
	svcTmpl string
	//go:embed template/server/server.go.tpl
	serverTmpl string
	//go:embed template/proto/client.go.tpl
	clientTmpl string
	//go:embed template/logic/logic.go.tpl
	logicTmpl string
	//go:embed template/config/config.go.tpl
	configTmpl string
	//go:embed template/config/config.yaml.gtpl
	configFileTmpl string
	//go:embed template/build.sh.gtpl
	buildSHTmpl string
	//go:embed template/Makefile.gtpl
	makefileTmpl string
	//go:embed template/Dockerfile.gtpl
	dockerfileTmpl string
	//go:embed template/gitignore.gtpl
	gitignoreTmpl string
)

func init() {
	flag.StringVar(&protoFile, "f", "", "proto filename")
	flag.Parse()

	if protoFile == "" {
		os.Exit(0)
	}
	spec.Template.Main = mainTmpl
	spec.Template.Session = sessionTmpl
	spec.Template.Svc = svcTmpl
	spec.Template.Server = serverTmpl
	spec.Template.Client = clientTmpl
	spec.Template.Logic = logicTmpl
	spec.Template.Config = configTmpl
	spec.Template.ConfigFile = configFileTmpl
	spec.Template.BuildSH = buildSHTmpl
	spec.Template.Makefile = makefileTmpl
	spec.Template.Dockerfile = dockerfileTmpl
	spec.Template.GitIgnore = gitignoreTmpl
}

func handleService(s *proto.Service) {
	if s.Comment != nil {
		spec.Service.Comment = strings.Join(functools.Map(func(v string) string {
			return "//" + v
		}, s.Comment.Lines), "\n")
	}
	spec.Service.Name = s.Name
}

func handlePackage(p *proto.Package) {
	spec.Package = p.Name
}

func handleRPC(rpc *proto.RPC) {
	var comment []string
	if rpc.Comment != nil {
		comment = functools.Map(func(v string) string {
			return "//" + v
		}, rpc.Comment.Lines)
	}
	spec.Service.Methods = append(spec.Service.Methods, generator.Method{
		Comment: strings.Join(comment, "\n"),
		Name:    rpc.Name,
		Param:   rpc.RequestType,
		Return:  rpc.ReturnsType,
	})
}

func mkdir() {
	os.MkdirAll("config", 0755)
	os.MkdirAll("logic", 0755)
	os.MkdirAll("server", 0755)
	os.MkdirAll("proto", 0755)
	os.MkdirAll("svc", 0755)
}

func init() {
	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("please run `go mod init` befor do this")
		os.Exit(0)
	}
	defer file.Close()
	data, err:=io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	slice := strings.Split(string(data), "\n")
	if len(slice) == 0 {
		panic("invalid file go.mod")
	}
	slice = strings.Split(slice[0], " ")
	if len(slice) != 2 {
		panic("invalid file go.mod")
	}
	spec.Module = slice[len(slice)-1]
}

func main() {
	file, err := os.Open(protoFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	definition, err := proto.NewParser(file).Parse()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-2)
	}

	proto.Walk(definition,
		proto.WithService(handleService),
		proto.WithPackage(handlePackage),
		proto.WithRPC(handleRPC),
	)
	mkdir()
	generator.GenerateProto(protoFile)
	generator.GenerateCode(&spec)
}
