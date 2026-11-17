package cmd

import (
	"Tim4334/gopm/cmd/env"

	"github.com/spf13/cobra"
)

var (
	BuildCmd = &cobra.Command{
		Use: env.BUILD,
		Short: env.BUILD_SHORT,
		Long: env.BUILD_LONG,
		RunE: func(cmd *cobra.Command, args []string) error {
			
		},
	}
)