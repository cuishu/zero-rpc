package generator

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

func genFileOverwrite(filename, tmpl string, spec any) {
	t, err := template.New(filename).Parse(tmpl)
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

func genFile(filename, tmpl string, spec any) {
	if _, err := os.Stat(filename); err == nil {
		return
	}
	t, err := template.New(filename).Parse(tmpl)
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

func genMain(spec *Spec) {
	genFileOverwrite("main.go", spec.Template.Main, spec)
}

func genSession(spec *Spec) {
	filename := "svc/session.go"
	genFileOverwrite(filename, spec.Template.Session, spec)
}

func genSvc(spec *Spec) {
	filename := "svc/svc.go"
	genFile(filename, spec.Template.Svc, spec)
}

func genServer(spec *Spec) {
	filename := "server/server.go"
	genFileOverwrite(filename, spec.Template.Server, spec)
}

func genClient(spec *Spec) {
	filename := "proto/client.go"
	genFileOverwrite(filename, spec.Template.Client, spec)
}

func genLogic(spec *Spec) {
	for _, logic := range spec.Service.Methods {
		filename := fmt.Sprintf("logic/%s.go", logic.Name)
		if _, err := os.Stat(filename); err == nil {
			continue
		}
		logic.Module = spec.Module.Name
		genFile(filename, spec.Template.Logic, logic)
	}
}

func genConfig(spec *Spec) {
	filename := "config/config.go"
	genFile(filename, spec.Template.Config, spec)
}

func genConfigFile(spec *Spec) {
	filename := "config/config.yaml"
	genFile(filename, spec.Template.ConfigFile, spec)
	genFile("config/config.yaml.example", spec.Template.ConfigFile, spec)
}

func genBuildSH(spec *Spec) {
	filename := "build.sh"
	if _, err := os.Stat(filename); err == nil {
		return
	}
	os.WriteFile(filename, []byte(spec.Template.BuildSH), 0644)
}

func genMakefile(spec *Spec) {
	filename := "Makefile"
	if _, err := os.Stat(filename); err == nil {
		return
	}
	os.WriteFile(filename, []byte(spec.Template.Makefile), 0644)
}

func genVersion() {
	filename := "VERSION"
	if _, err := os.Stat(filename); err == nil {
		return
	}
	os.WriteFile(filename, []byte("v0.0.1"), 0644)
}

func genDockerFile(spec *Spec) {
	filename := "Dockerfile"
	genFile(filename, spec.Template.Dockerfile, spec)
}

func genGitIgnore(spec *Spec) {
	filename := ".gitignore"
	if _, err := os.Stat(filename); err == nil {
		return
	}
	t, err := template.New(filename).Parse(spec.Template.GitIgnore)
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

func genPostgres(spec *Spec) {
	genFile("svc/postgres.go", spec.Template.Postgres, spec)
}

func genRedis(spec *Spec) {
	genFile("svc/redis.go", spec.Template.Redis, spec)
}

func GenerateCode(spec *Spec) {
	genMain(spec)
	genSession(spec)
	genSvc(spec)
	genServer(spec)
	genClient(spec)
	genLogic(spec)
	genConfig(spec)
	genConfigFile(spec)
	genBuildSH(spec)
	genMakefile(spec)
	genVersion()
	genDockerFile(spec)
	genGitIgnore(spec)
	genPostgres(spec)
	genRedis(spec)

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
