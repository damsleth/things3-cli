package cli

import (
	"github.com/ossianhempel/things3-cli/internal/things"
	"github.com/spf13/cobra"
)

// NewDeleteAreaCommand builds the delete-area subcommand.
func NewDeleteAreaCommand(app *App) *cobra.Command {
	opts := things.DeleteAreaOptions{}

	cmd := &cobra.Command{
		Use:   "delete-area [OPTIONS...] [--] [-|TITLE]",
		Short: "Delete an existing area",
		RunE: func(cmd *cobra.Command, args []string) error {
			rawInput, err := readInput(app.In, args)
			if err != nil {
				return err
			}

			script, err := things.BuildDeleteAreaScript(opts, rawInput)
			if err != nil {
				return err
			}
			return runScript(app, script)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.ID, "id", "", "ID of the area to delete")

	return cmd
}
