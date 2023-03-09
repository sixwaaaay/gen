{{.head}}

package server

import (
	{{if .notStream}}"context"{{end}}
    "google.golang.org/grpc"
    "net"
	{{.imports}}
	"go.uber.org/fx"

)

type {{.server}}Server struct {
	{{.params}}
	{{.unimplementedServer}}
}

type ServerOption struct {
    fx.In
    {{.params}}
}

func New{{.server}}Server(opt ServerOption) *{{.server}}Server {
	return &{{.server}}Server{
	    {{.initParams}}
	}
}

{{.funcs}}


var Module = fx.Module("server",
	fx.Provide(
	    New{{.server}}Server,
        {{.constructors}}
	),
	fx.Invoke(func(conf config.Config, s *{{.server}}Server, lf fx.Lifecycle) error {
		listener, err := net.Listen("tcp", conf.ListenOn)
		if err != nil {
			return err
		}
		grpcServer := grpc.NewServer()
		{{.pkg}}.Register{{.server}}Server(grpcServer, s)

		// graceful startup and shutdown

        lf.Append(fx.Hook{
            OnStart: func(ctx context.Context) error {
                go func() {
                    if err := grpcServer.Serve(listener); err != nil {
                        panic(err)
                    }
                }()
                return nil
            },
            OnStop: func(ctx context.Context) error {
                grpcServer.Stop()
                return nil
            },
            })
        return nil
	}),
)



