package cmd

import (
	"path"

	"github.com/purkhusid/bazel_vcs_differ/calculate"
	"github.com/spf13/cobra"
)

var (
	query                   string
	bazelFlag               string
	workspacePathFlag       string
	calculateOutputPathFlag string
)

func init() {
	rootCmd.AddCommand(calculateCmd)
	calculateCmd.Flags().StringVar(&query, "query", "//external:all-targets + deps(//...:all-targets)", "The query used to query the bazel target graph")
	calculateCmd.Flags().StringVar(&bazelFlag, "bazel", "", "Location of Bazel executable. By default uses bazel from path")
	calculateCmd.Flags().StringVar(&workspacePathFlag, "workspace", "", "Path to the workspace root (Required)")
	calculateCmd.Flags().StringVar(&calculateOutputPathFlag, "out", "", "Where the output should be written to (Required)")
	calculateCmd.MarkFlagRequired("workspace")
	calculateCmd.MarkFlagRequired("out")
}

var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculate the checksum for each target in the Bazel graph.",
	Long: `Calculates the sha256 checksum for each target in the graph and outputs it to a file. 
	       This file can then be diffed against another output with the compare command`,
	Run: func(cmd *cobra.Command, args []string) {
		bazel := bazelFlag
		workspacePath := workspacePathFlag
		calculateOutputPathFlag := calculateOutputPathFlag

		queryResult := calculate.GetQueryResult(workspacePath, bazel, query)
		calculator := calculate.NewCalculator(queryResult)
		hashedTargets := calculator.CalculateHashes()

		calculate.WriteResultsToFile(hashedTargets, path.Join(workspacePath, calculateOutputPathFlag))
	},
}
