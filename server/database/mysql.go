package database



import (
	"fmt"
	"time"

	"video_application/server/conf"

	"github.com/gin-gonic/gin"
	"gorm"
	_ "gorm/dialects/mysql"

)

var DB *gorm.DB

func init() {

	name := conf.Config.Mysql.Username
	pswd := conf.Config.Mysql.Password
	host := conf.Config.Mysql.Host
	port := conf.Config.Mysql.Port
	dbname := conf.Config.Mysql.DbName

	db, err := gorm.Open("database",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", name, pswd, host, port, dbname))
	if err != nil {
		panic(err.Error())
	}

	// 开发模式和生产模式日志开关
	debug := conf.Config.DebugMode
	if gin.DebugMode != debug {
		db.LogMode(false)
	} else {
		db.LogMode(true)
	}

	// 设置可重用连接的最大时间量 如果d<=0，则永远重用连接
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//设置到数据库的最大打开连接数 如果n<=0，则不限制打开的连接数 默认值为0
	db.DB().SetMaxOpenConns(0)
	// 设置空闲中的最大连接数 默认最大空闲连接数当前为2 如果n<=0，则不保留空闲连接
	db.DB().SetMaxIdleConns(10)

	DB = db
}

/*
	开启 database 事务操作
	支持一次传递多个 *gorm.DB 执行语句(exce)
*/
func Transaction(dbs ...*gorm.DB) bool {

	tx := DB.Begin()

	for _, db := range dbs {
		//tx = db
		//if tx.Error != nil {
		//	tx.Rollback()
		//	return false
		//}
		if db.Error != nil {
			tx.Rollback()
			return false
		}
	}

	tx.Commit()
	return true
}

