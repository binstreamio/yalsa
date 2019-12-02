package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var port = flag.Int("port", 8080, "http server listen port")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] live_stream_url\n", os.Args[0])
		flag.PrintDefaults()
	}

	gob.Register(&yalsaPoint{})

	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func(interruptChannel chan os.Signal) {
		for range interruptChannel {
			savePoints()
			os.Exit(0)
		}
	}(interruptChannel)
}

var sourceRetries = 0

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	fmt.Println(loadPoints())

	streamInfoChan := make(chan *yalsaFrame)

	go output(streamInfoChan)

	go func() {
		for {
			sourceRetries++
			iCtx, err := openStream(flag.Arg(0))
			if err == nil {
				demuxing(iCtx, streamInfoChan)
			}

			time.Sleep(time.Second * 3)
		}
	}()

	startHTTPServer(*port)
}
