package tcpnetwork

import (
	"log"
)

type Agent interface {
	Run()
	OnClose()
}

//写一个实例
type MyAgent struct {
	Connect     Conn
	UserData interface{}
}

func (a *MyAgent) Run() {
	for {
		data, err := a.Connect.ReadMsg()
		if err != nil {
			log.Println("read message: %v", err)
			break
		}
		
		log.Println("message: ", string(data))
		
	}
}

func (a *MyAgent) OnClose() {
	log.Println("agent close ")
}