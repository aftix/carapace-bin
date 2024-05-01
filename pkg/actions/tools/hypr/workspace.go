package hypr

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
)

type hyprWorkspace struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Monitor         string `json:"monitor,omitempty"`
	MonitorID       int    `json:"monitorID,omitempty"`
	Windows         int    `json:"windows,omitempty"`
	HasFullscreen   bool   `json:"hasfullscreen,omitempty"`
	LastWindow      string `json:"lastwindow,omitempty"`
	LastWindowTitle string `json:"lastwindowtitle,omitempty"`
}

// ActionWorkspaces completes Hyprland workspace identifier syntax
//
// 1
// +1
// e-2
// m+1
// r-1
// name:web
// special:magic
func ActionWorkspaces() carapace.Action {
	return carapace.ActionExecCommand("hyprctl", "-j", "workspaces")(func(output []byte) carapace.Action {
		var workspaces []hyprWorkspace
		if err := json.Unmarshal(output, &workspaces); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		return carapace.Batch(
			ActionWorkspaceIDs(workspaces),
			ActionWorkspaceRelativeIDs(workspaces),
			ActionWorkspaceNames(workspaces),
			ActionSpecialWorkspaces(workspaces),
			carapace.ActionValuesDescribed(
				"previous", "Previously active workspace",
				"empty", "First available empty workspace",
			),
		).ToA()
	})
}

// ActionWorkspaceIDs completes Hyprland workspace IDs
func ActionWorkspaceIDs(workspaces []hyprWorkspace) carapace.Action {
	values := make([]string, 0)
	for _, workspace := range workspaces {
		values = append(values, fmt.Sprint(workspace.Id), workspace.Name)
	}
	return carapace.ActionValuesDescribed(values...)
}

// ActionWorkspaceRelativeIDs completes Hyprland workspace IDs relative to the currently active workspace
func ActionWorkspaceRelativeIDs(workspaces []hyprWorkspace) carapace.Action {
	return carapace.ActionExecCommand("hyprctl", "-j", "activeworkspace")(func(output []byte) carapace.Action {
		var activeWorkspace hyprWorkspace
		if err := json.Unmarshal(output, &activeWorkspace); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		values := []string{"+", "Relative workspace ID (including empty workspaces)", "-", "Relative workspace ID (including empty workspaces)"}
		for _, workspace := range workspaces {
			if relPos := workspace.Id - activeWorkspace.Id; relPos > 0 {
				values = append(values, fmt.Sprint("e+", relPos), workspace.Name)
			} else if relPos < 0 {
				values = append(values, fmt.Sprint("e", relPos), workspace.Name)
			}
		}
		return carapace.ActionValuesDescribed(values...)
	})
}

// ActionWorkspaceRelativeOnMonitorIDs completes Hyprland workspace IDs relative to the currently active workspace and monitor
func ActionWorkspaceRelativeOnMonitorIDs(workspaces []hyprWorkspace) carapace.Action {
	return carapace.ActionExecCommand("hyprctl", "-j", "activeworkspace")(func(output []byte) carapace.Action {
		var activeWorkspace hyprWorkspace
		if err := json.Unmarshal(output, &activeWorkspace); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		monitorWorkspaces := slices.Clone(workspaces)
		slices.DeleteFunc(monitorWorkspaces, func(w hyprWorkspace) bool {
			return w.MonitorID != activeWorkspace.MonitorID
		})
		slices.SortFunc(monitorWorkspaces, func(a, b hyprWorkspace) int {
			return a.Id - b.Id
		})
		activeIndex := slices.IndexFunc(monitorWorkspaces, func(w hyprWorkspace) bool {
			return w.Id == activeWorkspace.Id
		})

		values := []string{"r", "Relative workspace on monitor including empty workspaces"}
		for i, workspace := range monitorWorkspaces {
			if relPos := i - activeIndex; relPos > 0 {
				values = append(values, fmt.Sprint("m+", relPos), workspace.Name)
			} else if relPos < 0 {
				values = append(values, fmt.Sprint("m", relPos), workspace.Name)
			}
		}
		return carapace.ActionValuesDescribed(values...)
	})
}

// ActionWorkspaceNames completes Hyprland workspace names
func ActionWorkspaceNames(workspaces []hyprWorkspace) carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		if len(c.Parts) == 0 {
			return carapace.ActionValues("name")
		}

		names := make([]string, 0)
		for _, workspace := range workspaces {
			if workspace.Name != "" && !strings.HasPrefix(workspace.Name, "special") {
				names = append(names, workspace.Name)
			}
		}
		return carapace.ActionValues(names...)
	})
}

// ActionSpecialWorkspaces completes Hyprland special workspaces
func ActionSpecialWorkspaces(workspaces []hyprWorkspace) carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		if len(c.Parts) == 0 {
			return carapace.ActionValuesDescribed("special", "scratchpad workspace (optionally named after :)")
		}

		names := make([]string, 0)
		for _, workspace := range workspaces {
			if workspace.Name != "" && strings.HasPrefix(workspace.Name, "special:") {
				names = append(names, strings.TrimPrefix(workspace.Name, "special:"))
			}
		}
		return carapace.ActionValues(names...)
	})
}
