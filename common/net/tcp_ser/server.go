package tcp_ser

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/pjh130/go/common/uuid"
)

/*
服务端结构体
*/
type Server struct {
	listener   net.Listener // 服务端监听器 监听xx端口
	maxClient  int          //最大连接数
	currClient int          //当前连接数
	clients    ClientTable  // 客户端列表 抽象出来单独维护和入参 更方便管理连接
	quit       chan *Client // 连接退出嗅探器 触发连接退出处理方法
	lock       sync.Mutex   //互斥
	parser     Parser       //必须实现的消息解析数据流的接口
	toDo       ToDoFunc     //消息处理函数
}

func StartServer(path string, parser Parser, toDo ToDoFunc) {
	if nil == parser {
		panic("[ Parser ] 数据包的解析接口必须实现")
		return
	}

	if nil == toDo {
		panic("[ ToDoFunc ] 处理接收到消息的业务逻辑接口必须实现")
		return
	}

	config, err := ReadConfig(path)
	if nil != err {
		str := fmt.Sprintf("加载配置文件[%s]出错: %s", path, err)
		panic(str)
		return
	}

	log.Println("服务端启动中...")
	//初始化服务端
	server := &Server{
		clients:    make(ClientTable),
		maxClient:  config.MaxClients,
		currClient: 0,
		parser:     parser,
		toDo:       toDo,
	}

	// 设置监听地址及端口
	addr := fmt.Sprintf("0.0.0.0:%d", config.Port)
	listener, err := net.Listen("tcp", addr)
	if nil == err {
		server.listener = listener
		log.Printf("开始监听服务器端口[%d]\n", config.Port)
	} else {
		log.Printf("监听[%d]端口失败:%s\n", config.Port, err)
		return
	}

	//开始工作
	server.Working()

	//启动监听端口
	var tempDelay time.Duration
	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Println("accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return
		}
		tempDelay = 0

		server.NewClient(conn)
	}
}

func (this *Server) GetClient(key string) *Client {
	this.lock.Lock()
	defer this.lock.Unlock()

	return this.clients[key]
}

func (this *Server) NewClient(conn net.Conn) {
	//获取UUID作为客户端的key
	key := uuid.NewV4().String()

	//创建一个客户端
	client := CreateClient(key, conn, this.parser)

	log.Printf("新客户端[%s][%s]，当前连接数[%d]最大连接数[%d]", client.Key, conn.RemoteAddr(), this.currClient, this.maxClient)

	//判断服务的最大客户端数量是否溢出
	if this.maxClient != -1 && this.currClient >= this.maxClient {
		res := MsgResponse{
			Key:  client.Key,
			Data: []byte("More than max connection!"),
		}
		client.PutOut(res)
		client.Close()
		return
	}

	//保存客户端
	this.lock.Lock()
	this.clients[key] = client
	this.currClient++
	this.lock.Unlock()

	//开启协程不断地处理和客户端的事件交互(处理业务逻辑)
	go func() {
		for {
			select {
			//处理接受到的消息.......................................
			case req := <-client.In:
				log.Println(string(req.Data))
				this.toDo(this, req)
			//客户端退出
			case quit := <-client.Quit:
				//调用客户端关闭方法
				quit.Close()
				log.Printf("客户端[%s]退出\n", quit.Key)
				this.lock.Lock()
				delete(this.clients, quit.Key)
				this.currClient--
				this.lock.Unlock()
			}

		}
	}()
}

/*

 */
func (this *Server) Working() {
	go func() {
		for {
			select {
			// 退出了一个连接
			case client := <-this.quit:
				// 调用客户端关闭方法
				client.Close()
				delete(this.clients, client.Key)
			}
		}
	}()
}
