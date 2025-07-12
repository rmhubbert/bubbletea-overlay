package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var Writer io.Writer

func Log(messages ...string) {
	if Writer == nil {
		return
	}
	spew.Fdump(Writer, messages)
}

func TimeStamp() error {
	_, err := fmt.Fprintln(Writer, time.Now().UnixMilli())
	if err != nil {
		return err
	}
	return nil
}

// InitWriter creates a file to use for debugging messages, if the DEBUG environment variable has
// been set.
func InitWriter() {
	var debugFile *os.File
	if _, ok := os.LookupEnv("DEBUG"); ok {
		var err error
		debugFile, err = os.OpenFile("debug.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
		if err != nil {
			log.Fatal(err)
		}
	}
	Writer = debugFile
}
