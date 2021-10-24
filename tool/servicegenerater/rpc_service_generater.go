package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"
)

// -s server_name
// -p package_name
// -v proto version
type ServerParams struct {
	ProtoVersion  int
	ServerName    string
	PackageName   string
	ProtoFileName string
}

func main() {
	var serverName = flag.String("s", "", "Input server name")
	var packageName = flag.String("p", "", "Input package name")
	var protoVersion = flag.Int("v", 3, "Input proto version")
	var protoFileName = flag.String("f", defaultFileName(), "Input proto version")

	flag.Parse()

	s := ServerParams{
		ProtoVersion:  *protoVersion,
		ServerName:    *serverName,
		PackageName:   *packageName,
		ProtoFileName: *protoFileName,
	}
	s.ServerName = strings.Title(s.ServerName)

	generate(s)
}

func generate(s ServerParams) {
	tmpl, err := template.New("rpc_service").Parse(serverProtoTemplate)
	if err != nil {
		panic(err)
	}

	fmt.Println(s.ProtoFileName)
	if _, err := os.Stat(s.ProtoFileName); !os.IsNotExist(err) {
		fmt.Println("file name is existed")
		return
	}

	f, err := os.Create(s.ProtoFileName)
	if err != nil {
		fmt.Println("file is not existed")
		return
	}
	defer f.Close()

	err = tmpl.Execute(f, s)
	if err != nil {
		panic(err)
	}
}

func defaultFileName() string {
	return fmt.Sprintf("my_rpc_service_%d.proto", time.Now().UnixNano())
}

var serverProtoTemplate = `syntax = "proto{{.ProtoVersion}}";

option go_package = "{{.PackageName}}";

package plutoservice;

service {{.ServerName}} {
  rpc Pong(BlankParams) returns (PongResponse) {}
}

message BlankParams{}

message PongResponse {
  string serverName = 1;
}
`
