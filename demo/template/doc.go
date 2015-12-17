package main

/*
Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象，这和Java或者C++中的this类
似，如果要访问当前对象的字段通过{{.FieldName}},但是需要注意一点：这个字段必须是导出的(字段首字母必
须是大写的),否则在渲染的时候就会报错

双大括号{{}}是区分模板代码和HTML的分隔符
括号里边可以是要显示输 出的数据,或者是控制语句,比如if判断式或者range循环体等
.|formatter表示对当前这个元素的值以 formatter 方式进行格式化输出
.|urlquery}即表示对当前元素的值进行转换以适合作为URL一部 分
.|html 表示对当前元素的值进行适合用于HTML 显示的字符转化,比如">"会被转义 成"&gt;"
*/
