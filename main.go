package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

const (
	// TODO: create dir if not exists
	CLIPBOARD = "~/.local/share/clipr/clipr.db"
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

	// check for current state
	// SELECT - put selected text in clipboard
	if c.state.val == 1 {
		c.copyToClipboard()
	}

	// use hot keys
	fmt.Println("\000use-hot-keys\x1ftrue")

	// get latest text from X clipboard
	out, err := exec.Command("xsel", "-ob").Output()
	if err != nil {
		c.errorlog.Fatalln(err)
	}

	// replace newline (\n) or carriage return (\r) with '\xA0'
	enctxt := rofiEncode(string(out))
	rofitxt := fmt.Sprintf("%s\000info\x1f{\"id\":1}", string(enctxt))
	fmt.Println(rofitxt)
	fmt.Println("This is a test\000info\x1f{\"id\":2}")
}
