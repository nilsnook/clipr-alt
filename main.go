package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
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

	// check for current state
	// SELECT - put selected text in clipboard
	if c.state.val == 1 {
		c.copy()
	}

	// use hot keys
	// fmt.Println("\000use-hot-keys\x1ftrue")

	// get latest text from X clipboard
	out, err := exec.Command("xsel", "-ob").Output()
	if err != nil {
		c.errorlog.Fatalln(err)
	}
	newEntry := entry{
		// replace newline (\n) or carriage return (\r) with '\xA0'
		Val: val(rofiEncode(string(out))),
		// update last modified time
		Meta: meta{
			LastModified: time.Now(),
		},
	}
	c.write(newEntry)

	// rofitxt := fmt.Sprintf("%s\000info\x1f{\"id\":1}", string(enctxt))
	// fmt.Println(rofitxt)
	// fmt.Println("This is a test\000info\x1f{\"id\":2}")

	for _, e := range c.clipboard.List {
		fmt.Println(e.Val)
	}
}
