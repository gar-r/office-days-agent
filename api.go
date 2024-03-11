package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func startApiServer(port int) {
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.GET("/list", listHandler)
	r.POST("/clear", clearHandler)
	r.Run(fmt.Sprintf(":%d", port))
}

func listHandler(ctx *gin.Context) {
	data, err := store.Load()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, data)
}

func clearHandler(ctx *gin.Context) {
	err := store.Clear()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	} else {
		ctx.Status(http.StatusOK)
	}
}
