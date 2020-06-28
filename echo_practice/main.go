package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func main() {
	e:= echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/",hello)
	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK,"welcome to echo")
}

func setMiddleware(e *echo.Echo) {
	// access log 输出到文件中
	accessLog, err := os.OpenFile("log/access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	// 同时输出到终端和文件
	middleware.DefaultLoggerConfig.Output = accessLog
	e.Use(middleware.Logger())

	// 自定义 middleware
	//e.Use(AutoLogin)

	e.Use(middleware.Recover())
}