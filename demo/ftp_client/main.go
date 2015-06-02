package main

import "fmt"
import "os"
import "io/ioutil"

func main() {
	// new ftp
	ftp := new(FTP)
	// set debug, default false
	ftp.Debug = true
	// connect
	ftp.Connect("10.0.0.0", 20)
	// login
	ftp.Login("smallfish", "****")
	// login failure
	if ftp.Code == 530 {
		fmt.Println("error: login failure")
		os.Exit(-1)
	}
	// pwd
	ftp.Pwd()
	fmt.Println("code:", ftp.Code, ", message:", ftp.Message)
	// mkdir new dir
	ftp.Mkd("/smallfish")
	// stor new file
	ftp.Request("TYPE I")
	b, _ := ioutil.ReadFile("/path/a.txt")
	ftp.Stor("/smallfish/a.txt", b)
	// quit
	ftp.Quit()
}
