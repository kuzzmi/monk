package main

import (
	gc "github.com/rthornton128/goncurses"
	"io/ioutil"
	"log"
)

type FilePanel struct {
	Height, Width, X, Y int
	Directory           string
	IsActive            bool
	Selected            int

	Panel *gc.Panel
}

func PrintHiglighted(window *gc.Window, y int, x int, index int, selected int, text string) {
	if index == selected {
		window.AttrOn(gc.A_STANDOUT)
	}
	window.MovePrint(y, x, text)
	if index == selected {
		window.AttrOff(gc.A_STANDOUT)
	}
}

func (fp *FilePanel) Draw() {
	window, _ := gc.NewWindow(fp.Height, fp.Width, fp.Y, fp.X)
	window.Box(0, 0)

	files, _ := ioutil.ReadDir(fp.Directory)

	PrintHiglighted(window, 1, 1, 0, fp.Selected, "..")
	for i, f := range files {
		PrintHiglighted(window, i+2, 1, i, fp.Selected-1, f.Name())
	}

	if fp.IsActive {
		var msg = " [ Active ] "
		window.MovePrint(0, (fp.Width-len(msg))/2, msg)
	}

	fp.Panel = gc.NewPanel(window)
}

func (fp *FilePanel) Redraw() {
	fp.Panel.Window().Erase()
	fp.Panel.Window().NoutRefresh()

	fp.Draw()

	gc.Update()
}

func (fp *FilePanel) Select(index int) {
	fp.Selected = index
	fp.Redraw()
}

func (fp *FilePanel) GoUp() {
	fp.Directory = fp.Directory + "../"
	fp.Select(0)
}

func (fp *FilePanel) ToggleActivity() {
	fp.IsActive = !fp.IsActive
	fp.Redraw()
}

func TogglePannels(panels [2]FilePanel, activePanel int) {
}

func main() {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)

	var panels [2]FilePanel
	rows, cols := stdscr.MaxYX()
	height, width := rows, cols/2
	y, x := 0, 0
	activePanel := 0

	panels[0] = FilePanel{Height: height, Width: width, Y: y, X: x, Directory: "./", Selected: 0, IsActive: true}
	panels[1] = FilePanel{Height: height, Width: width, Y: y, X: x + width, Directory: "/home/kuzzmi/", Selected: 2, IsActive: false}

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
			panels[activePanel].GoUp()
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
		case 'k':
			panels[activePanel].Select(panels[activePanel].Selected - 1)
		case 'j':
			panels[activePanel].Select(panels[activePanel].Selected + 1)
		}
	}
	stdscr.Delete()
}
