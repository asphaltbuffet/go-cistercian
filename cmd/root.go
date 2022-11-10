// Package cmd contains the CLI commands for go-cistercian.
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/asphaltbuffet/go-cistercian/pkg/cistercian"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := &cobra.Command{
		Use:   "go-cistercian [flags]",
		Short: "go-cistercian is a CLI tool for generating Cistercian numbers",
		Long:  `go-cistercian is a CLI tool for generating Cistercian numbers.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   RunRootCmd,
	}

	rootCmd.Flags().Bool("svg", false, "output SVG")

	cobra.CheckErr(rootCmd.Execute())
}

// RunRootCmd is the entry point for the CLI.
func RunRootCmd(cmd *cobra.Command, args []string) {
	num, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}

	isSvg, err := cmd.Flags().GetBool("svg")
	if err != nil {
		return
	}

	cipher, err := cistercian.GenerateCistercianNumber(num, isSvg)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	cmd.Println(cipher)
}
