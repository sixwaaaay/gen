package upgrade

import "github.com/spf13/cobra"

// Cmd describes an upgrade command.
var Cmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade gen to latest version",
	RunE:  upgrade,
}
