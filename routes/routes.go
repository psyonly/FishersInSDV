package routes

import (
	"fmt"
)

import (
	home "github.com/psyonly/FishersInSDV/controller/index"
	"github.com/psyonly/FishersInSDV/settings"
)

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
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
	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/", home.EntryPage)

		apiGroup.GET("/version", home.GetVersion)

		apiGroup.GET("/fishes", home.GetFishes)
	}

	// 启动路由监听请求并回应
	err := engine.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
