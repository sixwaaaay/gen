package main

import (

	{{.imports}}
    "go.uber.org/fx"
)


func main() {
	fx.New(
{{range .serviceNames}}       {{.ServerPkg}}.Module,
{{end}}

        fx.Provide(config.NewConfig),
    ).Run()
}
