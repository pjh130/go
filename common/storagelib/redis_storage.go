package storagelib

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"strings"
	"sync"
	"time"
)

var DriveRedisStorage string = "redis-storage"

type redisStorage struct {
	_pool *redis.Pool
	_buf  *bytes.Buffer
	sync.Mutex
}

func NewRedisStorage(redisAddr string, pwd string, poolSide int) Storage {
	pool, _ := NewRedisPool(redisAddr, pwd, poolSide)
	return &redisStorage{
		_pool: pool,
		_buf:  new(bytes.Buffer),
	}
}

func NewRedisPool(redisAddr string, pwd string, poolSide int) (*redis.Pool, error) {
	pool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.DialTimeout("tcp", redisAddr, 0, 1*time.Second, 1*time.Second)
		if err != nil {
			return nil, err
		}
		if pwd != "" {
			if _, err := c.Do("AUTH", pwd); err != nil {
				c.Close()
				return nil, err
			}
		}
		return c, err
	}, poolSide)
	return pool, nil
}

func (this *redisStorage) getByte(v interface{}) ([]byte, error) {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	enc := gob.NewEncoder(this._buf)
	err := enc.Encode(v)
	if err == nil {
		b := this._buf.Bytes()
		this._buf.Reset()
		return b, nil
	}
	if strings.Index(err.Error(), "type not registered") != -1 {
		panic(err)
	}
	return nil, err
}

func (this *redisStorage) decodeBytes(b []byte, dst interface{}) error {
	this.Mutex.Lock()
	defer this.Mutex.Unlock()
	this._buf.Write(b)
	dec := gob.NewDecoder(this._buf)
	err := dec.Decode(dst)
	this._buf.Reset()
	return err
}

func isBaseOfStruct(v interface{}) bool {
	valueType := reflect.TypeOf(v)
	kind := valueType.Kind()
	if kind == reflect.Ptr {
		kind = valueType.Elem().Kind()
	}
	return kind == reflect.Struct || kind == reflect.Map || kind == reflect.Array
}

func (this *redisStorage) getRedisBytes(key string) ([]byte, error) {
	conn := this._pool.Get()
	src, err := redis.Bytes(conn.Do("GET", key))
	conn.Close()
	return src, err
}

func (this *redisStorage) Driver() string {
	return DriveRedisStorage
}

func (this *redisStorage) Get(key string, dst interface{}) error {
	if isBaseOfStruct(dst) {
		src, err := this.getRedisBytes(key)
		if err == nil {
			err = this.decodeBytes(src, dst)
		}
		return err
	}
	return errors.New("dst must be struct")
}

func (this *redisStorage) GetBool(key string) (bool, error) {
	conn := this._pool.Get()
	src, err := redis.Bool(conn.Do("GET", key))
	conn.Close()
	return src, err
}

func (this *redisStorage) GetInt(key string) (int, error) {
	conn := this._pool.Get()
	src, err := redis.Int(conn.Do("GET", key))
	conn.Close()
	return src, err
}

func (this *redisStorage) GetInt64(key string) (int64, error) {
	conn := this._pool.Get()
	src, err := redis.Int64(conn.Do("GET", key))
	conn.Close()
	return src, err
}

func (this *redisStorage) GetString(key string) (string, error) {
	d, err := this.getRedisBytes(key)
	if err != nil {
		return "", err
	}
	return string(d), err
}

func (this *redisStorage) GetFloat64(key string) (float64, error) {
	conn := this._pool.Get()
	src, err := redis.Float64(conn.Do("GET", key))
	conn.Close()
	return src, err
}

//Get raw value
func (this *redisStorage) GetRaw(key string) (interface{}, error) {
	conn := this._pool.Get()
	replay, err := conn.Do("GET", key)
	conn.Close()
	return replay, err
}

func (this *redisStorage) Set(key string, v interface{}) error {
	var err error
	var redisValue interface{} = v

	if isBaseOfStruct(v) {
		redisValue, err = this.getByte(v)
	}

	conn := this._pool.Get()
	_, err = conn.Do("SET", key, redisValue)
	conn.Close()
	return err
}
func (this *redisStorage) Del(key string) {
	conn := this._pool.Get()
	conn.Do("DEL", key)
	defer conn.Close()
}

func (this *redisStorage) SetExpire(key string, v interface{}, seconds int64) error {
	var err error
	var redisValue interface{} = v

	if isBaseOfStruct(v) {
		redisValue, err = this.getByte(v)
	}

	conn := this._pool.Get()
	_, err = conn.Do("SETEX", key, seconds, redisValue)
	conn.Close()
	return err
}
