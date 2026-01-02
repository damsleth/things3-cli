package cli

import (
	"github.com/ossianhempel/things3-cli/internal/things"
	"github.com/spf13/cobra"
)

// NewUpdateCommand builds the update subcommand.
func NewUpdateCommand(app *App) *cobra.Command {
	opts := things.UpdateOptions{}

	cmd := &cobra.Command{
		Use:   "update [OPTIONS...] [--] [-|TITLE]",
		Short: "Update an existing todo",
		RunE: func(cmd *cobra.Command, args []string) error {
			rawInput, err := readInput(app.In, args)
			if err != nil {
				return err
			}

			if opts.AuthToken == "" {
				opts.AuthToken = authTokenFromEnv()
			}

			url, err := things.BuildUpdateURL(opts, rawInput)
			if err != nil {
				return err
			}
			return openURL(app, url)
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&opts.AuthToken, "auth-token", "", "Things URL scheme authorization token")
	flags.StringVar(&opts.ID, "id", "", "ID of the todo to update")
	flags.StringVar(&opts.Notes, "notes", "", "Replace notes")
	flags.StringVar(&opts.PrependNotes, "prepend-notes", "", "Prepend to notes")
	flags.StringVar(&opts.AppendNotes, "append-notes", "", "Append to notes")
	flags.StringVar(&opts.When, "when", "", "When to schedule the todo")
	flags.BoolVar(&opts.Later, "later", false, "Move the todo to Later")
	flags.StringVar(&opts.Deadline, "deadline", "", "Deadline for the todo")
	flags.StringVar(&opts.Tags, "tags", "", "Replace tags")
	flags.StringVar(&opts.AddTags, "add-tags", "", "Add tags")
	flags.BoolVar(&opts.Completed, "completed", false, "Mark the todo completed")
	flags.BoolVar(&opts.Canceled, "canceled", false, "Mark the todo canceled")
	flags.BoolVar(&opts.Canceled, "cancelled", false, "Mark the todo cancelled")
	flags.BoolVar(&opts.Reveal, "reveal", false, "Reveal the updated todo")
	flags.BoolVar(&opts.Duplicate, "duplicate", false, "Duplicate before updating")
	flags.StringVar(&opts.CompletionDate, "completion-date", "", "Completion date (ISO8601)")
	flags.StringVar(&opts.CreationDate, "creation-date", "", "Creation date (ISO8601)")
	flags.StringVar(&opts.Heading, "heading", "", "Heading within a project")
	flags.StringVar(&opts.List, "list", "", "Project or area to move to")
	flags.StringVar(&opts.ListID, "list-id", "", "Project or area ID to move to")
	flags.StringArrayVar(&opts.ChecklistItems, "checklist-item", nil, "Checklist item (repeatable)")
	flags.StringArrayVar(&opts.PrependChecklistItems, "prepend-checklist-item", nil, "Prepend checklist item (repeatable)")
	flags.StringArrayVar(&opts.AppendChecklistItems, "append-checklist-item", nil, "Append checklist item (repeatable)")

	return cmd
}
