package hypr

import (
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
)

type hyprMonitor struct {
	Id               int           `json:"id"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	Make             string        `json:"make"`
	Model            string        `json:"model"`
	Serial           string        `json:"serial"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	RefreshRate      float32       `json:"refreshRate"`
	X                int           `json:"x"`
	Y                int           `json:"y"`
	ActiveWorkspace  hyprWorkspace `json:"activeWorkspace"`
	SpecialWorkspace hyprWorkspace `json:"specialWorkspace"`
	Reserved         []int         `json:"reserved"`
	Scale            float32       `json:"scale"`
	Transform        int           `json:"transform"`
	Focused          bool          `json:"focused"`
	DpmsStatus       bool          `json:"dpmsStatus"`
	Vrr              bool          `json:"vrr"`
	ActivelyTearing  bool          `json:"activelyTearing"`
	Disabled         bool          `json:"disabled"`
	CurrentFormat    string        `json:"currentFormat"`
	AvailableModes   []string      `json:"availableModes"`
}

// ActionMonitors completes Hyprland monitor identifier sytax
//
// 1
// left
// DP-1
// desc:ViewSonic Corporation VX2703 SERIES T8G132800478
// current
// +1
func ActionMonitors() carapace.Action {
	return carapace.ActionExecCommand("hyprctl", "-j", "monitors")(func(output []byte) carapace.Action {
		var monitors []hyprMonitor
		if err := json.Unmarshal(output, &monitors); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		return carapace.Batch(
			ActionMonitorIDs(monitors),
			ActionMonitorRelativeIDs(monitors),
			ActionMonitorNames(monitors),
			ActionMonitorDescriptions(monitors),
			carapace.ActionValues("left", "right", "up", "down", "current"),
		).ToA()
	})
}

// ActionMonitorIDs completes Hyprland monitor IDs
func ActionMonitorIDs(monitors []hyprMonitor) carapace.Action {
	values := make([]string, 0)
	for _, monitor := range monitors {
		values = append(values, fmt.Sprint(monitor.Id), fmt.Sprintf("(%s) %s", monitor.Name, monitor.Description))
	}
	return carapace.ActionValuesDescribed(values...)
}

// ActionMonitorRelativeIDs completes Hyprland monitor IDs relative to the currently active monitor
func ActionMonitorRelativeIDs(monitors []hyprMonitor) carapace.Action {
	return carapace.ActionExecCommand("hyprctl", "-j", "activeworkspace")(func(output []byte) carapace.Action {
		var activeWorkspace hyprWorkspace
		if err := json.Unmarshal(output, &activeWorkspace); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		values := make([]string, 0)
		for _, monitor := range monitors {
			if relPos := monitor.Id - activeWorkspace.MonitorID; relPos > 0 {
				values = append(values, fmt.Sprint("+", relPos), fmt.Sprintf("(%s) %s", monitor.Name, monitor.Description))
			} else if relPos < 0 {
				values = append(values, fmt.Sprint(relPos), fmt.Sprintf("(%s) %s", monitor.Name, monitor.Description))
			}
		}
		return carapace.ActionValuesDescribed(values...)
	})
}

// ActionMonitorNames completes Hyprland monitor Names
func ActionMonitorNames(monitors []hyprMonitor) carapace.Action {
	values := make([]string, 0)
	for _, monitor := range monitors {
		values = append(values, monitor.Name, monitor.Description)
	}
	return carapace.ActionValuesDescribed(values...)
}

// ActionMonitorDescriptions completes Hyprland monitors by description
func ActionMonitorDescriptions(monitors []hyprMonitor) carapace.Action {
	values := make([]string, 0)
	for _, monitor := range monitors {
		values = append(values, monitor.Description, monitor.Name)
	}
	return carapace.ActionValuesDescribed(values...).Prefix("desc:")
}
