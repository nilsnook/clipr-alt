package main

import (
	"encoding/json"
	"errors"
	"os"
	"unicode/utf8"
)

func createDirIfNotExists(path string) (err error) {
	if _, err = os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(path, 0755)
	}
	return
}

func writeJSON(f *os.File, data clipboard) (err error) {
	err = json.NewEncoder(f).Encode(data)
	return
}

func readJSON(f *os.File) (data clipboard, err error) {
	err = json.NewDecoder(f).Decode(&data)
	return
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
