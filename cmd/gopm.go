package cmd

import (
	"Tim4334/gopm/cmd/vars"
	b "Tim4334/gopm/internal/build"
	i "Tim4334/gopm/internal/init"
	l "Tim4334/gopm/internal/lint"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	GopmCmd = &cobra.Command{
		Use:   vars.GOPM,
		Short: vars.GOPM_SHORT,
		Long:  vars.GOPM_LONG,
	}
	InitCmd = &cobra.Command{
		Use:   vars.INIT,
		Short: vars.INIT_SHORT,
		Long:  vars.INIT_LONG,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := i.Init(cmd, args)
			if err != nil {
				return err
			}
			return nil
		},
	}
	BuildCmd = &cobra.Command{
		Use:   vars.BUILD,
		Short: vars.BUILD_SHORT,
		Long:  vars.BUILD_LONG,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := b.Build(cmd, args)
			if err != nil {
				return err
			}
			return nil
		},
	}
	LintCmd = &cobra.Command{
		Use:   vars.LINT,
		Short: vars.LINT_SHORT,
		Long:  vars.LINT_LONG,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := l.Lint(cmd, args)
			if err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	GopmCmd.AddCommand(InitCmd, BuildCmd, LintCmd)
}

func Execute() {
	if err := GopmCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
