package cmd

import (
	"Tim4334/gopm/cmd/env"

	"github.com/spf13/cobra"
)

var (
	InitCmd = &cobra.Command{
		Use: env.INIT,
		Short: env.INIT_SHORT,
		Long: env.INIT_LONG,
		RunE: func(cmd *cobra.Command, args []string) error {
			
		},
	}
)