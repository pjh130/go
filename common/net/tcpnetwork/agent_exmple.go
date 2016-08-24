package tcpnetwork

import (
	"log"
)

//写一个实例
func MyAgentFunc(conn *TCPConn) Agent {
	a := &MyAgent{
		Key:     conn.Key,
		Connect: conn,
	}
	return a
}

type MyAgent struct {
	Key      string
	Connect  Conn
	UserData interface{}
}

func (a *MyAgent) Run() {
	//开启业务逻辑处理
	for {
		data, err := a.Connect.ReadMsg()
		if err != nil {
			log.Println("read message: %v", err)
			return
		}

		log.Println("receive: ", string(data))
		a.Connect.WriteMsg([]byte("Goodbye 1!"))
		a.Connect.WriteMsgOther(a.Key, []byte("Goodbye 2!"))
	}
}

func (a *MyAgent) OnClose() {
	log.Println("agent close ")
}

func (a *MyAgent) read() {

}
