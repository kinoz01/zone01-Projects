package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/jroimartin/gocui"
)

var conn net.Conn
var wg sync.WaitGroup

func main() {
	var err error
	conn, err = net.Dial("tcp", "localhost:8989")
	if err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
	defer conn.Close()

	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer gui.Close()

	gui.SetManagerFunc(layout)

	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := gui.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, sendMessage); err != nil {
		log.Panicln(err)
	}

	wg.Add(1)
	go receiveMessages(gui)

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	wg.Wait()
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("chat", 0, 0, maxX-1, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Chat"
		v.Autoscroll = true
		v.Wrap = true
	}

	if v, err := g.SetView("input", 0, maxY-4, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input"
		v.Editable = true

		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
	}

	return nil
}

func sendMessage(g *gocui.Gui, v *gocui.View) error {
	input := v.Buffer()
	v.Clear()
	v.SetCursor(0, 0)
	input = strings.TrimSpace(input)

	if input != "" {
		fmt.Fprintln(conn, input)
	}

	return nil
}

func receiveMessages(g *gocui.Gui) {
	defer wg.Done()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		g.Update(func(gui *gocui.Gui) error {
			v, err := g.View("chat")
			if err != nil {
				return err
			}
			fmt.Fprintln(v, msg)
			return nil
		})
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
