package rand

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

/**********************************************************************
 * 功能描述： 获取一个取值区间的随机数
 * 输入参数： min-最小值 max-最大值
 * 输出参数： 无
 * 返 回 值： error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func RandInt(min int, max int) int {
	r.Seed(time.Now().UTC().UnixNano())
	return min + r.Intn(max-min)
}

var (
	//数字和字母混搭
	numAlpha = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
	//纯数字
	numOnly = []byte(`0123456789`)
	//字母
	numLetter = []byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)
	//纯小写字母
	numLower = []byte(`abcdefghijklmnopqrstuvwxyz`)
	//纯大写字母
	numUpper = []byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZ`)
)

/**********************************************************************
 * 功能描述： 获取随机数字字符串
 * 输入参数： length-随机数字串的长度
 * 输出参数： 无
 * 返 回 值： error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func RandNum(length int) string {

	return string(RandomCreateBytes(length, numOnly))
}

//获取指定长度的随机数字和字母串
func RandAlpha(length int) string {

	return string(RandomCreateBytes(length, numAlpha))
}

//获取指定长度的随机字母串
func RandLetter(length int) string {

	return string(RandomCreateBytes(length, numLetter))
}

//获取指定长度的随机大写字母串
func RandUpper(length int) string {

	return string(RandomCreateBytes(length, numUpper))
}

//获取指定长度的随机小写字母串
func RandLower(length int) string {

	return string(RandomCreateBytes(length, numLower))
}

/**********************************************************************
 * 功能描述： 根据输入的数据随机指定长度内容
 * 输入参数： length-随机数字串的长度 alphabets-随机原始数据
 * 输出参数： 无
 * 返 回 值： []byte结果
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func RandomCreateBytes(length int, alphabets []byte) []byte {
	if length <= 0 {
		return []byte("")
	}

	if len(alphabets) == 0 {
		alphabets = numAlpha
	}
	var bytes = make([]byte, length)
	var randBy bool
	if num, err := rand.Read(bytes); num != length || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}
