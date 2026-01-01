package cli

import (
	"fmt"

	"github.com/ossianhempel/things3-cli/internal/db"
)

func buildTaskFilter(store *db.Store, status string, includeTrashed bool, all bool, project string, area string, tag string, search string, limit int, includeChecklist bool) (db.TaskFilter, error) {
	statusFilter, err := db.ParseStatus(status)
	if err != nil {
		return db.TaskFilter{}, fmt.Errorf("Error: %s", err)
	}
	if all {
		statusFilter = nil
		includeTrashed = true
	}

	projectID := ""
	if project != "" {
		projectID, err = store.ResolveProjectID(project)
		if err != nil {
			return db.TaskFilter{}, fmt.Errorf("Error: %s", err)
		}
	}

	areaID := ""
	if area != "" {
		areaID, err = store.ResolveAreaID(area)
		if err != nil {
			return db.TaskFilter{}, fmt.Errorf("Error: %s", err)
		}
	}

	tagID := ""
	if tag != "" {
		tagID, err = store.ResolveTagID(tag)
		if err != nil {
			return db.TaskFilter{}, fmt.Errorf("Error: %s", err)
		}
	}

	return db.TaskFilter{
		Status:           statusFilter,
		IncludeTrashed:   includeTrashed,
		ExcludeTrashedContext: true,
		ProjectID:        projectID,
		AreaID:           areaID,
		TagID:            tagID,
		Search:           search,
		Limit:            limit,
		IncludeChecklist: includeChecklist,
	}, nil
}
