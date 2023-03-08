package generator

import (
	_ "embed"
	"io/ioutil"
	"os"
	"path/filepath"

	conf "github.com/sixwaaaay/gen/config"
	"github.com/sixwaaaay/gen/rpc/parser"
	"github.com/sixwaaaay/gen/util/format"
	"github.com/sixwaaaay/gen/util/pathx"
)

//go:embed config.tpl
var configTemplate string

// GenConfig generates the configuration structure definition file of the rpc service
func (g *Generator) GenConfig(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetConfig()
	configFilename, err := format.FileNamingFormat(cfg.NamingFormat, "config")
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, configFilename+".go")
	if pathx.FileExists(fileName) {
		return nil
	}

	text, err := pathx.LoadTemplate(category, configTemplateFileFile, configTemplate)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, []byte(text), os.ModePerm)
}
