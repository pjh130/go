package main

//fmt 包实现了格式化 I/O 函数，类似于 C 的 printf 和 scanf。格式“占位符”衍生自 C，但比 C 更简单。
//【打印】
//占位符：
//[一般]
//%v	相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
//%#v	相应值的 Go 语法表示
//%T	相应值的类型的 Go 语法表示
//%%	字面上的百分号，并非值的占位符
//[布尔]
//%t	单词 true 或 false。
//[整数]
//%b	二进制表示
//%c	相应 Unicode 码点所表示的字符
//%d	十进制表示
//%o	八进制表示
//%q	单引号围绕的字符字面值，由 Go 语法安全地转义
//%x	十六进制表示，字母形式为小写 a-f
//%X	十六进制表示，字母形式为大写 A-F
//%U	Unicode 格式：U+1234，等同于 "U+%04X"
//[浮点数及其复合构成]
//%b	无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78
//%e	科学计数法，例如 -1234.456e+78
//%E	科学计数法，例如 -1234.456E+78
//%f	有小数点而无指数，例如 123.456
//%g	根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的 0）输出
//%G	根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的 0）输出
//[字符串与字节切片]
//%s	字符串或切片的无解译字节
//%q	双引号围绕的字符串，由 Go 语法安全地转义
//%x	十六进制，小写字母，每字节两个字符
//%X	十六进制，大写字母，每字节两个字符
//[指针]
//%p	十六进制表示，前缀 0x
//[注意]
//这里没有 'u' 标记。若整数为无符号类型，他们就会被打印成无符号的。类似地， 这里也不需要指定操作数的大小（int8，int64）。
//宽度与精度的控制格式以 Unicode 码点为单位。（这点与 C 的 printf 不同， 它以字节数为单位。）
//二者或其中之一均可用字符 '*' 表示， 此时它们的值会从下一个操作数中获取，该操作数的类型必须为 int。
//// 宽度与精度的控制以 Unicode 码点为单位
//fmt.Printf("\"%8s\"\n", "123456") // 最大长度为 8
//// "  123456"
//fmt.Printf("\"%8s\"\n", "你好")   // 最大长度为 8
//// "      你好"
//// 宽度与精度均可用字符 '*' 表示
//fmt.Printf("%0*.*f \n", 8, 3, 13.25) // 总长度 8，小数位数 3
//fmt.Printf("%08.3f \n", 13.25)       // 总长度 8，小数位数 3
//// 0013.250
//对数值而言，宽度为该数值占用区域的最小宽度；精度为小数点之后的位数。 但对于 %g/%G 而言，
//精度为所有数字的总数。例如，对于 123.45，格式 %6.2f 会打印 123.45，而 %.4g 会打印 123.5。%e
//和 %f 的默认精度为 6；但对于 %g 而言，它的默认精度为确定该值所必须的最小位数。
//对大多数值而言，宽度为输出的最小字符数，如果必要的话会为已格式化的形式填充空格。对字符串而言，
//精度为输出的最大字符数，如果必要的话会直接截断。
//// 宽度与精度标记字符串
//fmt.Printf("%8q", "ABC")         // 最小长度为 8（包括 %q 的引号字符）
////    "ABC"
//fmt.Printf("%.8q", "1234567890") // 最大长度为 8（不包括 %q 的引号字符）
//// "12345678"
//[其它标记]
//+	总打印数值的正负号；对于 %q（%+q）保证只输出 ASCII 编码的字符。
//-	在右侧而非左侧填充空格（左对齐该区域）
//#	备用格式：为八进制添加前导 0（%#o），为十六进制添加前导 0x（%#x）或
//0X（%#X），为 %p（%#p）去掉前导 0x；如果可能的话，%q（%#q）会打印原始（即反引号围绕的）
//字符串；如果是可打印字符，%U（%#U）会写出该字符的 Unicode 编码形式（如字符 x 会被打印成 U+0078 'x'）。
//' '	（空格）为数值中省略的正负号留出空白（% d）；以十六进制（% x, % X）打印字符串或切片时，
//在字节之间用空格隔开：fmt.Printf("% x\n", "Hello")
//// 48 65 6c 6c 6f
//0	填充前导的 0 而非空格；对于数字，这会将填充移到正负号之后
//[注意]
//标记有时会被占位符忽略，所以不要指望它们。例如十进制没有备用格式，因此 %#d 与 %d 的行为相同。
//对于每一个 Printf 类的函数，都有一个 Print 函数，该函数不接受任何格式化， 它等价于对每一个
//操作数都应用 %v。另一个变参函数 Println 会在操作数之间插入空白， 并在末尾追加一个换行符。
//不考虑占位符的话，如果操作数是接口值，就会使用其内部的具体值，而非接口本身。 因此：
//var i interface{} = 23
//fmt.Printf("%v\n", i)
//// 会打印 23
//若一个操作数实现了 Formatter 接口，该接口就能更好地用于控制格式化。
//若其格式（它对于 Println 等函数是隐式的 %v）对于字符串是有效的（%s %q %v %x %X），以下两条规则也适用：
//1。若一个操作数实现了 error 接口，Error 方法就能将该对象转换为字符串，随后会根据占位符的需要进行格式化。
//2。若一个操作数实现了 String() string 方法，该方法能将该对象转换为字符串，随后会根据占位符的需要进行格式化。
//为避免以下这类递归的情况：
//type X string
//func (x X) String() string { return Sprintf("<%s>", x) }
//需要在递归前转换该值：
//func (x X) String() string { return Sprintf("<%s>", string(x)) }
//[格式化错误]
//如果给占位符提供了无效的实参（例如将一个字符串提供给 %d），所生成的字符串会包含该问题的描述，如下例所示：
//类型错误或占位符未知：%!verb(type=value)
//Printf("%d", hi)
//// %!d(string=hi)
//实参太多：%!(EXTRA type=value)
//Printf("hi", "guys")
//// hi%!(EXTRA string=guys)
//实参太少：%!verb(MISSING)
//Printf("hi%d")
//// hi %!d(MISSING)
//宽度或精度不是 int 类型：%!(BADWIDTH）或 %!(BADPREC)
//Printf("%*s", 4.5, "hi")
//// %!(BADWIDTH)hi
//Printf("%.*s", 4.5, "hi")
//// %!(BADPREC)hi
//所有错误都始于“%!”，有时紧跟着单个字符（占位符），并以小括号括住的描述结尾。
//【扫描】
//一组类似的函数通过扫描已格式化的文本来产生值。Scan、Scanf 和 Scanln 从 os.Stdin 中读取；
//Fscan、Fscanf 和 Fscanln 从指定的 io.Reader中读取； Sscan、Sscanf 和 Sscanln 从实参
//字符串中读取。Scanln、Fscanln 和 Sscanln 在换行符处停止扫描，且需要条目紧随换行符之后；
//Scanf、Fscanf 和 Sscanf 需要输入换行符来匹配格式中的换行符；其它函数则将换行符视为空格。
//Scanf、Fscanf 和 Sscanf 根据格式字符串解析实参，类似于 Printf。例如，%x 会将一个整数扫
//描为十六进制数，而 %v 则会扫描该值的默认表现格式。格式化行为类似于 Printf，但也有如下例外：
//%p 没有实现
//%T 没有实现
//%e %E %f %F %g %G 都完全等价，且可扫描任何浮点数或复合数值
//%s 和 %v 在扫描字符串时会将其中的空格作为分隔符
//标记 # 和 + 没有实现
//在使用 %v 占位符扫描整数时，可接受友好的进制前缀 0（八进制）和 0x（十六进制）。
//宽度被解释为输入的文本（%5s 意为最多从输入中读取 5 个符文来扫描成字符串），而扫描函数则没有
//精度的语法（没有 %5.2f，只有 %5f）。
//当以某种格式进行扫描时，无论在格式中还是在输入中，所有非空的连续空白字符 （除换行符外）都等
//价于单个空格。由于这种限制，格式字符串文本必须匹配输入的文本，如果不匹配，扫描过程就会停止，并返回已扫描的实参数。
//在所有的扫描参数中，若一个操作数实现了 Scan 方法（即它实现了 Scanner 接口），该操作数将使用
//该方法扫描其文本。此外，若已扫描的实参数少于所提供的实参数，就会返回一个错误。
//所有需要被扫描的实参都必须是基本类型或实现了 Scanner 接口的类型。
//注意：Fscan 等函数会从输入中多读取一个字符（符文），因此，如果循环调用扫描函数，可能会跳过
//输入中的某些数据。一般只有在输入的数据中没有空白符时该问题才会出现。若提供给 Fscan 的读取器
//实现了 ReadRune，就会用该方法读取字符。若此读取器还实现了 UnreadRune 方法，就会用该方法保
//存字符，而连续的调用将不会丢失数据。若要为没有 ReadRune 和 UnreadRune 方法的读取器加上这些功能，需使用 bufio.NewReader。