package cmd

import (
	"path"

	"github.com/purkhusid/bazel_vcs_differ/calculate"
	"github.com/spf13/cobra"
)

var (
	bazelFlag         string
	workspacePathFlag string
	outputPathFlag    string
)

func init() {
	rootCmd.AddCommand(calculateCmd)
	calculateCmd.Flags().StringVar(&bazelFlag, "bazel", "", "Location of Bazel executable. By default uses bazel from path")
	calculateCmd.Flags().StringVar(&workspacePathFlag, "workspace", "", "Path to the workspace root (Required)")
	calculateCmd.Flags().StringVar(&outputPathFlag, "out", "", "Where the output should be written to (Required)")
	calculateCmd.MarkFlagRequired("workspace")
	calculateCmd.MarkFlagRequired("out")
}

var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Calculate the checksum for each target in the Bazel graph.",
	Long: `Calculates the sha256 checksum for each target in the graph and outputs it to a file. 
	       This file can then be diffed against another outputu with the diff command`,
	Run: func(cmd *cobra.Command, args []string) {
		bazel := bazelFlag
		workspacePath := workspacePathFlag
		outputPath := outputPathFlag

		queryResult := calculate.GetQueryResult(workspacePath, bazel)
		calculator := calculate.NewCalculator(queryResult)
		hashedTargets := calculator.CalculateHashes()

		calculate.WriteResultsToFile(hashedTargets, path.Join(workspacePath, outputPath))
	},
}
