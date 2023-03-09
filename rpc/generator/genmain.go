package generator

import (
	_ "embed"
	"fmt"
	"github.com/sixwaaaay/gen/util/collection"
	"path/filepath"
	"strings"

	conf "github.com/sixwaaaay/gen/config"
	"github.com/sixwaaaay/gen/rpc/parser"
	"github.com/sixwaaaay/gen/util"
	"github.com/sixwaaaay/gen/util/format"
	"github.com/sixwaaaay/gen/util/pathx"
)

//go:embed main.tpl
var mainTemplate string

type MainServiceTemplateData struct {
	Service   string
	ServerPkg string
	Pkg       string
}

// GenMain generates the main file of the rpc service, which is a rpc service program call entry
func (g *Generator) GenMain(ctx DirContext, proto parser.Proto, cfg *conf.Config,
	c *RpcContext) error {
	mainFilename, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	fileName := filepath.Join(ctx.GetMain().Filename, fmt.Sprintf("%v.go", mainFilename))
	imports := collection.NewSet()
	configImport := fmt.Sprintf(`"%v"`, ctx.GetConfig().Package)
	serverImport := fmt.Sprintf(`"%v"`, ctx.GetServer().Package)
	imports.AddStr(configImport, serverImport)

	var serviceNames []MainServiceTemplateData
	for _, e := range proto.Service {
		var (
			remoteImport string
			serverPkg    string
		)
		if !c.Multiple {
			serverPkg = "server"
			remoteImport = fmt.Sprintf(`"%v"`, ctx.GetServer().Package)
		} else {
			childPkg, err := ctx.GetServer().GetChildPackage(e.Name)
			if err != nil {
				return err
			}

			serverPkg = filepath.Base(childPkg + "Server")
			remoteImport = fmt.Sprintf(`%s "%v"`, serverPkg, childPkg)
		}
		imports.AddStr(remoteImport)
		serviceNames = append(serviceNames, MainServiceTemplateData{
			Service:   parser.CamelCase(e.Name),
			ServerPkg: serverPkg,
			Pkg:       proto.PbPackage,
		})
	}

	text, err := pathx.LoadTemplate(category, mainTemplateFile, mainTemplate)
	if err != nil {
		return err
	}

	etcFileName, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	return util.With("main").GoFmt(true).Parse(text).SaveTo(map[string]any{
		"serviceName":  etcFileName,
		"imports":      strings.Join(imports.KeysStr(), pathx.NL),
		"pkg":          proto.PbPackage,
		"serviceNames": serviceNames,
	}, fileName, false)
}
