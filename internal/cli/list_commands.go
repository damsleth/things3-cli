package cli

import (
	"github.com/ossianhempel/things3-cli/internal/db"
	"github.com/spf13/cobra"
)

func NewInboxCommand(app *App) *cobra.Command {
	return newTaskListCommand(app, "inbox", "List inbox tasks from the Things database", "incomplete", func(store *db.Store, filter db.TaskFilter) ([]db.Task, error) {
		return store.InboxTasks(filter)
	})
}

func NewAnytimeCommand(app *App) *cobra.Command {
	return newTaskListCommand(app, "anytime", "List Anytime tasks from the Things database", "incomplete", func(store *db.Store, filter db.TaskFilter) ([]db.Task, error) {
		return store.AnytimeTasks(filter)
	})
}

func NewSomedayCommand(app *App) *cobra.Command {
	return newTaskListCommand(app, "someday", "List Someday tasks from the Things database", "incomplete", func(store *db.Store, filter db.TaskFilter) ([]db.Task, error) {
		return store.SomedayTasks(filter)
	})
}

func NewUpcomingCommand(app *App) *cobra.Command {
	return newTaskListCommand(app, "upcoming", "List upcoming tasks from the Things database", "incomplete", func(store *db.Store, filter db.TaskFilter) ([]db.Task, error) {
		return store.UpcomingTasks(filter)
	})
}

func NewDeadlinesCommand(app *App) *cobra.Command {
	return newTaskListCommand(app, "deadlines", "List tasks with deadlines from the Things database", "incomplete", func(store *db.Store, filter db.TaskFilter) ([]db.Task, error) {
		return store.DeadlinesTasks(filter)
	})
}

func NewLogbookCommand(app *App) *cobra.Command {
	return newTaskListCommand(app, "logbook", "List logbook tasks from the Things database", "any", func(store *db.Store, filter db.TaskFilter) ([]db.Task, error) {
		return store.LogbookTasks(filter)
	})
}

func NewCompletedCommand(app *App) *cobra.Command {
	return newTaskListCommand(app, "completed", "List completed tasks from the Things database", "completed", func(store *db.Store, filter db.TaskFilter) ([]db.Task, error) {
		return store.CompletedTasks(filter)
	})
}

func NewCanceledCommand(app *App) *cobra.Command {
	return newTaskListCommand(app, "canceled", "List canceled tasks from the Things database", "canceled", func(store *db.Store, filter db.TaskFilter) ([]db.Task, error) {
		return store.CanceledTasks(filter)
	})
}

func NewTrashCommand(app *App) *cobra.Command {
	return newTaskListCommand(app, "trash", "List trashed tasks from the Things database", "any", func(store *db.Store, filter db.TaskFilter) ([]db.Task, error) {
		return store.TrashTasks(filter)
	})
}
