package filepanel

import (
	gc "github.com/rthornton128/goncurses"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type FilePanel struct {
	Height, Width, X, Y int
	Directory           string
	IsActive            bool
	Selected            int
	Files               []os.FileInfo

	Panel *gc.Panel
}

func PrintHiglighted(window *gc.Window, y int, x int, index int, selected int, text string, isActive bool) {
	if index == selected && isActive {
		window.AttrOn(gc.A_STANDOUT)
	}
	window.MovePrint(y, x, text)
	if index == selected && isActive {
		window.AttrOff(gc.A_STANDOUT)
	}
}

func RowOutput(name string, size int64, isDir bool, width int) string {
	var sizeString = strconv.FormatInt(size, 10) + " "
	if isDir {
		sizeString = ""
		name += "/"
	}
	var spacesBetween = width - 1 - len(name) - len(sizeString)
	var spaces = make([]string, spacesBetween)
	return name + strings.Join(spaces, " ") + sizeString
}

func (fp *FilePanel) Draw() {
	window, _ := gc.NewWindow(fp.Height, fp.Width, fp.Y, fp.X)
	window.Box(0, 0)

	files, _ := ioutil.ReadDir(fp.Directory)
	fp.Files = files

	PrintHiglighted(window, 1, 1, 0, fp.Selected, "..", fp.IsActive)
	for i, f := range files {
		PrintHiglighted(window, i+2, 1, i, fp.Selected-1, RowOutput(f.Name(), f.Size(), f.Mode().IsDir(), fp.Width), fp.IsActive)
	}

	var msg = " [ " + fp.Directory + " ] "
	if fp.IsActive {
		window.AttrOn(gc.A_BOLD)
		window.AttrOn(gc.A_STANDOUT)
	}
	window.MovePrint(0, (fp.Width-len(msg))/2, msg)
	if fp.IsActive {
		window.AttrOff(gc.A_BOLD)
		window.AttrOff(gc.A_STANDOUT)
	}

	fp.Panel = gc.NewPanel(window)
}

func (fp *FilePanel) Redraw() {
	fp.Draw()
	gc.Update()
}

func (fp *FilePanel) Select(index int) {
	if index > -1 && index < len(fp.Files)+1 {
		fp.Selected = index
		fp.Redraw()
	}
}

func (fp *FilePanel) HideHidden() {
	var newFiles []os.FileInfo
	for i := range fp.Files {
		if fp.Files[i].Name()[0] == '.' {
			newFiles = append(newFiles, fp.Files[i])
		}
	}
	fp.Files = newFiles
	fp.Redraw()
}

func (fp *FilePanel) Execute() {
	if fp.Selected == 0 {
		fp.GoUp()
		return
	}
	var SelectedFile = fp.Directory + fp.Files[fp.Selected-1].Name()
	fileInfo, err := os.Stat(SelectedFile)
	if err != nil {
		log.Fatal(err)
	}
	if fileInfo.IsDir() {
		fp.GoToDir(SelectedFile)
	}
}

func (fp *FilePanel) GoToDir(folderName string) {
	fp.Directory = folderName + "/"
	fp.Select(0)
}

func (fp *FilePanel) GoUp() {
	var folders = strings.Split(fp.Directory, "/")
	fp.Directory = strings.Join(folders[:len(folders)-2], "/") + "/"
	fp.Select(0)
}

func (fp *FilePanel) ToggleActivity() {
	fp.IsActive = !fp.IsActive
	fp.Redraw()
}
