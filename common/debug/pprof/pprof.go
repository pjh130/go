package pprof

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

/**********************************************************************
 * 功能描述： 在运行中实时监控调优
 * 输入参数： webaddr-地址和端口号
 * 输出参数： 无
 * 返 回 值： error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151215           V1.0            panpan            创建
 ************************************************************************/
func StartAdminHttp(webaddr string) error {
	adminServeMux := http.NewServeMux()
	adminServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	adminServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	adminServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	adminServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	adminServeMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	err := http.ListenAndServe(webaddr, adminServeMux)
	if err != nil {
		x := fmt.Sprintf("http.ListenAdServe(\"%s\") failed (%s)", webaddr, err.Error())
		fmt.Println(x)
		return err
	}
	return nil
}
