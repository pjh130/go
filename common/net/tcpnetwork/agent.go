package tcpnetwork

type Agent interface {
	Run()
	OnClose()
}
