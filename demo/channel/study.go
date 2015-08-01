package main

/*
在main goroutine如果没有在写入的channel中close，就不能用rangeq去读取数据
在main goroutine线，期望从管道中获得一个数据，而这个数据必须是其他goroutine线放入管道的
但是其他goroutine线都已经执行完了(all goroutines are asleep)，那么就永远不会有数据放入管道。
所以，main goroutine线在等一个永远不会来的数据，那整个程序就永远等下去了。
这显然是没有结果的，所以这个程序就说“算了吧，不坚持了，我自己自杀掉，报一个错给代码作者，我被deadlock了
*/
/*
关闭2次
ch := make(chan bool)
close(ch)
close(ch)  // 这样会panic的，channel不能close两次

读取的时候channel提前关闭了
ch := make(chan string)
close(ch)
i := <- ch // 不会panic, i读取到的值是空 "",  如果channel是bool的，那么读取到的是false

向已经关闭的channel写数据
ch := make(chan string)
close(ch)
ch <- "good" // 会panic的
判断channel是否close
i, ok := <- ch
if ok {
    println(i)
} else {
    println("channel closed")
}

for循环读取channel
for i := range ch { // ch关闭时，for循环会自动结束
    println(i)
}
防止读取超时
select {
    case <- time.After(time.Second*2):
        println("read channel timeout")
    case i := <- ch:
        println(i)
}

防止写入超时
// 其实和读取超时很像
select {
    case <- time.After(time.Second *2):
        println("write channel timeout")
    case ch <- "hello":
        println("write ok")
}
*/
