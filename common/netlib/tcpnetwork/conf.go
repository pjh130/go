package tcpnetwork

import (
	"time"
)

//数据定义
var (
	MaxMsgLen       uint32 = 4096
	MinMsgLen       uint32 = 1
	Timeout                = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = false
	MaxConnNum             = 1024
	PendingWriteNum        = 1024
	ServerAddr             = "127.0.0.1:8888"
)