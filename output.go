package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/3d0c/gmf"
	gocache "github.com/patrickmn/go-cache"
)

type yalsaPoint struct {
	TimeStamp int `json:"timestamp_ms"`
	KeyFrame  int `json:"keyframe"`
	BitRate   int `json:"bitrate"`
	F1        int `json:"avg_fps_1s"`
	F10       int `json:"avg_fps_10s"`
	F100      int `json:"avg_fps_100s"`

	SourceRetries int `json:"source_retries"`
}

var (
	framesCache = gocache.New(time.Minute*2, time.Second)
	pointsCache = gocache.New(time.Hour*24, time.Second)

	lastPointKey = ""
)

func savePoints() {
	pointsCache.SaveFile(fmt.Sprintf("yasla_%d.data", *port))
}

func loadPoints() error {
	return pointsCache.LoadFile(fmt.Sprintf("yasla_%d.data", *port))
}

func count() {
	frames := framesCache.Items()
	timeNow := time.Now()

	f1, f10, f100 := 0, 0, 0
	keyframe := 0
	bitrate := 0

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
				keyframe++
			}
			bitrate += frame.Size
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

	tms := time.Now().UnixNano() / int64(time.Millisecond)
	lastPointKey = fmt.Sprint(tms)

	pointsCache.SetDefault(lastPointKey, &yalsaPoint{
		TimeStamp:     int(tms),
		KeyFrame:      keyframe,
		BitRate:       bitrate * 8 / 1024,
		F1:            f1,
		F10:           f10,
		F100:          f100,
		SourceRetries: sourceRetries,
	})
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
			var point *yalsaPoint
			v, ok := pointsCache.Get(lastPointKey)
			if ok {
				point = v.(*yalsaPoint)
			} else {
				point = &yalsaPoint{}
			}

			fmt.Printf("[%s] %3d %3d %3d %5dK %2d\n",
				time.Now().Format("0102T15:04:05.000Z07"),
				point.F1, point.F10, point.F100, point.BitRate, point.KeyFrame)
		}
	}
}
