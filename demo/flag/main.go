package main

import (
	"flag"
	"fmt"
)

func main() {
	//	flag.Usage()

	Test1()

	Test2()
}

/*
几点注意事项：
1，通过flag.String(), Bool(), Int()等方式来定义命令行中需要使用的flag。
2，在定义完flag后，通过调用flag.Parse()来进行对命令行参数的解析。
3，命令行参数的格式可以是：
-flag // 代表bool值，相当于-flag=true
-flag xxx （使用空格，一个 - 符号）// non-boolean flags only  不支持bool值标志
--flag xxx （使用空格，两个 - 符号）
-flag=xxx （使用等号，一个 - 符号）
--flag=xxx （使用等号，两个 - 符号）
其中，布尔类型的参数防止解析时的二义性，应该使用等号的方式指定。
*/
func Test1() {
	data_path := flag.String("D", "/home/manu/sample/", "DB data path")
	log_file := flag.String("l", "/home/manu/sample.log", "log file")
	nowait_flag := flag.Bool("W", false, "do not wait until operation completes")

	flag.Parse()

	var cmd string = flag.Arg(0)

	//	fmt.Println(data_path)
	fmt.Printf("action   : %s\n", cmd)
	fmt.Printf("data path: %s\n", *data_path)
	fmt.Printf("log file : %s\n", *log_file)
	fmt.Printf("nowait     : %v\n", *nowait_flag)

	fmt.Printf("-------------------------------------------------------\n")

	fmt.Printf("there are %d non-flag input param\n", flag.NArg())
	for i, param := range flag.Args() {
		fmt.Printf("#%d    :%s\n", i, param)
	}

}

func Test2() {
	var Input_pstrName = flag.String("name", "gerry", "input ur name")
	var Input_piAge = flag.Int("age", 20, "input ur age")
	var Input_flagvar int

	flag.IntVar(&Input_flagvar, "flagname", 1234, "help message for flagname")

	flag.Parse()

	// After parsing, the arguments after the flag are available as the slice flag.Args() or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg()-1
	// Args returns the non-flag command-line arguments
	// NArg is the number of arguments remaining after flags have been processed
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	fmt.Println("name=", *Input_pstrName)
	fmt.Println("age=", *Input_piAge)
	fmt.Println("flagname=", Input_flagvar)

	/*
		output:
		mba:filter gerryyang$ ./flag --name "aaa" -age=123 -flagname=0x11 para1 para2 para3
		args=[para1 para2 para3], num=3
		arg[0]=para1
		arg[1]=para2
		arg[2]=para3
		name= aaa
		age= 123
		flagname= 17
	*/
}
