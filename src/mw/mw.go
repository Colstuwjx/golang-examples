package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// source: https://gist.github.com/jerblack/4b98ba48ed3fb1d9f7544d2b1a1be287

func main() {
	fn := logOutput()
	defer fn()
	var words = "Hello\nDoctor\nName\nContinue\nYesterday\nTomorrow"
	for i := 0; i < 10; i++ {
		log.Println(i) // stderr

		fmt.Println(i)     // stdout
		fmt.Println(words) // stdout
	}
}

func logOutput() func() {
	logfile := `logfile`
	// open file read/write | create if not exist | clear file at open if exists
	f, _ := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	// save existing stdout (/dev/stdout) | MultiWriter writes to saved stdout and file
	out := os.Stdout
	mw := io.MultiWriter(out, f)

	// get pipe reader and writer | writes to pipe writer come out pipe reader
	r, w, _ := os.Pipe()

	// replace stdout,stderr with pipe writer | all writes to stdout, stderr will go through pipe instead (fmt.print, log)
	os.Stdout = w
	os.Stderr = w

	// writes with log.Print should also write to mw
	log.SetOutput(mw)

	//create channel to control exit | will block until all copies are finished
	exit := make(chan bool)

	go func() {
		// copy all reads from pipe to multiwriter, which writes to stdout and file
		_, _ = io.Copy(mw, r)
		// when r or w is closed copy will finish and true will be sent to channel
		exit <- true
	}()

	// function to be deferred in main until program exits
	return func() {
		// close writer then block on exit channel | this will let mw finish writing before the program exits
		_ = w.Close()
		<-exit
		// close file after all writes have finished
		_ = f.Close()
	}
}
