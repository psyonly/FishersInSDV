package main

import (
	"fmt"
)

import (
	"github.com/psyonly/FishersInSDV/routes"
	"github.com/psyonly/FishersInSDV/settings"
)

func main() {
	var (
		err error
	)

	// 读取配置
	err = settings.InitSettings()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 初始化并启动路由
	err = routes.InitRoute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
