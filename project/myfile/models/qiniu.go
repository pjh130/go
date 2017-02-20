package models

import (
	"log"

	"github.com/pjh130/go/project/myfile/utils"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

type PutRet struct {
	Hash string `json:"hash"` // 如果 uptoken 没有指定 ReturnBody，那么返回值是标准的 PutRet 结构
	Key  string `json:"key"`
}

func init() {

}

func FileUpload(filepath string) error {
	//初始化AK，SK
	conf.ACCESS_KEY = utils.Config.Qiniu["Access"]
	conf.SECRET_KEY = utils.Config.Qiniu["Secret"]

	//创建一个Client
	c := kodo.New(0, nil)

	//设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: utils.Config.Qiniu["Bucket"],
		//设置Token过期时间
		Expires: 3600,
	}
	//生成一个上传token
	token := c.MakeUptoken(policy)

	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet

	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)
	//打印返回的信息
	log.Println(ret)
	//打印出错信息
	if res != nil {
		log.Println("io.Put failed:", res)
		return res
	}
	return nil
}

func FileOverwrite(key string, filepath string) error {
	//初始化AK，SK
	conf.ACCESS_KEY = utils.Config.Qiniu["Access"]
	conf.SECRET_KEY = utils.Config.Qiniu["Secret"]

	//创建一个Client
	c := kodo.New(0, nil)

	//设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: utils.Config.Qiniu["Bucket"] + ":" + key,
		//设置Token过期时间
		Expires: 3600,
	}
	//生成一个上传token
	token := c.MakeUptoken(policy)

	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	//设置上传文件的路径
	//调用PutFile方式上传，这里的key需要和上传指定的key一致
	res := uploader.PutFile(nil, &ret, token, key, filepath, nil)
	//打印返回的信息
	log.Println(ret)
	//打印出错信息
	if res != nil {
		log.Println("io.Put failed:", res)
		return res
	}

	return nil
}

func FileRename(name1, name2 string) error {
	return nil
}

func FileDelete(key string) error {
	//初始化AK，SK
	conf.ACCESS_KEY = utils.Config.Qiniu["Access"]
	conf.SECRET_KEY = utils.Config.Qiniu["Secret"]

	//new一个Bucket管理对象
	c := kodo.New(0, nil)
	p := c.Bucket(utils.Config.Qiniu["Bucket"])

	//调用Delete方法删除文件
	res := p.Delete(nil, key)
	//打印返回值以及出错信息
	if res == nil {
		log.Println("Delete success")
	} else {
		log.Println(res)
		return res
	}

	return nil
}
