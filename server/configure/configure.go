package configure

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Myconf struct {
	Port string
	DebugMode string
	Mysql
	MyRedis
}

type Mysql struct {
	Username string
	Password string
	Host string
	Port string
	DbName string
}
type MyRedis struct {
	Username string
	Password string
	Address string
}

// 全局配置对象
var Config Myconf


// 读取配置文件
func ReadConfigureFile(filepath string) (err error ) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("配置文件不存在: [ %s ]",err)
		return
	}

	//把yaml形式的字符串解析成struct类型
	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		fmt.Printf("配置文件解析错误: [ %s ]",err)
		return
	}

	if(Config.Port==""){
		fmt.Println("端口设置错误")
	}
	Config.Port = ":" + Config.Port
	return
}
