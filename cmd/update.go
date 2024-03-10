package cmd

import (
	"fmt"

	"github.com/anowlbear/gitspan/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates all git repositories in the specified directory",
	Run: func(cmd *cobra.Command, args []string) {
		repos := util.GetRepos(args[0], -1)
		fmt.Println()
		for _, repo := range repos {
			fmt.Println(repo)
		}
	},
}
