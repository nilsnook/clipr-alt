package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

const (
	// Enter
	SELECT = 1
	// kb-custom-1
	DELETE = 10
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "clipr.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// create new clipr with specified log file
	c := newClipr(f)
	// initialize current state
	c.initCurrentState()
	// initialize clipboard
	c.initClipboard()

	// handle rofi events
	switch c.state.val {
	case SELECT:
		c.copySelection()
		c.getLatestTextFromClipboard()
	case DELETE:
		c.deleteSelection()
	default:
		c.getLatestTextFromClipboard()
	}

	// use hot keys
	// for events like delete
	fmt.Println("\000use-hot-keys\x1ftrue")
	// render clipboard
	c.renderClipboard()
}
