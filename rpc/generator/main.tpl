package main

import (
	"flag"
	"fmt"

	{{.imports}}
    "github.com/uber-go/fx"
)


func main() {
	fx.New(
{{range .serviceNames}}       {{.ServerPkg}}.Module,
{{end}}

        fx.Provide(config.NewConfig),
    ).Run()
}
