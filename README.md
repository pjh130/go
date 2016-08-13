golang

###GO网站
* [GO语言中文网](http://studygolang.com/)
* [Golang中国社区](http://www.golangtc.com/)

###go程序变小
* go build -ldflags "-s -w"  (go install类似)
* 解释一下参数的意思：
* -ldflags： 表示将后面的参数传给连接器（5/6/8l）
* -s：去掉符号信息（然后panic时候的stack trace就没有任何文件名/行号信息了，这个等价于普通C/C++程序被strip的效果）
* -w：去掉DWARF调试信息。得到的程序就不能用gdb调试了

=========
###================开源项目================
* [谷歌官方维护了一个基于go语言的开源项目列表](https://github.com/golang/go/wiki/Projects)

* [Awesome Go收集了 Go 语言的流行库，框架和软件](https://github.com/avelino/awesome-go)

* [OPEN经验库](http://www.open-open.com/lib/view/open1396063913278.html#Configuration_File_Parsers)

* [开源社区](http://www.oschina.net/project/lang/358/go)


###Go 
* [Go 语言包管理](https://gopm.io/)
* [Search for Go Packages](https://godoc.org/)

###================测试工具================
* [http性能测试](https://github.com/rakyll/boom)


###================文件系统================
* [分布式文件系统 SeaweedFS](https://github.com/chrislusf/seaweedfs)

* [Syncthing: 一个在计算机之间同步文件/文件夹的私密安全同步工具](https://github.com/syncthing/syncthing)

* [分布式文件系统 SeaweedFS](https://github.com/chrislusf/seaweedfs)

* [分布式文件系统 IPFS](https://github.com/ipfs/go-ipfs)

* [http的文件服务器 gohttp](https://github.com/codeskyblue/gohttp)


###================WEB================
* [Go静态网站生成器Hugo] (https://github.com/spf13/hugo)


###===============消息推送================
* [gopush-cluster是一套golang开发的实时消息推送集群](https://github.com/Terry-Mao/gopush-cluster)

* [开源消息系统 NSQ](https://github.com/nsqio/nsq)


###================文件监控================
* [File system notifications for Go](https://github.com/howeyc/fsnotify)


###================游戏框架================
* [游戏服务器框架 gonet](http://gonet2.github.io/)

* [Leaf 游戏服务器框架](https://github.com/name5566/leaf)


###================爬虫================
* [Go 爬虫软件 Pholcus](https://github.com/henrylee2cn/pholcus)

* [分布式爬虫 ants](https://github.com/wcong/ants-go)

* [爬虫框架 go_spider](https://github.com/hu17889/go_spider)

* [基于docker的分布式爬虫服务](https://github.com/huichen/zerg)


###================任务系统================
* [轻量级异步定时任务系统 kingtask](https://github.com/kingsoft-wps/kingtask)

* [Go 异步任务队列 Go Machinery](https://github.com/RichardKnop/machinery)


###================词库================
* [Go 分词库 GoJieba](https://github.com/yanyiwu/gojieba)
* [悟空全文搜索引擎](https://github.com/huichen/wukong)
* [现代化的文本索引库，可以做多种形式的索引及搜索查询,对于文档内容索引应该是非常好用的](https://github.com/blevesearch/bleve)
* [Go中文分词](https://github.com/huichen/sego)


###================Reader================
* [开源的 Google Reader 替代品 GoRead](https://github.com/mjibson/goread)


###================网络库================
* [go 的简单网络框架 kendynet-go](https://github.com/sniperHW/kendynet-go)
* [Teleport是一款适用于分布式系统的高并发API框架](https://github.com/henrylee2cn/teleport)
* [qTunnel使用的安全套接字隧道](https://github.com/getqujing/qtunnel)
* [Fast HTTP implementation for Go](https://github.com/valyala/fasthttp)


###================json================
* [ffjson: faster JSON for Go](https://github.com/pquerna/ffjson)
* [Jason is an easy-to-use JSON library for Go](https://github.com/antonholmquist/jason)
* [gojson attempts to generate go struct definitions from json documents](https://github.com/ChimeraCoder/gojson)
* [Go-Json-Rest](https://github.com/ant0ine/go-json-rest)
* [easyjson](https://github.com/mailru/easyjson)

###================算法库================
* [大量算法库](https://github.com/henrylee2cn/algorithm)


###================存储================
* [Yoke是Postgres的高可用集群，具有自动切换和自动集群恢复](https://github.com/nanopack/yoke)
* [KiteQ基于go+protobuff实现的多种持久化方案的mq框架](https://github.com/blackbeans/kiteq)
* [cockroach新型的分布式SQL数据库](https://github.com/cockroachdb/cockroach)


###================RPC================
* [RobustIRC 是不会有网络中断情况的 IRC](https://github.com/robustirc/robustirc)


###================监控服务================
* [pingd 是世界上最简单的监控服务](https://github.com/pinggg/pingd)


###================图形框架================
* [GoQt是一个Go语言的GUI工具包](https://github.com/visualfc/goqt)
