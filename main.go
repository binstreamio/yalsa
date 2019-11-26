package main

import (
	"flag"
	"time"
)

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

			time.Sleep(time.Second)
		}
	}()

	startHTTPServer(8080)
}
