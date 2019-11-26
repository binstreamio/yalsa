package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func _feedPoints(c *gin.Context) {
	lastPoint, ok := c.GetQuery("last_point")
	if !ok {
		lastPoint = "0"
	}

	points := pointsCache.Items()

	keys := make([]string, len(points))
	i := 0
	for key := range points {
		if key > lastPoint {
			keys[i] = key
			i++
		}
	}
	keys = keys[:i]

	sort.Strings(keys)
	resPoints := make([]*yalsaPoint, len(keys))
	for i := range keys {
		resPoints[i] = points[keys[i]].Object.(*yalsaPoint)
	}

	c.JSON(http.StatusOK, resPoints)
}

func startHTTPServer(port int) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/v1/points", _feedPoints)

	r.StaticFile("/home", "./html/index.html") // for dev purpose

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", MustAsset("html/index.html"))
	})

	fmt.Println("Running http server on", fmt.Sprintf(":%d", port))

	r.Run(fmt.Sprintf(":%d", port))
}
