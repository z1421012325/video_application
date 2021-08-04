package cache

import (
	"time"

	"video_application/server/conf"

	"github.com/go-redis/redis"
)



type RedisPool struct {
	*redis.Client
}

var (
	//RDB *cache.Client
	err   error
	RPool RedisPool
)

// 被引用自动初始化 前提需要配置文件加载
// todo 之后进行更加详细的配置
func init() {
	rb := redis.NewClient(&redis.Options{
		Addr:         conf.Config.MyRedis.Address,
		Password:     conf.Config.MyRedis.Password,
		DB:           0,
		DialTimeout:  0,
		ReadTimeout:  time.Second,
		WriteTimeout: 0,
		PoolSize:     0,
		MinIdleConns: 10,
		MaxConnAge:   0,
	})

	_, err = rb.Ping().Result()
	if err != nil {
		panic(err)
	}
	//RDB = rb
	RPool.Client = rb
}



func CacheGet(key string) (value string){
	value,_ = RPool.Get(key).Result()
	return value
}


func CacheSet(key,value string,ex time.Duration) (err error){
	_,err =RPool.Set(key,value,ex).Result()
	return
}

func CacheDel(key string) (err error){
	_,err = RPool.Del(key).Result()
	return
}


func CacheVideoNewLook(key string,score float64,vid int64)(err error) {
	_,err = RPool.ZAdd(key,redis.Z{
		Score:  score,
		Member: vid,
	}).Result()
	return
}









