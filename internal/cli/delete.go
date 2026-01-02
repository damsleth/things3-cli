package cli

import (
	"github.com/ossianhempel/things3-cli/internal/things"
	"github.com/spf13/cobra"
)

// NewDeleteCommand builds the delete subcommand.
func NewDeleteCommand(app *App) *cobra.Command {
	opts := things.DeleteTodoOptions{}
	var confirm string

	cmd := &cobra.Command{
		Use:   "delete [OPTIONS...] [--] [-|TITLE]",
		Short: "Delete an existing todo",
		RunE: func(cmd *cobra.Command, args []string) error {
			rawInput, err := readInput(app.In, args)
			if err != nil {
				return err
			}

			target := deleteConfirmTarget(opts.ID, rawInput)
			if err := confirmDelete(app, "todo", target, confirm); err != nil {
				return err
			}

			script, err := things.BuildDeleteTodoScript(opts, rawInput)
			if err != nil {
				return err
			}
			return runScript(app, script)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.ID, "id", "", "ID of the todo to delete")
	flags.StringVar(&confirm, "confirm", "", "Confirm deletion by typing the todo ID or title")

	return cmd
}
