package main

import (
	"bytes"
	"fmt"
	"github.com/jroimartin/gocui"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Content struct {
	name string
	dir  bool
}

var wd bytes.Buffer

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()

		_, err := v.Line(cy+2)
		if err != nil {
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func cursorReset(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, oy := v.Origin()
		if err := v.SetCursor(0, 0); err != nil && oy > 0 {
			if err := v.SetOrigin(0, 0); err != nil {
				return err
			}
		}
	}
	return nil
}

func moveToDir(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	var newPath string
	if l == ".." {
		newPath = removeFromPath(wd.String())
	} else {
		newPath = addToPath(wd.String(), l)
	}

	fi, err := os.Stat(newPath)

	if err != nil {
		log.Println(err)
		return nil
	}

	if !fi.IsDir() {
		return nil
	}

	wd.Reset()
	wd.WriteString(newPath)
	update(g, newPath)

	return nil
}

func addToPath(path string, new string) string {
	return filepath.Join(path, new)
}

func removeFromPath(path string) string {
	return filepath.Dir(path)
}

func update(g *gocui.Gui, dir string) {
	v, err := g.View("top")
	if err != nil {
		log.Panicln(err)
	}

	v.Clear()
	fmt.Fprintln(v, dir)

	v, err = g.View("main")
	if err != nil {
		log.Panicln(err)
	}

	v.Clear()
	if dir != "/" {
		fmt.Fprintln(v, "..")
	}

	for _, c := range getDirContent(dir) {
		fmt.Fprintln(v, c.name)
	}

	cursorReset(g, v)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func keyBindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("main", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", gocui.KeyEnter, gocui.ModNone, moveToDir); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("top", 0, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		fmt.Fprintln(v, getWd())
	}
	if v, err := g.SetView("main", 0, 3, maxX-1, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack

		update(g, getWd())

	}
	if _, err := g.SetCurrentView("main"); err != nil {
		log.Panicln(err)
	}

	return nil
}

func getDirContent(wd string) []Content {
	files, err := ioutil.ReadDir(wd)

	if err != nil {
		log.Fatal(err)
	}

	count := len(files)
	content := make([]Content, 0, count)

	for _, f := range files {
		content = append(content, Content{f.Name(), f.IsDir()})
	}

	return content
}

func getWd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func main() {

	wd.WriteString(getWd())

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keyBindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
