#!/bin/bash

export PKG_CONFIG_PATH=/usr/lib/ffmpeg/pkgconfig/

go-bindata -o htmldata.go html/index.html
go build -o yalsa