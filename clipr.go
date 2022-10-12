package main

import (
	"log"
	"os"
	"time"
)

type info struct {
	Id int `json:"id"`
}

type state struct {
	val  int
	info info
	arg  string
}

type val string

type meta struct {
	LastModified time.Time `json:"last_modified"`
}

type entry struct {
	Val  val  `json:"val"`
	Meta meta `json:"meta"`
}

type clipboard struct {
	List []entry `json:"list"`
}

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
	}
}
