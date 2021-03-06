package models

import (
	"bufio"
	"github.com/PuerkitoBio/goquery"
	"github.com/pjh130/go/demo/forex/utils"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	RES_FILE = "res/money.txt"
)

type MoneyCode struct {
	Country string
	Name    string
	Code    string
}

var CodeList []MoneyCode

func InitMoneyCode() {

	//打开文件，并进行相关处理
	f, err := os.Open(RES_FILE)
	if err != nil {
		log.Println(err)
		return
	}

	//文件关闭
	defer f.Close()

	//将文件作为一个io.Reader对象进行buffered I/O操作
	br := bufio.NewReader(f)
	for {
		//每次读取一行
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else {
			temp := strings.TrimSpace(string(line[0:]))
			if len(temp) > 0 {
				var hzRegexp = regexp.MustCompile(`\s+`)
				items := hzRegexp.Split(temp, -1)
				if len(items) >= 3 {
					var code MoneyCode
					code.Country = strings.TrimSpace(items[0])
					code.Name = strings.TrimSpace(items[1])
					code.Code = strings.TrimSpace(items[2])
					CodeList = append(CodeList, code)
				}
			}
		}
	}

	log.Println("CodeList length: ", len(CodeList))
}

func StartCollect() {
	if false {
		go CollectData()
	} else {
		ticker := time.NewTicker(time.Duration(utils.WAIT_TIME) * time.Minute)
		for {
			select {
			case <-ticker.C:
				go CollectData()
			}
		}
	}
}

func CollectData() {
	for i := 0; i < len(CodeList); i++ {
		item := CodeList[i]

		var add Forex
		add.Country = item.Country
		add.Name = item.Name
		add.MoneyCode = item.Code

		go ForexSina(add)

		time.Sleep(3 * time.Second)
	}
}

//注：100外币 兑 人民币
func ForexSina(item Forex) {

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("Start forex:", item.MoneyCode)

	//如果传入的价格是CNY，不做查找
	if "CNY" == item.MoneyCode {
		item.Rate = 1
		item.Modify = time.Now()
		log.Println(item.MoneyCode, " : ", item.Rate)

		InsertCode(item)
		return
	}

	//样例
	//http://biz.finance.sina.com.cn/forex/forex.php?startdate=2012-01-01&enddate=2015-07-07&money_code=USD&type=0

	//查询当天的价格
	startdate := time.Now().Format("2006-01-02")
	enddate := time.Now().Format("2006-01-02")

	url := "http://biz.finance.sina.com.cn/forex/forex.php?"
	url = url + "startdate=" + startdate
	url = url + "&enddate=" + enddate
	url = url + "&money_code=" + item.MoneyCode
	url = url + "&type=0"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println(err)
		return
	}

	//查找支持的货币类型
	bSpupport := false
	doc.Find("#money_code").Find("option").Each(func(i int, s *goquery.Selection) {
		v, _ := s.Attr("value")
		if item.MoneyCode == v {
			bSpupport = true
		}
	})

	//如果不支持，中断本次查找
	if false == bSpupport {
		log.Println("Not support: ", item.MoneyCode)
		return
	}

	//使用最后一个央行中间价格
	v := ""
	doc.Find(`table[class="list_table"]`).Find("tbody").Find("tr").Find("td").Each(func(i int, s *goquery.Selection) {
		v = s.Text()
	})

	rate, err := strconv.ParseFloat(v, 64)
	if nil == err {
		item.Rate = rate / 100
		item.Modify = time.Now()
		log.Println(item.MoneyCode, " : ", rate)

		InsertCode(item)
	}
}
