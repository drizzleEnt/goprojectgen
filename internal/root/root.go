package root

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prg",
	Short: "golang project generation",
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init project",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Flags().GetString("name")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)

}
