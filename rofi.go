package main

import (
	"fmt"
	"io"
	"os/exec"
)

func (c *clipr) copy() {
	txt := rofiDecode(c.state.arg)
	if len(txt) == 0 {
		c.errorlog.Fatalln("Selection empty! Copy to clipboard failed.")
		return
	}

	// copy selection to clipboard
	pr, pw := io.Pipe()
	xselcmd := exec.Command("xsel", "-ib")
	xselcmd.Stdin = pr
	go func() {
		defer pw.Close()
		fmt.Fprintf(pw, txt)
	}()
	err := xselcmd.Run()
	if err != nil {
		c.errorlog.Fatalln(err)
	}
}
