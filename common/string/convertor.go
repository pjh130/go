package string

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

// convert any numeric value to int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

//数字转换成字节
func NumToBytes(data interface{}, bigEndian bool) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})

	var err error
	if bigEndian {
		err = binary.Write(bytesBuffer, binary.BigEndian, data)
	} else {
		err = binary.Write(bytesBuffer, binary.LittleEndian, data)
	}
	return bytesBuffer.Bytes(), err
}

//字节转换成数字
func BytesToNum(b []byte, data interface{}, bigEndian bool) error {
	bytesBuffer := bytes.NewBuffer(b)

	var err error
	if bigEndian {
		err = binary.Read(bytesBuffer, binary.BigEndian, data)
	} else {
		err = binary.Read(bytesBuffer, binary.LittleEndian, data)
	}

	return err
}

//解析int参数
func ParseUrlInt64(values url.Values, key string, base int, bitSize int) (int64, error) {
	temp := values.Get(key)
	return strconv.ParseInt(temp, base, bitSize)
}

//解析int参数
func ParseUrlInt(values url.Values, key string) (int, error) {
	temp := values.Get(key)
	return strconv.Atoi(temp)
}

//解析uint64参数
func ParseUrlUint64(values url.Values, key string, base int, bitSize int) (uint64, error) {
	temp := values.Get(key)
	return strconv.ParseUint(temp, base, bitSize)
}

//解析float64参数
func ParseUrlFloat64(values url.Values, key string, bitSize int) (float64, error) {
	temp := values.Get(key)
	return strconv.ParseFloat(temp, bitSize)
}

//解析Bool参数
func ParseUrlBool(values url.Values, key string) (bool, error) {
	temp := values.Get(key)
	return strconv.ParseBool(temp)
}
