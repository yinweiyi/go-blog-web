package redis

import (
	"blog/vendors/config"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var Pool *redis.Pool

func InitRedis() *redis.Pool {
	var (
		host     = config.GetString("database.redis.host")
		port     = config.GetString("database.redis.port")
		password = config.GetString("database.redis.password")
	)

	Pool = &redis.Pool{
		MaxIdle:     256,
		MaxActive:   0,
		IdleTimeout: time.Duration(120),
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%s", host, port),
				redis.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				redis.DialDatabase(0),
				redis.DialPassword(password),
			)
		},
	}

	return Pool
}

func Exec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := Pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	params := make([]interface{}, 0)
	params = append(params, key)

	if len(args) > 0 {
		for _, v := range args {
			params = append(params, v)
		}
	}
	return con.Do(cmd, params...)
}

//设置值
func Set(key, value string, expire int) error {
	_, err := Exec("set", key, value)
	if err != nil {
		return err
	}
	if expire > 0 {
		_, err = Exec("expire", key, expire)
	}
	return err
}

func ToString(data interface{}, err error) (string, error) {
	return redis.String(data, err)
}
