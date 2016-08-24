package tcpnetwork

import (
	"log"
	"net"
	"sync"
	"time"

	"github.com/pjh130/go/common/net/tcpnetwork/utils"
)

type AgentFunc func(*TCPConn) Agent

type TCPServer struct {
	Addr            string
	MaxConnNum      int
	PendingWriteNum int
	NewAgent        AgentFunc
	ln              net.Listener
	conns           map[string]Conn
	mutexConns      sync.Mutex
	wgLn            sync.WaitGroup
	wgConns         sync.WaitGroup

	msgParser *MsgParser
}

func CreateServer(agent AgentFunc) *TCPServer {
	ser := &TCPServer{
		NewAgent: agent,
	}

	ser.init()

	return ser
}

func (server *TCPServer) Start() {
	if server.NewAgent == nil {
		log.Fatal("NewAgent must not be nil")
		return
	}

	go server.run()
}

func (server *TCPServer) init() {
	utils.InitConfig()

	//设置初始化参数
	server.Addr = utils.Cfg.ServerAddr
	server.MaxConnNum = utils.Cfg.MaxConnNum
	server.PendingWriteNum = utils.Cfg.PendingWriteNum

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatal("%v", err)
	}

	if server.MaxConnNum <= 0 {
		server.MaxConnNum = utils.Cfg.MaxConnNum
		log.Println("invalid MaxConnNum, reset to %v", server.MaxConnNum)
	}
	if server.PendingWriteNum <= 0 {
		server.PendingWriteNum = utils.Cfg.PendingWriteNum
		log.Println("invalid PendingWriteNum, reset to %v", server.PendingWriteNum)
	}

	server.ln = ln
	server.conns = make(map[string]Conn)

	// msg parser
	msgParser := NewMsgParser()
	server.msgParser = msgParser
}

func (server *TCPServer) run() {
	server.wgLn.Add(1)
	defer server.wgLn.Done()
	log.Println("run", server.Addr)
	var tempDelay time.Duration
	for {
		conn, err := server.ln.Accept()
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
		connTcp := newTCPConn(conn, server.PendingWriteNum, server.msgParser)

		server.mutexConns.Lock()
		if len(server.conns) >= server.MaxConnNum {
			server.mutexConns.Unlock()
			conn.Close()
			log.Println("too many connections")
			continue
		}
		server.conns[connTcp.Key] = connTcp
		server.mutexConns.Unlock()

		server.wgConns.Add(1)

		agent := server.NewAgent(connTcp)

		//监控连接退出
		go func() {
			//监控需要发送消息给其他链接
			go func() {
				select {
				case other := <-connTcp.Other:
					//找出存在的链接
					server.mutexConns.Lock()
					to := server.conns[other.Key]
					server.mutexConns.Unlock()
					if nil != to {
						to.WriteMsg(other.Data)
					}
				}
			}()

			agent.Run()

			//如果业务逻辑结束则需要清理工作
			//1、关闭链接
			connTcp.Close()

			//2、清除链接的缓存列表
			server.mutexConns.Lock()
			delete(server.conns, connTcp.Key)
			server.mutexConns.Unlock()

			//3、业务逻辑自身清理
			agent.OnClose()

			server.wgConns.Done()
		}()
	}
}

func (server *TCPServer) Close() {
	server.ln.Close()
	server.wgLn.Wait()

	server.mutexConns.Lock()
	for key := range server.conns {
		server.conns[key].Close()
	}
	server.conns = nil
	server.mutexConns.Unlock()
	server.wgConns.Wait()
}
