package randlib

import (
	"math/rand"
	"strconv"
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
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

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
func RandString(length int) string {
	var v string
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		n := rand.Intn(9)
		v += strconv.Itoa(n)
	}

	return v
}
