package main

import (
	"encoding/json"
	"fmt"

	"github.com/reconquest/karma-go"
	"github.com/reconquest/pkg/log"

	"github.com/kovetskiy/i3ipc"
)

type WindowFocusEvent struct {
	Container struct {
		ID                 int64   `json:"id"`
		Type               string  `json:"type"`
		Orientation        string  `json:"orientation"`
		ScratchpadState    string  `json:"scratchpad_state"`
		Percent            float64 `json:"percent"`
		Urgent             bool    `json:"urgent"`
		Focused            bool    `json:"focused"`
		Output             string  `json:"output"`
		Layout             string  `json:"layout"`
		WorkspaceLayout    string  `json:"workspace_layout"`
		LastSplitLayout    string  `json:"last_split_layout"`
		Border             string  `json:"border"`
		CurrentBorderWidth int     `json:"current_border_width"`
		Rect               struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"rect"`
		DecoRect struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"deco_rect"`
		WindowRect struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"window_rect"`
		Geometry struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"geometry"`
		Name             string `json:"name"`
		Window           int    `json:"window"`
		WindowType       string `json:"window_type"`
		WindowProperties struct {
			Class        string      `json:"class"`
			Instance     string      `json:"instance"`
			Title        string      `json:"title"`
			TransientFor interface{} `json:"transient_for"`
		} `json:"window_properties"`
		Nodes          []interface{} `json:"nodes"`
		FloatingNodes  []interface{} `json:"floating_nodes"`
		Focus          []interface{} `json:"focus"`
		FullscreenMode int           `json:"fullscreen_mode"`
		Sticky         bool          `json:"sticky"`
		Floating       string        `json:"floating"`
		Swallows       []interface{} `json:"swallows"`
	} `json:"container"`
}

