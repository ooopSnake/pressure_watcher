package main

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"os"
)

type RequestHandler func() interface{}

func StartServer(addr string, fun RequestHandler) {
	e := echo.New()
	defer e.Close()
	e.HidePort = true
	e.HideBanner = true
	e.GET("/*", func(ctx echo.Context) error {
		jsonObj := fun()
		return ctx.JSONPretty(http.StatusOK, jsonObj, "  ")
	})
	err := e.Start(addr)
	if err != nil {
		fmt.Println("http server start error : ", err)
		os.Exit(2)
	}
}
