package generator

import (
	"fmt"
	"github.com/sixwaaaay/gen/rpc/execx"
	"github.com/sixwaaaay/gen/util/stringx"
	"github.com/stretchr/testify/assert"
	"go/build"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRpcGenerate(t *testing.T) {
	_ = Clean()
	g := NewGenerator("gozero", true)
	err := g.Prepare()
	assert.NoError(t, err)
	projectName := stringx.Rand()
	src := filepath.Join(build.Default.GOPATH, "src")
	_, err = os.Stat(src)
	if err != nil {
		return
	}

	projectDir := filepath.Join(src, projectName)
	srcDir := projectDir
	defer func() {
		_ = os.RemoveAll(srcDir)
	}()
	common, err := filepath.Abs(".")
	assert.Nil(t, err)

	// case go path
	t.Run("GOPATH", func(t *testing.T) {
		ctx := &RpcContext{
			Src: "./test.proto",
			ProtocCmd: fmt.Sprintf("protoc -I=%s test.proto --go_out=%s --go_opt=Mbase/common.proto=./base --go-grpc_out=%s",
				common, projectDir, projectDir),
			IsGooglePlugin: true,
			GoOutput:       projectDir,
			GrpcOutput:     projectDir,
			Output:         projectDir,
		}
		err = g.Generate(ctx)
		assert.Nil(t, err)
		_, err = execx.Run("go test "+projectName, projectDir)
		if err != nil {
			assert.True(t, func() bool {
				return strings.Contains(err.Error(),
					"not in GOROOT") || strings.Contains(err.Error(), "cannot find package")
			}())
		}
	})

	// case go mod
	t.Run("GOMOD", func(t *testing.T) {
		workDir := projectDir
		name := filepath.Base(projectDir)
		_, err = execx.Run("go mod init "+name, workDir)
		assert.NoError(t, err)

		projectDir = filepath.Join(workDir, projectName)
		ctx := &RpcContext{
			Src: "./test.proto",
			ProtocCmd: fmt.Sprintf("protoc -I=%s test.proto --go_out=%s --go_opt=Mbase/common.proto=./base --go-grpc_out=%s",
				common, projectDir, projectDir),
			IsGooglePlugin: true,
			GoOutput:       projectDir,
			GrpcOutput:     projectDir,
			Output:         projectDir,
		}
		err = g.Generate(ctx)
		assert.Nil(t, err)
	})
}
