package main

import (
	"errors"
	"os"
	"path"
)

const (
	DB_DIR  = ".local/share/clipr"
	DB_FILE = "clipr.db"
)

func (c *clipr) getDBFile(flag int) (f *os.File) {
	DB := path.Join(c.dbdir, DB_FILE)
	f, err := os.OpenFile(DB, flag, 0644)
	if err != nil {
		c.errorlog.Fatalln(err)
	}
	return
}

func (c *clipr) initClipboard() {
	c.createClipboardIfNotExists()
	c.read()
}

func (c *clipr) createClipboardIfNotExists() {
	var f *os.File
	defer f.Close()

	DB := path.Join(c.dbdir, DB_FILE)
	if _, err := os.Stat(DB); errors.Is(err, os.ErrNotExist) {
		// create dir
		err = createDirIfNotExists(c.dbdir)
		if err != nil {
			c.errorlog.Fatalln(err)
		}

		// create db file
		f, err = os.OpenFile(DB, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			c.errorlog.Fatalln(err)
		}

		// create a clipboard and save to file
		x := clipboard{
			List: []entry{},
		}
		err = writeJSON(f, x)
		if err != nil {
			c.errorlog.Fatalln(err)
		}
	}
}

func (c *clipr) read() {
	var err error
	c.clipboard, err = readJSON(c.getDBFile(os.O_RDONLY))
	if err != nil {
		c.errorlog.Fatalln(err)
	}
}

func (c *clipr) write(e entry) {
	c.clipboard.List = append(c.clipboard.List, e)
	s := newSet(c.clipboard.List...)
	c.clipboard.List = s.entries()
	err := writeJSON(c.getDBFile(os.O_WRONLY), c.clipboard)
	if err != nil {
		c.errorlog.Fatalln(err)
	}
}

func (c *clipr) delete(e entry) {
	s := newSet(c.clipboard.List...)
	s.delete(e)
	c.clipboard.List = s.entries()
	err := writeJSON(c.getDBFile(os.O_WRONLY), c.clipboard)
	if err != nil {
		c.errorlog.Fatalln(err)
	}
}
