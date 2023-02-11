package generator

import "os/exec"

func GenerateProto(filename string) {
	cmd := exec.Command("protoc", "--go_out=plugins=grpc:.", filename)
	cmd.Run()
}