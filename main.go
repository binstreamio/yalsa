package main

import (
	"flag"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	iCtx, err := openStream(flag.Arg(0))
	if err != nil {
		return
	}

	streamInfoChan := make(chan *yalsaFrame)

	go demuxing(iCtx, streamInfoChan)

	output(streamInfoChan)
}
