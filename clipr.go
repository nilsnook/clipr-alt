package main

import (
	"log"
	"os"
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

func newClipr(f *os.File) *clipr {
	infolog := log.New(f, "INFO\t", log.LstdFlags)
	errorlog := log.New(f, "ERROR\t", log.LstdFlags|log.Lshortfile)
	return &clipr{
		infolog:  infolog,
		errorlog: errorlog,
		// TODO: init clipboard
	}
}
