package sysstat

import (
	"fmt"
	"log"
	"os"

	"github.com/drizzleent/goprojectgen/internal/create"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sysstat",
	Short: "machine statistic",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Привет! Используй sysstat to show machine statistic.")
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
		create.CreateBaseStruct(name)
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
