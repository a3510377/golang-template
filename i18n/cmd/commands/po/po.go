package po

import (
	"os"

	"test/i18n"
	"test/i18n/cmd/po"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "po",
	Short: "po",
}

func init() {
	Command.AddCommand(generatePoCommand)
}

var generatePoCommand = &cobra.Command{
	Use:   "generate [input folder]",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	Run:   generatePo,
}

func generatePo(cmd *cobra.Command, args []string) {
	folder := args[0]
	lang := ""

	if len(args) > 1 {
		lang = i18n.FormatLocale(args[1])
	}

	po := po.GeneratePo(folder, lang)
	po.Write(os.Stdout)
}
