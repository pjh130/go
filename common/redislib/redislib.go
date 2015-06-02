package redislib

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
	"time"
)

//全局变量
var redispder = &RedisProvider{}

func InitRedisSpder(RedisAddr, RedisPwd string) {
	err := redispder.RedisInit(120, RedisAddr, RedisPwd)
	if nil != err {
		fmt.Println("InitRedisSpder err: " + err.Error())
		os.Exit(2)
		return
	} else {
		fmt.Println("InitRedisSpder OK")
	}
}

func DoRedisComm(commandName string, args ...interface{}) (interface{}, error) {
	return redispder.Do(commandName, args...)
}

// redis max pool size
var MAX_POOL_SIZE = 100

// redis session provider
type RedisProvider struct {
	maxlifetime int64
	savePath    string
	poolsize    int
	password    string
	poollist    *redis.Pool
}

func (rp *RedisProvider) RedisInit(maxlifetime int64, redisAddr string, pwd string) error {
	rp.maxlifetime = maxlifetime
	rp.password = pwd
	rp.poolsize = MAX_POOL_SIZE

	rp.poollist = redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.DialTimeout("tcp", redisAddr, 0, 1*time.Second, 1*time.Second)
		if err != nil {
			return nil, err
		}
		if rp.password != "" {
			if _, err := c.Do("AUTH", rp.password); err != nil {
				c.Close()
				return nil, err
			}
		}
		return c, err
	}, rp.poolsize)

	return rp.poollist.Get().Err()
}

func (rp *RedisProvider) Do(commandName string, args ...interface{}) (interface{}, error) {
	c := rp.poollist.Get()
	if nil == c {
		return nil, errors.New("Can't get conn from poolist")
	}

	defer c.Close()

	reply, err := c.Do(commandName, args...)

	if nil != err {
		fmt.Println("redis do err: ", err)
	}

	return reply, err
}
