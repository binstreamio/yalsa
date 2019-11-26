package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var port = flag.Int("port", 8080, "http server listen port")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] live_stream_url\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	streamInfoChan := make(chan *yalsaFrame)

	go output(streamInfoChan)

	go func() {
		for {
			iCtx, err := openStream(flag.Arg(0))
			if err == nil {
				demuxing(iCtx, streamInfoChan)
			}

			time.Sleep(time.Second * 5)
		}
	}()

	startHTTPServer(*port)
}
