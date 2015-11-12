package main

import (
	fp "github.com/kuzzmi/monk/filepanel"
	gc "github.com/rthornton128/goncurses"
	"log"
)

func main() {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)

	var panels [2]fp.FilePanel
	rows, cols := stdscr.MaxYX()
	height, width := rows, cols/2
	y, x := 0, 0
	activePanel := 0

	panels[0] = fp.FilePanel{Height: height, Width: width, Y: y, X: x, Directory: "./", Selected: 0, IsActive: true}
	panels[1] = fp.FilePanel{Height: height, Width: width, Y: y, X: x + width, Directory: "/home/kuzzmi/", Selected: 2, IsActive: false}

	panels[0].Draw()
	panels[1].Draw()

	gc.UpdatePanels()
	gc.Update()

	stdscr.Keypad(true)
	// stdscr.GetChar()

main:
	for {
		switch panels[activePanel].Panel.Window().GetChar() {
		case 'q':
			break main
		case gc.KEY_RETURN:
			panels[activePanel].Execute()
		case gc.KEY_TAB:
			panels[0].ToggleActivity()
			panels[1].ToggleActivity()
			gc.UpdatePanels()
			gc.Update()
			if activePanel == 0 {
				activePanel = 1
			} else {
				activePanel = 0
			}
		case 'u':
			panels[activePanel].GoUp()
		case 'k':
			panels[activePanel].Select(panels[activePanel].Selected - 1)
		case 'j':
			panels[activePanel].Select(panels[activePanel].Selected + 1)
		case 'i':
			panels[activePanel].HideHidden()
		}
	}
	stdscr.Delete()
}
