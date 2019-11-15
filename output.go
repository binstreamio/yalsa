package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/3d0c/gmf"
	gocache "github.com/patrickmn/go-cache"
)

var (
	framesCache = gocache.New(time.Minute*2, time.Second)

	fps1SecondCache   = gocache.New(time.Hour*24, time.Second)
	fps10SecondCache  = gocache.New(time.Hour*24, time.Second)
	fps100SecondCache = gocache.New(time.Hour*24, time.Second)

	f1, f10, f100 int
	keyframe      int
)

func count() {
	frames := framesCache.Items()
	timeNow := time.Now()

	f1, f10, f100 = 0, 0, 0
	keyframe = 0

	var f100Diff time.Duration

	for t, f := range frames {
		it, _ := strconv.ParseInt(t, 10, 64)
		ft := time.Unix(it/int64(time.Second), it%int64(time.Second))
		frame := f.Object.(*yalsaFrame)

		if frame.MediaType != gmf.AVMEDIA_TYPE_VIDEO {
			continue
		}

		if timeNow.Sub(ft) < time.Second {
			f1++

			if frame.Keyframe {
				keyframe = 1
			}
		}

		if timeNow.Sub(ft) < time.Second*10 {
			f10++
		}

		if timeNow.Sub(ft) < time.Second*100 {
			if timeNow.Sub(ft) > f100Diff {
				f100Diff = timeNow.Sub(ft)
			}
			f100++
		}
	}

	f10 /= 10
	if int(f100Diff/time.Second) > 0 {
		f100 /= int(f100Diff / time.Second)
	}

	fps1SecondCache.Add(fmt.Sprint(time.Now().UnixNano()), f1, 0)
	fps10SecondCache.Add(fmt.Sprint(time.Now().UnixNano()), f10, 0)
	fps100SecondCache.Add(fmt.Sprint(time.Now().UnixNano()), f100, 0)
}

func output(streamInfoChan chan *yalsaFrame) {
	startTime := time.Now()
	t := time.NewTicker(time.Second)
	for {
		select {
		case f := <-streamInfoChan:
			if time.Now().Sub(startTime) > 1*time.Second {
				framesCache.Add(fmt.Sprint(time.Now().UnixNano()), f, 0)
			}
		case <-t.C:
			count()
			fmt.Printf("[%s] %3d %3d %3d %2d\n", time.Now().Format("0102T15:04:05.000Z07"), f1, f10, f100, keyframe)

		}
	}
}
