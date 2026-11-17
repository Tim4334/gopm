package cmd

import (
	"Tim4334/gopm/cmd/env"

	"github.com/spf13/cobra"
)

var (
	LintCmd = &cobra.Command{
		Use: env.LINT,
		Short: env.LINT_SHORT,
		Long: env.LINT_LONG,
		RunE: func(cmd *cobra.Command, args []string) error {
			
		},
	}
)