package {{.packageName}}

import (
	"context"

	{{.imports}}
    "go.uber.org/fx"

)

type {{.logicName}} struct {
    conf  config.Config
}

type {{.logicName}}Option struct {
    fx.In
    Config config.Config
}

func New{{.logicName}}(opt {{.logicName}}Option) *{{.logicName}} {
	return &{{.logicName}}{
        conf: opt.Config,
	}
}
{{.functions}}
