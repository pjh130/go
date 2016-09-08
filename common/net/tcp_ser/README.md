使用方法参看example中的例子
	调用函数func StartServer(path string, parser Parser, toDo ToDoFunc)
	1、配置文件路径

	2、实现消息的解析接口
		type Parser interface {
			Decode(conn net.Conn) ([]byte, error)
			Encode(data []byte) ([]byte, error)
		}
	3、实现业务逻辑函数type ToDoFunc func(MsgResquest) []MsgResponse

