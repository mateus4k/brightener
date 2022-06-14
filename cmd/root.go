package cmd

import (
	"os"

	"github.com/mateus4k/brightener/handler"
	"github.com/spf13/cobra"
)

var input string
var brightness float64
var rootCmd = &cobra.Command{
	Use:   "brightener",
	Short: "Adjust the brightness of all images within a folder",
	Long: `Adjust the brightness of all images within a folder.
Usage:
brightener --brightness=20 --input=./images`,
	Run: run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&input, "input", "i", "", "Source directory to read from")
	rootCmd.Flags().Float64VarP(&brightness, "brightness", "b", 0, "Brightness percentage to change the original image. The percentage must be in range (-100, 100). The percentage = 0 gives the original image. The percentage = -100 gives solid black image. The percentage = 100 gives solid white image.")
	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("brightness")
}

func run(cmd *cobra.Command, args []string) {
	handler.Handle(input, brightness)
}
