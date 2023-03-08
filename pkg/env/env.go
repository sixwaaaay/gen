package env

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sixwaaaay/gen/internal/version"
	sortedmap "github.com/sixwaaaay/gen/pkg/collection"
	"github.com/sixwaaaay/gen/pkg/protoc"
	"github.com/sixwaaaay/gen/pkg/protocgengo"
	"github.com/sixwaaaay/gen/pkg/protocgengogrpc"
	"github.com/sixwaaaay/gen/util/pathx"
)

var genEnv *sortedmap.SortedMap

const (
	OS                     = "TOOL_OS"
	Arch                   = "TOOL_ARCH"
	Home                   = "TOOL_HOME"
	Debug                  = "TOOL_DEBUG"
	Cache                  = "TOOL_CACHE"
	Version                = "TOOL_VERSION"
	ProtocVersion          = "PROTOC_VERSION"
	ProtocGenGoVersion     = "PROTOC_GEN_GO_VERSION"
	ProtocGenGoGRPCVersion = "PROTO_GEN_GO_GRPC_VERSION"

	envFileDir = "env"
)

// init initializes the gen environment variables, the environment variables of the function are set in order,
// please do not change the logic order of the code.
func init() {
	defaultHome, err := pathx.GetDefaultHome()
	if err != nil {
		log.Fatalln(err)
	}
	genEnv = sortedmap.New()
	genEnv.SetKV(OS, runtime.GOOS)
	genEnv.SetKV(Arch, runtime.GOARCH)
	existsEnv := readEnv(defaultHome)
	if existsEnv != nil {
		genHome, ok := existsEnv.GetString(Home)
		if ok && len(genHome) > 0 {
			genEnv.SetKV(Home, genHome)
		}
		if debug := existsEnv.GetOr(Debug, "").(string); debug != "" {
			if strings.EqualFold(debug, "true") || strings.EqualFold(debug, "false") {
				genEnv.SetKV(Debug, debug)
			}
		}
		if value := existsEnv.GetStringOr(Cache, ""); value != "" {
			genEnv.SetKV(Cache, value)
		}
	}
	if !genEnv.HasKey(Home) {
		genEnv.SetKV(Home, defaultHome)
	}
	if !genEnv.HasKey(Debug) {
		genEnv.SetKV(Debug, "False")
	}

	if !genEnv.HasKey(Cache) {
		cacheDir, _ := pathx.GetCacheDir()
		genEnv.SetKV(Cache, cacheDir)
	}

	genEnv.SetKV(Version, version.BuildVersion)
	protocVer, _ := protoc.Version()
	genEnv.SetKV(ProtocVersion, protocVer)

	protocGenGoVer, _ := protocgengo.Version()
	genEnv.SetKV(ProtocGenGoVersion, protocGenGoVer)

	protocGenGoGrpcVer, _ := protocgengogrpc.Version()
	genEnv.SetKV(ProtocGenGoGRPCVersion, protocGenGoGrpcVer)
}

func Print() string {
	return strings.Join(genEnv.Format(), "\n")
}

func Get(key string) string {
	return GetOr(key, "")
}

func GetOr(key, def string) string {
	return genEnv.GetStringOr(key, def)
}

func readEnv(genHome string) *sortedmap.SortedMap {
	envFile := filepath.Join(genHome, envFileDir)
	data, err := os.ReadFile(envFile)
	if err != nil {
		return nil
	}
	dataStr := string(data)
	lines := strings.Split(dataStr, "\n")
	sm := sortedmap.New()
	for _, line := range lines {
		_, _, err = sm.SetExpression(line)
		if err != nil {
			continue
		}
	}
	return sm
}

func WriteEnv(kv []string) error {
	defaultGenHome, err := pathx.GetDefaultHome()
	if err != nil {
		log.Fatalln(err)
	}
	data := sortedmap.New()
	for _, e := range kv {
		_, _, err := data.SetExpression(e)
		if err != nil {
			return err
		}
	}
	data.RangeIf(func(key, value any) bool {
		switch key.(string) {
		case Home, Cache:
			path := value.(string)
			if !pathx.FileExists(path) {
				err = fmt.Errorf("[writeEnv]: path %q is not exists", path)
				return false
			}
		}
		if genEnv.HasKey(key) {
			genEnv.SetKV(key, value)
			return true
		} else {
			err = fmt.Errorf("[writeEnv]: invalid key: %v", key)
			return false
		}
	})
	if err != nil {
		return err
	}
	envFile := filepath.Join(defaultGenHome, envFileDir)
	return os.WriteFile(envFile, []byte(strings.Join(genEnv.Format(), "\n")), 0o777)
}
