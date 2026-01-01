package cli

import (
	"github.com/ossianhempel/things3-cli/internal/things"
	"github.com/spf13/cobra"
)

// NewAddCommand builds the add subcommand.
func NewAddCommand(app *App) *cobra.Command {
	opts := things.AddOptions{}

	cmd := &cobra.Command{
		Use:   "add [OPTIONS...] [--] [-|TITLE]",
		Short: "Add a new todo",
		RunE: func(cmd *cobra.Command, args []string) error {
			rawInput, err := readInput(app.In, args)
			if err != nil {
				return err
			}

			url := things.BuildAddURL(opts, rawInput)
			return openURL(app, url)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.When, "when", "", "When to schedule the todo")
	flags.StringVar(&opts.Deadline, "deadline", "", "Deadline for the todo")
	flags.BoolVar(&opts.Completed, "completed", false, "Mark the todo completed")
	flags.BoolVar(&opts.Canceled, "canceled", false, "Mark the todo canceled")
	flags.BoolVar(&opts.Canceled, "cancelled", false, "Mark the todo cancelled")
	flags.StringArrayVar(&opts.ChecklistItems, "checklist-item", nil, "Checklist item (repeatable)")
	flags.StringVar(&opts.CreationDate, "creation-date", "", "Creation date (ISO8601)")
	flags.StringVar(&opts.CompletionDate, "completion-date", "", "Completion date (ISO8601)")
	flags.StringVar(&opts.List, "list", "", "Project or area to add to")
	flags.StringVar(&opts.ListID, "list-id", "", "Project or area ID to add to")
	flags.StringVar(&opts.Heading, "heading", "", "Heading within a project")
	flags.BoolVar(&opts.Reveal, "reveal", false, "Reveal the newly created todo")
	flags.BoolVar(&opts.ShowQuickEntry, "show-quick-entry", false, "Show the quick entry dialog")
	flags.StringVar(&opts.Notes, "notes", "", "Notes for the todo")
	flags.StringVar(&opts.Tags, "tags", "", "Comma-separated tags")
	flags.StringVar(&opts.TitlesRaw, "titles", "", "Comma-separated titles for multiple todos")
	flags.StringVar(&opts.UseClipboard, "use-clipboard", "", "Use clipboard content")

	return cmd
}
