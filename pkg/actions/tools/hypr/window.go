package hypr

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/carapace-sh/carapace"
)

type hyprClient struct {
	Address        string        `json:"address"`
	Mapped         bool          `json:"mapped"`
	Hidden         bool          `json:"hidden"`
	At             []int         `json:"at"`
	Size           []int         `json:"size"`
	Workspace      hyprWorkspace `json:"workspace"`
	Floating       bool          `json:"floating"`
	Monitor        int           `json:"monitor"`
	Class          string        `json:"class"`
	Title          string        `json:"title"`
	InitialClass   string        `json:"initialClass"`
	InitialTitle   string        `json:"initialTitle"`
	Pid            int           `json:"pid"`
	Xwayland       bool          `json:"xwayland"`
	Pinned         bool          `json:"pinned"`
	Fullscreen     bool          `json:"fullscreen"`
	FullscreenMode int           `json:"fullscreenMode"`
	FakeFullscreen bool          `json:"fakeFullscreen"`
	Grouped        []string      `json:"grouped"`
	Swallowing     string        `json:"swallowing"`
	FocusHistoryID int           `json:"focusHistoryID"`
}

// ActionWindows completes Hyprland window selector syntax
//
// class:kitty
// title:kitty,xwayland:1
func ActionWindows() carapace.Action {
	return carapace.ActionExecCommand("hyprctl", "-j", "clients")(func(output []byte) carapace.Action {
		var clients []hyprClient
		if err := json.Unmarshal(output, &clients); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				ActionClasses(clients),
				ActionTitles(clients),
				ActionInitialClasses(clients),
				ActionInitialTitles(clients),
				actionBool("xwayland"),
				actionBool("floating"),
				actionBool("fullscreen"),
				actionBool("pinned"),
				actionBool("focus"),
				carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
					if len(c.Parts) == 0 {
						return carapace.ActionValuesDescribed(
							"workspace", "match workspace id which window is on",
							"onworkspace", "match workspace id which window is on (supports selectors)",
						)
					}
					if c.Parts[0] != "workspace" && c.Parts[0] != "onworkspace" {
						return carapace.ActionValues()
					}
					return ActionWorkspaces()
				}),
			).ToA()
		})
	})
}

// ActionClasses completes window classes
func ActionClasses(clients []hyprClient) carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		if len(c.Parts) == 0 {
			return carapace.ActionValuesDescribed("class", "match window class against given regex")
		}
		if c.Parts[0] != "class" {
			return carapace.ActionValues()
		}

		classes := make([]string, 0)
		for _, client := range clients {
			if client.Class != "" {
				classes = append(classes, client.Class)
			}
		}
		slices.Sort(classes)
		slices.Compact(classes)
		return carapace.ActionValues(classes...)
	})
}

// ActionTitles completes window titles
func ActionTitles(clients []hyprClient) carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		if len(c.Parts) == 0 {
			return carapace.ActionValuesDescribed("title", "match window title against given regex")
		}
		if c.Parts[0] != "title" {
			return carapace.ActionValues()
		}

		titles := make([]string, 0)
		for _, client := range clients {
			if client.Title != "" {
				titles = append(titles, client.Title)
			}
		}
		slices.Sort(titles)
		slices.Compact(titles)
		return carapace.ActionValues(titles...)
	})
}

// ActionInitialClasses completes inital window classes
func ActionInitialClasses(clients []hyprClient) carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		if len(c.Parts) == 0 {
			return carapace.ActionValuesDescribed("initialClass", "match window inital class against given regex")
		}
		if c.Parts[0] != "initialClass" {
			return carapace.ActionValues()
		}

		initialClasses := make([]string, 0)
		for _, client := range clients {
			if client.InitialClass != "" {
				initialClasses = append(initialClasses, client.InitialClass)
			}
		}
		slices.Sort(initialClasses)
		slices.Compact(initialClasses)
		return carapace.ActionValues(initialClasses...)
	})
}

// ActionInitialTitles completes initial window titles
func ActionInitialTitles(clients []hyprClient) carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		if len(c.Parts) == 0 {
			return carapace.ActionValuesDescribed("initialTitle", "match window title against given regex")
		}
		if c.Parts[0] != "initialTitle" {
			return carapace.ActionValues()
		}

		initialTitles := make([]string, 0)
		for _, client := range clients {
			if client.InitialTitle != "" {
				initialTitles = append(initialTitles, client.InitialTitle)
			}
		}
		slices.Sort(initialTitles)
		slices.Compact(initialTitles)
		return carapace.ActionValues(initialTitles...)
	})
}

// Produce an action for matching a hyprland boolean window property
func actionBool(name string) carapace.Action {
	return carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
		if len(c.Parts) == 0 {
			return carapace.ActionValuesDescribed(name, fmt.Sprintf("match window %s mode", name))
		}
		if c.Parts[0] != "name" {
			return carapace.ActionValues()
		}

		return carapace.ActionValues("0", "1")
	})
}
