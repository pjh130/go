package tcp_ser

import (
	"time"
)

type Who interface {
	Id() int64
	Token() string
	Addr() string
	Time() time.Time
}
