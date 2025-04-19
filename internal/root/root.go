package root

import (
	"fmt"
	"log"
	"os"

	"github.com/drizzleent/goprojectgen/internal/root/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prg",
	Short: "golang project generation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Привет! Используй `prg init <projectname>` чтобы создать проект.")
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init project",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Printf("failed to read name flag: %s", err.Error())
			os.Exit(1)
		}
		commands.CreateBaseStruct(name)
		fmt.Printf("created new project %s\n", name)
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

	initCmd.Flags().StringP("name", "n", "app", "Project name")
}