type WorkspaceFocusEvent struct {
	Current struct {
		ID                 int64   `json:"id"`
		Type               string  `json:"type"`
		Orientation        string  `json:"orientation"`
		ScratchpadState    string  `json:"scratchpad_state"`
		Percent            float64 `json:"percent"`
		Urgent             bool    `json:"urgent"`
		Focused            bool    `json:"focused"`
		Output             string  `json:"output"`
		Layout             string  `json:"layout"`
		WorkspaceLayout    string  `json:"workspace_layout"`
		LastSplitLayout    string  `json:"last_split_layout"`
		Border             string  `json:"border"`
		CurrentBorderWidth int     `json:"current_border_width"`
		Rect               struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"rect"`
		DecoRect struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"deco_rect"`
		WindowRect struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"window_rect"`
		Geometry struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"geometry"`
		Name string `json:"name"`
		Num  int    `json:"num"`
		Gaps struct {
			Inner  int `json:"inner"`
			Outer  int `json:"outer"`
			Top    int `json:"top"`
			Right  int `json:"right"`
			Bottom int `json:"bottom"`
			Left   int `json:"left"`
		} `json:"gaps"`
		Window     interface{} `json:"window"`
		WindowType interface{} `json:"window_type"`
		Nodes      []struct {
			ID                 int64   `json:"id"`
			Type               string  `json:"type"`
			Orientation        string  `json:"orientation"`
			ScratchpadState    string  `json:"scratchpad_state"`
			Percent            float64 `json:"percent"`
			Urgent             bool    `json:"urgent"`
			Focused            bool    `json:"focused"`
			Output             string  `json:"output"`
			Layout             string  `json:"layout"`
			WorkspaceLayout    string  `json:"workspace_layout"`
			LastSplitLayout    string  `json:"last_split_layout"`
			Border             string  `json:"border"`
			CurrentBorderWidth int     `json:"current_border_width"`
			Rect               struct {
				X      int `json:"x"`
				Y      int `json:"y"`
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"rect"`
			DecoRect struct {
				X      int `json:"x"`
				Y      int `json:"y"`
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"deco_rect"`
			WindowRect struct {
				X      int `json:"x"`
				Y      int `json:"y"`
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"window_rect"`
			Geometry struct {
				X      int `json:"x"`
				Y      int `json:"y"`
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"geometry"`
			Name       interface{} `json:"name"`
			Window     interface{} `json:"window"`
			WindowType interface{} `json:"window_type"`
			Nodes      []struct {
				ID                 int64   `json:"id"`
				Type               string  `json:"type"`
				Orientation        string  `json:"orientation"`
				ScratchpadState    string  `json:"scratchpad_state"`
				Percent            float64 `json:"percent"`
				Urgent             bool    `json:"urgent"`
				Focused            bool    `json:"focused"`
				Output             string  `json:"output"`
				Layout             string  `json:"layout"`
				WorkspaceLayout    string  `json:"workspace_layout"`
				LastSplitLayout    string  `json:"last_split_layout"`
				Border             string  `json:"border"`
				CurrentBorderWidth int     `json:"current_border_width"`
				Rect               struct {
					X      int `json:"x"`
					Y      int `json:"y"`
					Width  int `json:"width"`
					Height int `json:"height"`
				} `json:"rect"`
				DecoRect struct {
					X      int `json:"x"`
					Y      int `json:"y"`
					Width  int `json:"width"`
					Height int `json:"height"`
				} `json:"deco_rect"`
				WindowRect struct {
					X      int `json:"x"`
					Y      int `json:"y"`
					Width  int `json:"width"`
					Height int `json:"height"`
				} `json:"window_rect"`
				Geometry struct {
					X      int `json:"x"`
					Y      int `json:"y"`
					Width  int `json:"width"`
					Height int `json:"height"`
				} `json:"geometry"`
				Name             string `json:"name"`
				Window           int    `json:"window"`
				WindowType       string `json:"window_type"`
				WindowProperties struct {
					Class        string      `json:"class"`
					Instance     string      `json:"instance"`
					Title        string      `json:"title"`
					TransientFor interface{} `json:"transient_for"`
				} `json:"window_properties,omitempty"`
				Nodes          []interface{} `json:"nodes"`
				FloatingNodes  []interface{} `json:"floating_nodes"`
				Focus          []interface{} `json:"focus"`
				FullscreenMode int           `json:"fullscreen_mode"`
				Sticky         bool          `json:"sticky"`
				Floating       string        `json:"floating"`
				Swallows       []interface{} `json:"swallows"`
			} `json:"nodes"`
			FloatingNodes  []interface{} `json:"floating_nodes"`
			Focus          []int64       `json:"focus"`
			FullscreenMode int           `json:"fullscreen_mode"`
			Sticky         bool          `json:"sticky"`
			Floating       string        `json:"floating"`
			Swallows       []interface{} `json:"swallows"`
		} `json:"nodes"`
		FloatingNodes  []interface{} `json:"floating_nodes"`
		Focus          []int64       `json:"focus"`
		FullscreenMode int           `json:"fullscreen_mode"`
		Sticky         bool          `json:"sticky"`
		Floating       string        `json:"floating"`
		Swallows       []interface{} `json:"swallows"`
	} `json:"current"`
	Old struct {
		ID                 int64   `json:"id"`
		Type               string  `json:"type"`
		Orientation        string  `json:"orientation"`
		ScratchpadState    string  `json:"scratchpad_state"`
		Percent            float64 `json:"percent"`
		Urgent             bool    `json:"urgent"`
		Focused            bool    `json:"focused"`
		Output             string  `json:"output"`
		Layout             string  `json:"layout"`
		WorkspaceLayout    string  `json:"workspace_layout"`
		LastSplitLayout    string  `json:"last_split_layout"`
		Border             string  `json:"border"`
		CurrentBorderWidth int     `json:"current_border_width"`
		Rect               struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"rect"`
		DecoRect struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"deco_rect"`
		WindowRect struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"window_rect"`
		Geometry struct {
			X      int `json:"x"`
			Y      int `json:"y"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"geometry"`
		Name string `json:"name"`
		Num  int    `json:"num"`
		Gaps struct {
			Inner  int `json:"inner"`
			Outer  int `json:"outer"`
			Top    int `json:"top"`
			Right  int `json:"right"`
			Bottom int `json:"bottom"`
			Left   int `json:"left"`
		} `json:"gaps"`
		Window         interface{}   `json:"window"`
		WindowType     interface{}   `json:"window_type"`
		Nodes          []interface{} `json:"nodes"`
		FloatingNodes  []interface{} `json:"floating_nodes"`
		Focus          []interface{} `json:"focus"`
		FullscreenMode int           `json:"fullscreen_mode"`
		Sticky         bool          `json:"sticky"`
		Floating       string        `json:"floating"`
		Swallows       []interface{} `json:"swallows"`
	} `json:"old"`
}

var ghostsMap = map[string][]int{}

func main() {
	i3, err := i3ipc.GetIPCSocket()
	if err != nil {
		log.Fatalf(err, "unable to get ipc socket")
	}

	defer i3.Close()

	windows, err := i3ipc.Subscribe(i3ipc.I3WindowEvent)
	if err != nil {
		log.Fatalf(err, "unable to subscribe for i3 window events")
	}

	for {
		select {
		case windowEvent := <-windows:
			if windowEvent.Change != "focus" {
				continue
			}

			var event WindowFocusEvent
			err := json.Unmarshal(windowEvent.Payload, &event)
			if err != nil {
				log.Fatalf(err, "unable to unmarshal window focus event")
			}

			root, err := i3.GetTree()
			if err != nil {
				log.Fatalf(err, "unable to get i3 tree")
			}

			toggle(i3, root, event)
		}
	}
}

func match(node i3ipc.I3Node, win int) bool {
	for _, sub := range node.Nodes {
		if int(sub.Window) == win {
			return true
		}

		if match(sub, win) {
			return true
		}
	}

	return false
}

func count(node i3ipc.I3Node) (int, int) {
	var windows int
	var ghosts int
	for _, sub := range node.Nodes {
		if sub.Window != 0 {
			windows++
		}

		if sub.Window == 0 && len(sub.Nodes) == 0 {
			ghosts++
		}

		if len(sub.Nodes) != 0 {
			subWindows, subGhosts := count(sub)

			windows += subWindows
			ghosts += subGhosts
		}
	}

	return windows, ghosts
}

func toggle(i3 *i3ipc.IPCSocket, root i3ipc.I3Node, event WindowFocusEvent) error {
	const BLOCK_CONTENT = "content"
	const WORKSPACE_SCRATCHPAD = "__i3_scratch"

	var target i3ipc.I3Node
	var found bool

loop:
	for _, output := range root.Nodes {
		for _, block := range output.Nodes {
			if block.Name == BLOCK_CONTENT {
				for _, workspace := range block.Nodes {
					if workspace.Name == WORKSPACE_SCRATCHPAD {
						continue
					}

					if match(workspace, event.Container.Window) {
						found = true
						target = workspace
						break loop
					}
				}
			}
		}
	}

	if !found {
		return nil
	}

	windows, ghosts := count(target)

	if windows == 1 && ghosts == 1 || windows > 1 {
		cached, ok := ghostsMap[target.Name]
		if !ok {
			return nil
		}

		for _, id := range cached {
			if int64(id) == event.Container.ID {
				continue
			}

			_, err := i3.Command(fmt.Sprintf("[con_id=%d] kill", id))
			if err != nil {
				return karma.Format(
					err,
					"unable to kill ghost window",
				)
			}
		}

		delete(ghostsMap, target.Name)

		ghosts = 0
	}

	if windows == 1 && ghosts == 0 {
		left, err := newGhost(i3)
		if err != nil {
			return err
		}

		err = moveGhost(i3, left, "left")
		if err != nil {
			return err
		}

		right, err := newGhost(i3)
		if err != nil {
			return err
		}

		for i := 0; i < 3; i++ {
			err = moveGhost(i3, right, "right")
			if err != nil {
				return err
			}
		}

		_, err = i3.Command(fmt.Sprintf("[id=%d] focus", event.Container.Window))
		if err != nil {
			return karma.Format(
				err,
				"unable to focus window",
			)
		}

		_, err = i3.Command(fmt.Sprintf("[id=%d] resize set 60 ppt", event.Container.Window))
		if err != nil {
			return karma.Format(
				err,
				"unable to resize window",
			)
		}

		ghostsMap[target.Name] = []int{left, right}
		return nil
	}

	return nil
}

func moveGhost(i3 *i3ipc.IPCSocket, id int, direction string) error {
	cmd := fmt.Sprintf("[con_id=%d] move %s", id, direction)

	reply, err := i3.Command(cmd)
	if err != nil {
		return karma.Format(
			err,
			"unable to move ghost window",
		)
	}

	if reply == false {
		return fmt.Errorf("command failed: %s", cmd)
	}

	return nil
}

func newGhost(i3 *i3ipc.IPCSocket) (int, error) {
	reply, err := i3.Raw(i3ipc.I3Command, "open")
	if err != nil {
		return 0, karma.Format(
			err,
			"unable to create a ghost window",
		)
	}

	var items []struct {
		ID int `json:"id"`
	}

	err = json.Unmarshal(reply, &items)
	if err != nil {
		return 0, karma.Format(
			err,
			"unable to unmarshal json reply: %s", string(reply),
		)
	}

	if len(items) == 0 {
		return 0, karma.Format(
			err,
			"invalid reply from 'open' command because no items found: %s",
			string(reply),
		)
	}

	return items[0].ID, nil
}
