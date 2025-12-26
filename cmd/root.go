/*
Copyright © 2025 Simen Kristiansen simenbolstad@gmail.com
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// getDataFilePath returns the configured data file path from Viper.
// It supports defaults, environment variables, and future config file support.
func getDataFilePath() string {
	return viper.GetString("datafile")
}

func init() {
	// Initialize Viper configuration
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// Fallback to current directory if home directory cannot be determined
		homeDir = "."
	}
	
	defaultDataFile := filepath.Join(homeDir, ".tridos.json")
	viper.SetDefault("datafile", defaultDataFile)
	
	// Bind to environment variable TRI_DATA_FILE
	viper.SetEnvPrefix("TRI")
	viper.BindEnv("datafile", "TRI_DATA_FILE")
	
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
