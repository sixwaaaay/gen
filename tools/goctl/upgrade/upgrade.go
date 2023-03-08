package upgrade

import (
	"fmt"
	"runtime"

	"github.com/sixwaaaay/gen/rpc/execx"
	"github.com/spf13/cobra"
)

// upgrade gets the latest gen by
// go install github.com/sixwaaaay/gen@latest
func upgrade(_ *cobra.Command, _ []string) error {
	cmd := `GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/sixwaaaay/gen@latest`
	if runtime.GOOS == "windows" {
		cmd = `set GOPROXY=https://goproxy.cn,direct && go install github.com/sixwaaaay/gen@latest`
	}
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}
