package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jordan-wright/email"
)

var appPath string

func main() {
	var str string
	args := os.Args
	iNeed := 10
	appPath = filepath.Dir(args[0])
	fmt.Println(appPath)

	if len(args) < iNeed {
		str = fmt.Sprintf("程序启动至少需要%d个参数\n", iNeed)
		writeLog(str)
		<-time.After(time.Second * 10)
		return
	} else {
		fmt.Printf("程序启动参数足够\n")
	}

	var user string = args[1]     //发送邮件的帐户名
	var pwd string = args[2]      //发送邮件的密码
	var fromUser string = args[3] //发送者
	var toUser []string           //发给谁
	var temp string = args[4]
	if strings.Contains(temp, "|") {
		toUser = strings.Split(temp, "|")
	} else {
		toUser = strings.Split(temp, ";")
	}

	var subject string = args[5] //标题
	var files []string           //发送的多个附件
	temp = args[6]
	if strings.Contains(temp, "|") {
		files = strings.Split(temp, "|")
	} else {
		files = strings.Split(temp, ";")
	}

	var addrSMTP string = args[7] //SMTP地址(如：smtp.qq.com:587)
	var host string = args[8]     //邮件服务器地址(如：smtp.qq.com)
	var content string = args[9]  //邮件里边的正文内容

	var bt bytes.Buffer
	// NewEmail返回一个email结构体的指针
	e := email.NewEmail()
	// 发件人
	e.From = fromUser
	// 收件人(可以有多个)
	e.To = toUser
	// 邮件主题
	e.Subject = subject

	// 邮件正文
	bt.WriteString(content)
	e.Text = bt.Bytes()

	// 以路径将文件作为附件添加到邮件中
	for _, f := range files {
		e.AttachFile(f)
	}
	// 发送邮件(如果使用QQ邮箱发送邮件的话，passwd不是邮箱密码而是授权码)
	err := e.Send(addrSMTP, smtp.PlainAuth("", user, pwd, host))

	//把邮件发送结果写入到日志中去
	if err == nil {
		str = fmt.Sprintf("【发送成功】 收件人【%s】附件【%s】", toUser, files)
	} else {
		fmt.Println(err)
		str = fmt.Sprintf("【发送失败】 收件人【%s】附件【%s】 %s", toUser, files, err)
	}

	writeLog(str)
}

func writeLog(str string) {
	f, err := os.OpenFile(appPath+"\\email.log", os.O_CREATE|os.O_APPEND, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		write := time.Now().Format("2006-01-02 15:04:05 ") + str + "\r\n"
		_, err = f.WriteString(write)
	}
}
