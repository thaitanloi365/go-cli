package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
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

type buildInfo struct {
	Env        string `json:"env"`
	GitCommit  string `json:"git_commit"`
	GitBranch  string `json:"git_branch"`
	GitState   string `json:"git_state"`
	GitSummary string `json:"git_summary"`
	BuildDate  string `json:"build_date"`
	Version    string `json:"version"`
}

func (i *buildInfo) print() {
	data, err := jsoniter.MarshalIndent(&i, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func showInfo() {
	var i = buildInfo{
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

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "info",
	Short: "App Infomartion",
	Run: func(cmd *cobra.Command, args []string) {
		showInfo()
	},
}
