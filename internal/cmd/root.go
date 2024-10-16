package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "MyApp es una aplicación de ejemplo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ejecutando MyApp")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
