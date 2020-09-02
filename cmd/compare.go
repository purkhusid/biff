package cmd

import (
	"github.com/purkhusid/biff/compare"
	"github.com/spf13/cobra"
)

var (
	leftFlag              string
	rightFlag             string
	comparePutputPathFlag string
)

func init() {
	rootCmd.AddCommand(compareCommand)
	compareCommand.Flags().StringVar(&leftFlag, "left", "", "The 'left' side of the comparison (Required)")
	compareCommand.Flags().StringVar(&rightFlag, "right", "", "The 'right' side of the comparison (Required)")
	compareCommand.Flags().StringVar(&comparePutputPathFlag, "out", "", "Where the output should be written to (Required)")
	compareCommand.MarkFlagRequired("left")
	compareCommand.MarkFlagRequired("right")
	compareCommand.MarkFlagRequired("out")
}

var compareCommand = &cobra.Command{
	Use:   "compare",
	Short: "Compares two outputs from the calculate command",
	Long:  `Compares two outputs from the calculate command and outputs a file that can be fed to the '--target_pattern_file' option in bazel`,
	Run: func(cmd *cobra.Command, args []string) {
		compare.CompareAndWriteToOutput(leftFlag, rightFlag, comparePutputPathFlag)
	},
}
