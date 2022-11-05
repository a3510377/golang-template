package commands

import (
	"test/i18n/cmd/commands/po"

	"github.com/spf13/cobra"
)

var i18nCommand = &cobra.Command{
	Use:   "i18n",
	Short: "i18n",
}

func init() {
	i18nCommand.AddCommand(po.Command)
}

// Execute executes the i18n command
func Execute() error {
	return i18nCommand.Execute()
}
