package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Build info
var (
	Env        = ""
	GitCommit  = ""
	GitBranch  = ""
	GitState   = ""
	GitSummary = ""
	BuildDate  = ""
	Version    = ""
)

type info struct {
	Env        string `json:"env"`
	GitCommit  string `json:"git_commit"`
	GitBranch  string `json:"git_branch"`
	GitState   string `json:"git_state"`
	GitSummary string `json:"git_summary"`
	BuildDate  string `json:"build_date"`
	Version    string `json:"version"`
}

func (i *info) print() {
	data, err := json.MarshalIndent(&i, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}


// RootCmd root command
var RootCmd = &cobra.Command{
	Use:   "EZIE Logistics TPL",
	Short: "EZIE Logistics TPL APIs",
}

var cfgFile string
var pwd string

// Execute Execute root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
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

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is %s/.env)", dir))

}

func setup()  {
  info()
}

func info()  {
  var i = info{
		Env:        Env,
		GitCommit:  GitCommit,
		GitBranch:  GitBranch,
		GitState:   GitState,
		GitSummary: GitSummary,
		BuildDate:  BuildDate,
		Version:    Version,
  }

  i.print()
}
