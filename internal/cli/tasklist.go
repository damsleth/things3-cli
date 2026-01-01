package cli

import (
	"github.com/ossianhempel/things3-cli/internal/db"
	"github.com/spf13/cobra"
)

type taskListRunner func(store *db.Store, filter db.TaskFilter) ([]db.Task, error)

func newTaskListCommand(app *App, use string, short string, defaultStatus string, runner taskListRunner) *cobra.Command {
	var dbPath string
	status := defaultStatus
	var project string
	var area string
	var tag string
	var includeTrashed bool
	var all bool
	var limit int
	var asJSON bool
	var noHeader bool
	var recursive bool

	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		RunE: func(cmd *cobra.Command, args []string) error {
			store, _, err := db.OpenDefault(dbPath)
			if err != nil {
				return formatDBError(err)
			}
			defer store.Close()

			filter, err := buildTaskFilter(store, status, includeTrashed, all, project, area, tag, "", limit, recursive)
			if err != nil {
				return err
			}

			tasks, err := runner(store, filter)
			if err != nil {
				return formatDBError(err)
			}
			return printTasks(app.Out, tasks, asJSON, noHeader)
		},
	}

	cmd.Flags().StringVarP(&dbPath, "db", "d", "", "Path to Things database (overrides THINGSDB)")
	cmd.Flags().StringVar(&dbPath, "database", "", "Alias for --db")
	cmd.Flags().StringVar(&status, "status", defaultStatus, "Filter by status: incomplete, completed, canceled, any")
	cmd.Flags().StringVarP(&project, "filter-project", "p", "", "Filter by project title or ID")
	cmd.Flags().StringVar(&project, "project", "", "Alias for --filter-project")
	cmd.Flags().StringVarP(&area, "filter-area", "a", "", "Filter by area title or ID")
	cmd.Flags().StringVar(&area, "area", "", "Alias for --filter-area")
	cmd.Flags().StringVarP(&tag, "filter-tag", "t", "", "Filter by tag title or ID")
	cmd.Flags().StringVar(&tag, "filtertag", "", "Alias for --filter-tag")
	cmd.Flags().StringVar(&tag, "tag", "", "Alias for --filter-tag")
	cmd.Flags().IntVar(&limit, "limit", 200, "Limit number of results (0 = no limit)")
	cmd.Flags().BoolVar(&includeTrashed, "include-trashed", false, "Include trashed tasks")
	cmd.Flags().BoolVar(&all, "all", false, "Include completed, canceled, and trashed tasks")
	cmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Include checklist items in JSON output")
	cmd.Flags().BoolVarP(&asJSON, "json", "j", false, "Output JSON")
	cmd.Flags().BoolVar(&noHeader, "no-header", false, "Suppress header row")

	return cmd
}
