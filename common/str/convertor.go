package str

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"unsafe"
)

//判断是否是小字序
func IsLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb
	return (b == 0x04)
}

//数字转换成字节(不支持int和uint类型，需要根据平台和编译器转换成int32和uint32)
func NumToBytes(data interface{}, bigEndian bool) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	var err error

	//由于编译器和平台int和uint占用的字节数不一样,需要做特殊处理
	t := reflect.TypeOf(data)
	switch data.(type) {
	case *int:
		if 4 == t.Size() {
			var temp32 int32 = int32(*data.(*int))
			if bigEndian {
				err = binary.Write(bytesBuffer, binary.BigEndian, temp32)
			} else {
				err = binary.Write(bytesBuffer, binary.LittleEndian, temp32)
			}
		} else {
			var temp64 int64 = int64(*data.(*int))
			if bigEndian {
				err = binary.Write(bytesBuffer, binary.BigEndian, temp64)
			} else {
				err = binary.Write(bytesBuffer, binary.LittleEndian, temp64)
			}
		}
	case int:
		if 4 == t.Size() {
			var temp32 int32 = int32(data.(int))
			if bigEndian {
				err = binary.Write(bytesBuffer, binary.BigEndian, temp32)
			} else {
				err = binary.Write(bytesBuffer, binary.LittleEndian, temp32)
			}
		} else {
			var temp64 int64 = int64(data.(int))
			if bigEndian {
				err = binary.Write(bytesBuffer, binary.BigEndian, temp64)
			} else {
				err = binary.Write(bytesBuffer, binary.LittleEndian, temp64)
			}
		}
	case *uint:

	case uint:
		var temp32 uint32 = uint32(data.(uint))
		var temp64 uint64 = uint64(data.(uint))

		if 4 == t.Size() {
			if bigEndian {
				err = binary.Write(bytesBuffer, binary.BigEndian, temp32)
			} else {
				err = binary.Write(bytesBuffer, binary.LittleEndian, temp32)
			}
		} else {
			if bigEndian {
				err = binary.Write(bytesBuffer, binary.BigEndian, temp64)
			} else {
				err = binary.Write(bytesBuffer, binary.LittleEndian, temp64)
			}
		}
	default:
		if bigEndian {
			err = binary.Write(bytesBuffer, binary.BigEndian, data)
		} else {
			err = binary.Write(bytesBuffer, binary.LittleEndian, data)
		}
	}

	return bytesBuffer.Bytes(), err
}

//字节转换成数字
func BytesToNum(b []byte, data interface{}, bigEndian bool) error {
	bytesBuffer := bytes.NewBuffer(b)
	var err error

	//由于编译器和平台int和uint占用的字节数不一样,需要做特殊处理
	t := reflect.TypeOf(data)
	switch data.(type) {
	case *int:
		if 4 == t.Size() {
			var temp32 int32 = 0
			if bigEndian {
				err = binary.Read(bytesBuffer, binary.BigEndian, &temp32)
			} else {
				err = binary.Read(bytesBuffer, binary.LittleEndian, &temp32)
			}
			p := reflect.ValueOf(data).Elem()
			p.SetInt(int64(temp32))
		} else {
			var temp64 int64 = 0
			if bigEndian {
				err = binary.Read(bytesBuffer, binary.BigEndian, &temp64)
			} else {
				err = binary.Read(bytesBuffer, binary.LittleEndian, &temp64)
			}
			p := reflect.ValueOf(data).Elem()
			p.SetInt(temp64)
		}
		//	case int:/*如果传进来的不是指针，赋值不会成功的，不走这个分支，往下走让它报错*/
	case *uint:
		if 4 == t.Size() {
			var temp32 uint32 = 0
			if bigEndian {
				err = binary.Read(bytesBuffer, binary.BigEndian, &temp32)
			} else {
				err = binary.Read(bytesBuffer, binary.LittleEndian, &temp32)
			}
			p := reflect.ValueOf(data).Elem()
			p.SetUint(uint64(temp32))
		} else {
			var temp64 uint64 = 0
			if bigEndian {
				err = binary.Read(bytesBuffer, binary.BigEndian, &temp64)
			} else {
				err = binary.Read(bytesBuffer, binary.LittleEndian, &temp64)
			}
			p := reflect.ValueOf(data).Elem()
			p.SetUint(temp64)
		}
		//	case uint:/*如果传进来的不是指针，赋值不会成功的，不走这个分支，往下走让它报错*/
	default:
		if bigEndian {
			err = binary.Read(bytesBuffer, binary.BigEndian, data)
		} else {
			err = binary.Read(bytesBuffer, binary.LittleEndian, data)
		}
	}

	return err
}

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
