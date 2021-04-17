package routes

import (
	"fmt"
	"github.com/mattn/go-colorable"
)

import (
	controller "github.com/psyonly/FishersInSDV/controller/index"
)

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-colorable"
)

var (
	// 总路由
	engine *gin.Engine
)

func InitRoute() error {
	// 启用gin的颜色
	gin.ForceConsoleColor()
	// 替换writer
	gin.DefaultWriter = colorable.NewColorableStdout()
	// 初始化engine
	engine = gin.Default()
	// 路由注册
	engine.GET("/", controller.EntryPage)

	// 启动路由监听请求并回应
	err := engine.Run(":9876")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
