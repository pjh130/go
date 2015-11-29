package redislib

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

/**********************************************************************
 * 功能描述： 生成一个redis操作对象
 * 输入参数： redisAddr-地址 pwd-密码 poolsize-连接池数量
 * 输出参数： 无
 * 返 回 值： *RedisProvider-对象 error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func NewRedisProvider(redisAddr string, pwd string, poolsize int) (*RedisProvider, error) {
	v := &RedisProvider{}
	err := v.RedisInit(redisAddr, pwd, poolsize)
	if nil != err {
		fmt.Println("NewRedisProvider err: " + err.Error())
		return nil, err
	} else {
		fmt.Println("NewRedisProvider OK")
	}

	return v, nil
}

// redis max pool size
var MAX_POOL_SIZE = 100

// redis session provider
type RedisProvider struct {
	savePath string
	poolsize int
	password string
	poollist *redis.Pool
}

func (rp *RedisProvider) RedisInit(redisAddr string, pwd string, poolsize int) error {
	rp.password = pwd
	rp.poolsize = poolsize

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
