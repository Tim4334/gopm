package cmd

import (
	"Tim4334/gopm/cmd/env"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	GopmCmd = &cobra.Command{
		Use: env.GOPM,
		Short: env.GOPM_SHORT,
		Long: env.GOPM_LONG,
	}
)

func init() {
	GopmCmd.AddCommand(InitCmd,BuildCmd,LintCmd)
}

func Execute() {
	if err := GopmCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}