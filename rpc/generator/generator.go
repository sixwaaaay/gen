package generator

import (
	"github.com/sirupsen/logrus"
	conf "github.com/sixwaaaay/gen/config"
	"github.com/sixwaaaay/gen/env"
)

// Generator defines the environment needs of rpc service generation
type Generator struct {
	log     *logrus.Logger
	cfg     *conf.Config
	verbose bool
}

// NewGenerator returns an instance of Generator
func NewGenerator(style string, verbose bool) *Generator {
	cfg, err := conf.NewConfig(style)
	if err != nil {
		logrus.Fatalln(err)
	}

	return &Generator{
		log:     logrus.StandardLogger(),
		cfg:     cfg,
		verbose: verbose,
	}
}

// Prepare provides environment detection generated by rpc service,
// including go environment, protoc, whether protoc-gen-go is installed or not
func (g *Generator) Prepare() error {
	return env.Prepare(true, true, g.verbose)
}