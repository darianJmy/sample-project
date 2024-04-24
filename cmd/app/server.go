package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"sample-project/api/server/router"
	"sample-project/cmd/app/options"
)

func NewSampleServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()

	cmd := &cobra.Command{
		Use: "sample-server",
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := s.Complete(); err != nil {
				return err
			}

			if err := s.Registry(); err != nil {
				return err
			}

			return Run(s)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	s.BindFlags(cmd)

	return cmd
}

func Run(opts *options.ServerRunOptions) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", opts.ComponentConfig.Default.Listen),
		Handler: opts.HttpEngine,
	}

	router.InstallRouters(opts)

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
