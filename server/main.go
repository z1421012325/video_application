package main


import (
	"fmt"
	"os"

	"video_application/server/configure"

	_ "video_application/server/cache"		// 数据库与redis链接
	_ "video_application/server/database"

	"video_application/server/service"
)


func main() {
	// go run .\server\main.go .\config.yml
	if len(os.Args[1]) <= 1 {
		fmt.Println("查询不到配置文件")
		os.Exit(1)
	}

	// 读取配置
	err := configure.ReadConfigureFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	err = service.RunService(configure.Config.Port)
	if err != nil {
		os.Exit(1)
	}
}
