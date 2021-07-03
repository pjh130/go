golang

###GO网站
* [【必看】Awesome Go收集了 Go 语言的流行库，框架和软件](https://github.com/avelino/awesome-go)
* [Go必看学习教程](http://www.topgoer.com/)
* [Go技术文章精选by韩亚军](https://hanyajun.com/golang/go_article_2019/)
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
* [Golang优秀开源项目汇总](https://github.com/hackstoic/golang-open-source-projects/blob/master/golang-open-source-projects.md)
* [Golang筛选过的优秀开源项目汇总](https://github.com/hackstoic/golang-open-source-projects)
* [谷歌官方维护了一个基于go语言的开源项目列表](https://github.com/golang/go/wiki/Projects)
* [Awesome Go收集了 Go 语言的流行库，框架和软件](https://github.com/avelino/awesome-go)
* [OPEN经验库](http://www.open-open.com/lib/view/open1396063913278.html#Configuration_File_Parsers)
* [开源社区](http://www.oschina.net/project/lang/358/go)
* [为互联网IT人打造的中文版awesome-go](https://github.com/hackstoic/golang-open-source-projects)
* [有哪些值得学习的 Go 语言开源项目](https://www.zhihu.com/question/20801814)
* [Go最新资料汇总链接](https://github.com/ty4z2008/Qix/blob/master/golang.md)

###Go 
* [Go 语言包管理](https://gopm.io/)
* [Search for Go Packages](https://godoc.org/)

###================基础类库================
* [开源的基础类库](https://github.com/dropbox/godropbox)

###================golang协程池================
* [Package pool implements a limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation](https://github.com/go-playground/pool)
* [ants是一个受fasthttp启发的高性能协程池](https://github.com/panjf2000/ants)

###================语言编码转换================
* [character-set conversion library implemented in Go](https://github.com/axgle/mahonia)

###================测试工具================
* [http性能测试](https://github.com/rakyll/boom)
* [Testify - Thou Shalt Write Tests](https://github.com/stretchr/testify)

###================文件系统================
* [分布式文件系统 SeaweedFS](https://github.com/chrislusf/seaweedfs)
* [Syncthing: 一个在计算机之间同步文件/文件夹的私密安全同步工具](https://github.com/syncthing/syncthing)
* [分布式文件系统 SeaweedFS](https://github.com/chrislusf/seaweedfs)
* [分布式文件系统 IPFS](https://github.com/ipfs/go-ipfs)
* [http的文件服务器 gohttp](https://github.com/codeskyblue/gohttp)
* [gocryptfs 是一个使用 Go 开发的，加密的覆盖文件系统](https://github.com/rfjakob/gocryptfs)
* [bfs 是使用 Go 编写的分布式文件系统（小文件存储）](https://github.com/Terry-Mao/bfs)
* [Go 实现的跨平台文件系统监控库](https://github.com/fsnotify/fsnotify)

###================WEB================
* [Go静态网站生成器Hugo](https://github.com/spf13/hugo)
* [Negroni 是一个很地道的 web 中间件](https://github.com/urfave/negroni)

###===============消息推送================
* [gopush-cluster是一套golang开发的实时消息推送集群](https://github.com/Terry-Mao/gopush-cluster)
* [开源消息系统 NSQ](https://github.com/nsqio/nsq)
* [nats是一个开源的，云原生的消息系统](https://github.com/nats-io/nats-server)

###================文件监控================
* [File system notifications for Go](https://github.com/howeyc/fsnotify)

###================游戏框架================
* [游戏服务器框架 gonet](http://gonet2.github.io/)
* [Leaf 游戏服务器框架](https://github.com/name5566/leaf)
* [高效的跨平台服务器网络库](https://github.com/davyxu/cellnet)
* [mqant是简洁高效高性能的分布式游戏服务器框架](https://github.com/liangdas/mqant)
* [xingo高性能网络库，游戏开发脚手架](https://github.com/viphxin/xingo)
* [nano - 重量轻，设备，高性能的基于 golang 游戏服务器架构](https://github.com/lonng/nano)
* [goworld- 可扩展的游戏服务器引擎，具有空间实体框架和热插拔功能](https://github.com/xiaonanln/goworld)

###================爬虫，下载================
* [DHT实现了BitTorrent DHT协议](https://github.com/shiyanhui/dht)
* [Go 爬虫软件 Pholcus](https://github.com/henrylee2cn/pholcus)
* [分布式爬虫 ants](https://github.com/wcong/ants-go)
* [爬虫框架 go_spider](https://github.com/hu17889/go_spider)
* [基于docker的分布式爬虫服务](https://github.com/huichen/zerg)

###================任务系统================
* [轻量级异步定时任务系统 kingtask](https://github.com/kingsoft-wps/kingtask)
* [Go 异步任务队列 Go Machinery](https://github.com/RichardKnop/machinery)

###================词库、搜索引擎================
* [Go 分词库 GoJieba](https://github.com/yanyiwu/gojieba)
* [Go中文分词](https://github.com/huichen/sego)
* [悟空全文搜索引擎](https://github.com/huichen/wukong)
* [现代化的文本索引库，可以做多种形式的索引及搜索查询,对于文档内容索引应该是非常好用的](https://github.com/blevesearch/bleve)
* [GoLucene 是 Java 的 Lucene 搜索引擎的 Go 语言移植版本](https://github.com/balzaczyy/golucene)
* ["结巴"中文分词的Golang语言版本](https://github.com/yanyiwu/gojieba)

###================Reader================
* [开源的 Google Reader 替代品 GoRead](https://github.com/mjibson/goread)

###================网络库================
* [Gorilla WebSocket is a Go implementation of the WebSocket protocol](https://github.com/gorilla/websocket)
* [go 的简单网络框架 kendynet-go](https://github.com/sniperHW/kendynet)
* [Teleport是一款适用于分布式系统的高并发API框架](https://github.com/henrylee2cn/teleport)
* [qTunnel使用的安全套接字隧道](https://github.com/getqujing/qtunnel)
* [gnet 是一个高性能、轻量级、非阻塞的事件驱动 Go 网络框架](https://github.com/panjf2000/gnet)
* [cellnet是一个组件化、高扩展性、高性能的开源服务器网络库](https://github.com/davyxu/cellnet)

###================WEB框架================
* [fasthttp号称是比go原生的net/http快10倍](https://github.com/valyala/fasthttp)
* [高性能分支从httprouter第一个适合的路由器 fasthttp](https://github.com/buaazp/fasthttprouter)
* [Echo是个快速的HTTP路由器和微型Web框架](https://github.com/labstack/echo)
* [Gin 它具有类似于 martini 的 API，性能更高的Web框架](https://github.com/gin-gonic/gin)
* [Revel是一个高生产力的Go语言Web框架](https://github.com/revel/revel)
* [beego 是一种用于 Go 编程语言的开源高性能 Web 框架](https://github.com/astaxie/beego)

###================json================
* [ffjson: faster JSON for Go](https://github.com/pquerna/ffjson)
* [Jason is an easy-to-use JSON library for Go](https://github.com/antonholmquist/jason)
* [gojson attempts to generate go struct definitions from json documents](https://github.com/ChimeraCoder/gojson)
* [Go-Json-Rest](https://github.com/ant0ine/go-json-rest)
* [easyjson](https://github.com/mailru/easyjson)
* [go-simplejson](https://github.com/bitly/go-simplejson)
* [jsoniter （json-iterator）是一款快且灵活的 JSON 解析器](https://github.com/json-iterator/go)

###================算法库================
* [大量算法库](https://github.com/henrylee2cn/algorithm)

###================存储================
* [etcd 是一个高可用的 Key/Value 存储系统，主要用于分享配置和服务发现。](https://github.com/etcd-io/etcd)
* [cache2go - key/value 内存缓存，支持基于超时的自动无效功能](https://github.com/muesli/cache2go)
* [Groupcache 是一个缓存和缓存填充库，在许多情况下用于替代 memcached](https://github.com/golang/groupcache)
* [GeeCache：模仿 groupcache 实现的分布式缓存](https://geektutu.com/post/geecache.html)
* [Yoke是Postgres的高可用集群，具有自动切换和自动集群恢复](https://github.com/nanopack/yoke)
* [KiteQ基于go+protobuff实现的多种持久化方案的mq框架](https://github.com/blackbeans/kiteq)
* [cockroach新型的分布式SQL数据库](https://github.com/cockroachdb/cockroach)

###================RPC================
* [RobustIRC 是不会有网络中断情况的 IRC](https://github.com/robustirc/robustirc)

###================监控服务================
* [pingd 是世界上最简单的监控服务](https://github.com/pinggg/pingd)

###================图形框架================
* [GoQt是一个Go语言的GUI工具包](https://github.com/visualfc/goqt)
* [Golang Desktop Automation.键盘鼠标屏幕位图事件窗口](https://github.com/go-vgo/robotgo)

###================excel================
* [Excelize ](https://github.com/Luxurioust/excelize)
* [Go (golang) library for reading and writing XLSX files](https://github.com/tealeg/xlsx)

###================IM框架================
* [分布式可伸缩 IM 服务器 FishChat](https://github.com/oikomi/FishChatServer)
* [Go 开发的 IM 和推送服务 goim](https://github.com/Terry-Mao/goim)
* [Mattermost 采用 Go 语言开发，这是一个开源的团队通讯服务](https://github.com/mattermost/platform)
* [GoBelieve IM云平台服务端](https://github.com/GoBelieveIO/im_service)

###================电商系统================
* [基于Go语言开发的开源电商系统](https://github.com/qor/qor)
* [Go2o是Google Go语言结合领域驱动设计DDD的开源O2O实现](https://github.com/jsix/go2o)

###================图像处理================
* [Gift 包提供一整套有用的图像处理过滤器](https://github.com/disintegration/gift)

###================蓝牙================
* [Gatt 是一个 Go 语言包，用来构建低功耗蓝牙外设](https://github.com/paypal/gatt)

###================验证码服务================
* [Golang 实现的验证码服务](https://github.com/jianxinio/captcha)
* [go语言验证码服务器](https://github.com/hanguofeng/gocaptcha)

###================网络代理================
* [goproxy库编写的http代理服务器+图片cache保存脚本](https://github.com/elazarl/goproxy)

###================分布式系统================
* [Consul是HashiCorp公司推出的开源工具，用于实现分布式系统的服务发现与配置](https://github.com/hashicorp/consul)
* [etcd是一个高可用的键值存储系统，主要用于共享配置和服务发现](https://github.com/coreos/etcd)


* [go 实现直播服务](https://github.com/qieangel2013/livego)

###================go 绘制图表================
* [plot 是一个用 Go 语言实现的绘图库](https://github.com/gonum/plot)