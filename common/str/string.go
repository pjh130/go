package str

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"gopkg.in/iconv.v1"
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

// Use golang.org/x/text/encoding. Get page body and change it to utf-8
func ChangeCharsetEncoding(charset string, sor io.ReadCloser) string {
	ischange := true
	var tr transform.Transformer
	cs := strings.ToLower(charset)
	if cs == "gbk" {
		tr = simplifiedchinese.GBK.NewDecoder()
	} else if cs == "gb18030" {
		tr = simplifiedchinese.GB18030.NewDecoder()
	} else if cs == "hzgb2312" || cs == "gb2312" || cs == "hz-gb2312" {
		tr = simplifiedchinese.HZGB2312.NewDecoder()
	} else {
		ischange = false
	}

	var destReader io.Reader
	if ischange {
		transReader := transform.NewReader(sor, tr)
		destReader = transReader
	} else {
		destReader = sor
	}

	var sorbody []byte
	var err error
	if sorbody, err = ioutil.ReadAll(destReader); err != nil {
		return ""
	}
	bodystr := string(sorbody)

	return bodystr
}

//转换字符编码
func ChangeCodeType(src string, srcType, destType string) (string, error) {
	dest := ""

	//	cd, err := iconv.Open("gbk", "utf-8") // convert utf-8 to gbk
	cd, err := iconv.Open(destType, srcType)
	if err != nil {
		return dest, err
	}
	defer cd.Close()

	dest = cd.ConvString(src)

	return dest, err
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

//把数字字符串列表转换成数字列表
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

// Convert any type to string
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

// HexStr2int converts hex format string to decimal number.
func HexStr2Int(hexStr string) (int, error) {
	num := 0
	length := len(hexStr)
	for i := 0; i < length; i++ {
		char := hexStr[length-i-1]
		factor := -1

		switch {
		case char >= '0' && char <= '9':
			factor = int(char) - '0'
		case char >= 'a' && char <= 'f':
			factor = int(char) - 'a' + 10
		default:
			return -1, fmt.Errorf("invalid hex: %s", string(char))
		}

		num += factor * PowInt(16, i)
	}
	return num, nil
}

// Int2HexStr converts decimal number to hex format string.
func Int2HexStr(num int) (hex string) {
	if num == 0 {
		return "0"
	}

	for num > 0 {
		r := num % 16

		c := "?"
		if r >= 0 && r <= 9 {
			c = string(r + '0')
		} else {
			c = string(r + 'a' - 10)
		}
		hex = c + hex
		num = num / 16
	}
	return hex
}

// PowInt is int type of math.Pow function.
func PowInt(x int, y int) int {
	if y <= 0 {
		return 1
	} else {
		if y%2 == 0 {
			sqrt := PowInt(x, y/2)
			return sqrt * sqrt
		} else {
			return PowInt(x, y-1) * x
		}
	}
}
