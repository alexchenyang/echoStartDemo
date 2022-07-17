package main

import (
	"github.com/alexchenyang/echoStartDemo/controller"
	"github.com/labstack/echo/v4"
)

// 加上配置文件读取，conf, 日志  build.sh  control.sh start.sh
func main() {

	echoInst := echo.New()

	registerRouter(echoInst)

	echoInst.Start(":8080")

}

func registerRouter(echoInstance *echo.Echo) {
	routerGroup := echoInstance.Group("startDemo")

	routerGroup.Match([]string{echo.POST, echo.GET}, "/test1", controller.Test1)
}
