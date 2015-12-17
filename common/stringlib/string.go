package stringlib

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//判断是否是数字字符串
func IsNum(a string) bool {
	reg, _ := regexp.Compile("^\\d+$")
	return reg.MatchString(a)
}

//格式化json字串
func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}
	return jsons
}

//字串转换成unicode
func Str2Uinicode(str string) string {
	var v string
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		s := string(r[i])
		if len(s) <= 1 {
			v = v + s
		} else {
			v = v + "&#" + fmt.Sprintf("%d", r[i]) + ";"
		}
	}

	fmt.Println("v: ", v)

	return v
}

//把数字列表转换成字符
func Uint64ListToString(list []uint64, split string) string {
	var s string

	for i := 0; i < len(list); i++ {
		add := strconv.FormatUint(list[i], 10)
		if len(add) > 0 {
			if len(s) > 0 {
				s = s + split
			}
			s = s + add
		}
	}

	return s
}

//把数字列表转换成字符
func Uint64ListFromString(str string, split string) []uint64 {
	var v []uint64

	list := strings.Split(str, split)

	for i := 0; i < len(list); i++ {
		add, err := strconv.ParseUint(list[i], 10, 30)
		if nil == err {
			v = append(v, add)
		}
	}

	return v
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

//获取时间字串
func GetNowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//检查用户名是否正确
func IsUsernameValid(name string) error {
	var err error

	//判断用户名长度
	length := 6
	if len(name) < length {
		return errors.New("长度不能少于" + strconv.Itoa(length))
	}

	length = 15
	if len(name) > length {
		return errors.New("长度不能大于" + strconv.Itoa(length))
	}

	//是否含有非法字符
	//需要加/转义的正则表达式特殊字符  ^ $ ( ) [ ] { } . ? + * |
	var hzRegexp = regexp.MustCompile("[@]")
	if true == (hzRegexp.MatchString(name)) {
		return errors.New("含有非法字符")
	}

	//对单个字符校验
	r := []rune(name)
	for i := 0; i < len(r); i++ {
		s := string(r[i])

		//必须是键盘输入符号
		if len(s) > 1 {
			return errors.New("不是合法字符" + strconv.Itoa(length))
		}

		//判断是否是字母开头
		if 0 == i {
			var hzRegexp = regexp.MustCompile("[a-zA-Z]$")
			if false == (hzRegexp.MatchString(s)) {
				return errors.New("必须以字母开头")
			}
		}
	}
	return err
}

//整形转换成字节
func IntToBytesBig(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToIntBig(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

//整形转换成字节
func IntToBytesLittle(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToIntLittle(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)

	return int(x)
}

type argInt []int

// get int by index from int slice
func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}

// interface to string
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}
