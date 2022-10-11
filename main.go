package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"unicode/utf8"
)

type info struct {
	Id int `json:"id"`
}

type state struct {
	val  int
	info info
	arg  string
}

type text struct {
	id     int
	val    string
	marked bool
}

type clipboard []text

type clipr struct {
	infolog   *log.Logger
	errorlog  *log.Logger
	state     state
	clipboard clipboard
}

var f *os.File

func newClipr(f *os.File) *clipr {
	infolog := log.New(f, "INFO\t", log.LstdFlags)
	errorlog := log.New(f, "ERROR\t", log.LstdFlags|log.Lshortfile)
	return &clipr{
		infolog:  infolog,
		errorlog: errorlog,
	}
}

func (c *clipr) initCurrentState() {
	if val := os.Getenv("ROFI_RETV"); val != "" {
		c.infolog.Printf("State: %s", val)
		c.state.val, _ = strconv.Atoi(val)
	}

	if info := os.Getenv("ROFI_INFO"); info != "" {
		c.infolog.Printf("Info: %s", info)
		json.Unmarshal([]byte(info), &c.state.info)
	}

	args := os.Args
	if len(args) > 1 {
		for k, v := range args {
			if k == 1 {
				c.infolog.Printf("Arg: %s", v)
				c.state.arg = v
			}
		}
	}
}

func (c *clipr) initClipboard() {
	// TODO: Read from a json file saved on disk
}

func (c *clipr) copyToClipboard() {
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

func rofiEncode(txt string) (enctxt string) {
	rtxt := make([]rune, 0, utf8.RuneCountInString(txt))
	for _, ch := range txt {
		if ch == '\n' || ch == '\r' {
			rtxt = append(rtxt, '\xA0')
		} else {
			rtxt = append(rtxt, ch)
		}
	}
	enctxt = string(rtxt)
	return
}

func rofiDecode(enctxt string) (txt string) {
	ctxt := make([]rune, 0, utf8.RuneCountInString(enctxt))
	for _, ch := range enctxt {
		if ch == '\xA0' {
			ctxt = append(ctxt, '\n')
		} else {
			ctxt = append(ctxt, ch)
		}
	}
	txt = string(ctxt)
	return
}

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
