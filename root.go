package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "EZIE Logistics TPL",
	Short: "EZIE Logistics TPL APIs",
}

var cfgFile string
var pwd string

func execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
		os.Exit(1)
	}
}

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pwd = dir

	cobra.OnInitialize(setup)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is %s/.env)", dir))

}

func setup() {
}
