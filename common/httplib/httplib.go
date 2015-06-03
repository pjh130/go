package httplib

import (
	"crypto/tls"
	"fmt"
	HttpXie "github.com/astaxie/beego/httplib"
	"io/ioutil"
)

/*
httplib 包里面支持如下的方法返回 request 对象：
Get(url string)
Post(url string)
Put(url string)
Delete(url string)
Head(url string)


*/
func example1() {

	req := HttpXie.Get("http://beego.me/")

	//如果请求的网站是 HTTPS 的，那么我们就需要设置 client 的 TLS 信息，如下所示：
	//关于如何设置这些信息请访问： http://gowalker.org/crypto/tls#Config
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	//设置 header 信息
	req.Header("Accept-Encoding", "gzip,deflate,sdch")
	req.Header("Host", "beego.me")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")

	//支持超时设置
	req.SetTimeout(30, 30)

	//设置请求参数
	req.Param("username", "astaxie")
	req.Param("password", "123456")

	//发送大片的数据
	bt, err := ioutil.ReadFile("hello.txt")
	if err != nil {

	}
	req.Body(bt)

	//httplib支持文件直接上传接口
	b := HttpXie.Post("http://beego.me/")
	b.Param("username", "astaxie")
	b.Param("password", "123456")
	b.PostFile("uploadfile1", "httplib.pdf")
	b.PostFile("uploadfile2", "httplib.txt")

	//返回 Response 对象，req.Response() 方法
	//返回 bytes,req.Bytes() 方法
	//返回 string，req.String() 方法
	//保存为文件，req.ToFile(filename) 方法
	//解析为 JSON 结构，req.ToJson(&result) 方法
	//解析为 XML 结构，req.ToXml(&result) 方法
	str, err := req.String()
	if err != nil {
		return
	}
	fmt.Println(str)
}
