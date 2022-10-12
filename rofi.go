package main

import (
	"io"
	"os/exec"
	"strings"
)

func (c *clipr) copySelection() {
	txt := rofiDecode(c.state.arg)
	if len(txt) == 0 {
		c.errorlog.Fatalln("Selection empty! Failed to copy to clipboard.")
		return
	}

	// copy selection to system clipboard
	pr, pw := io.Pipe()
	xselcmd := exec.Command("xsel", "-ib")
	xselcmd.Stdin = pr
	go func() {
		defer pw.Close()
		io.Copy(pw, strings.NewReader(txt))
	}()
	err := xselcmd.Run()
	if err != nil {
		c.errorlog.Fatalln(err)
	}
}

func (c *clipr) deleteSelection() {
	txt := rofiDecode(c.state.arg)
	if len(txt) == 0 {
		c.errorlog.Fatalln("Selection empty! Failed to delete clipboard entry.")
		return
	}
	c.delete(entry{
		Val: val(txt),
	})

	// if the last entry is deleted
	// clear system clipboard as well
	if len(c.clipboard.List) == 0 {
		err := exec.Command("xsel", "-cb").Run()
		if err != nil {
			c.errorlog.Fatalln(err)
		}
	}
}
