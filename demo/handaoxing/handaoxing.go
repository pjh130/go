package main

import (
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/iconv.v1"
)

type Chapter struct {
	url  string
	name string
}

//转换字符编码
func ChangeCodeType(src string, srcType, destType string) (string, error) {
	dest := ""

	//	cd, err := iconv.Open("gbk", "utf-8") // convert utf-8 to gbk
	cd, err := iconv.Open(destType, srcType)
	if err != nil {
		return dest, err
	}
	defer cd.Close()

	dest = cd.ConvString(src)

	return dest, err
}

func GetStringItems(src, begin, end string) []string {
	var items []string
	temp := src
	//	log.Println(temp)
	for {
		b := strings.Index(temp, begin)
		if b == -1 {
			break
		}

		e := strings.Index(temp, end)
		if e == -1 {
			break
		}

		add := strings.TrimSpace(temp[b+len(begin) : e])
		items = append(items, add)
		temp = temp[e+len(end):]
	}

	return items
}

//获取
func GetPartAll(str string) (map[string][]Chapter, error) {
	lst := make(map[string][]Chapter)

	doc, err := goquery.NewDocument(str)
	if err != nil {
		log.Println(err)
		return lst, err
	}

	//先找出列表
	s := doc.Find("#list dl")
	if nil == s {
		log.Println("Find nil")
		return lst, errors.New("Can't find #list")
	}

	file, err := os.Create("test.txt")
	if err != nil {
		log.Println(err)
		return lst, err
	}
	defer file.Close()

	html, _ := s.Html()

	temp := html
	begin := "<dt>"
	end := "</dt>"

	for {
		b := strings.Index(temp, begin)
		if b == -1 {
			break
		}

		e := strings.Index(temp, end)
		if e == -1 {
			break
		}

		//获取dt的值
		dt, _ := ChangeCodeType(strings.TrimSpace(temp[b+len(begin):e]), "gbk", "utf-8")

		add := make([]Chapter, 0)
		tempNew := ""

		//判断是否还存在下一个dt
		if n := strings.Index(temp[e+len(end):], begin); n != -1 {
			tempNew = temp[e+len(end) : e+len(end)+n]
		} else {
			tempNew = temp[e+len(end):]
		}

		//找出dd的值
		items := GetStringItems(tempNew, "<dd>", "</dd>")
		if len(items) > 0 {
			for _, dd := range items {
				d, _ := ChangeCodeType(dd, "gbk", "utf-8")

				if len(d) > 0 {
					var cc Chapter
					v1 := "href=\""
					v2 := "\">"
					index1 := strings.Index(d, v1)
					index2 := strings.Index(d, v2)

					v3 := "</a>"
					index3 := strings.Index(d, v3)
					if index1 != -1 && index2 != -1 && index3 != -1 {
						href := d[index1+len(v1) : index2]
						name := d[index2+len(v2) : index3]

						//处理特殊字符
						name = strings.Replace(name, "，", "_", -1)
						name = strings.Replace(name, ",", "_", -1)
						name = strings.Replace(name, "。", "_", -1)
						name = strings.Replace(name, ".", "_", -1)
						name = strings.Replace(name, "！", "", -1)
						name = strings.Replace(name, "！", "", -1)

						cc.url = str + strings.TrimSpace(href)
						cc.name = strings.TrimSpace(name)

						add = append(add, cc)
					} else {
						//							log.Println("******************index123 < 0")
					}
				} else {
					//						log.Println("!!!!!!!!!!!!!!!!!!!!!!!len dd < 0")
				}
			}
		} else {
			//				log.Println("~~~~~~~~~~~~~~~~~~~~~items < 0")
		}

		if len(add) <= 0 {
			//				log.Println("@@@@@@@@@@@@@@@@@@@@@@len(add) <= 0")
		}

		if len(dt) > 0 && len(add) > 0 {
			//重名处理
			k := 0
			dtNew := dt
			for {
				k++
				_, ok := lst[dtNew]
				if false == ok {
					break
				} else {
					//为了避免重复处理dt
					dtNew = dt + strconv.Itoa(k)
				}
			}
			lst[dtNew] = add
		}

		//循环偏移
		temp = temp[e+len(end):]
	}

	return lst, err
}

//获取
func GetPart(str string) (map[string][]Chapter, error) {
	lst := make(map[string][]Chapter)

	doc, err := goquery.NewDocument(str)
	if err != nil {
		log.Println(err)
		return lst, err
	}

	//先找出列表
	s := doc.Find("#list")
	if nil == s {
		log.Println("Find nil")
		return lst, errors.New("Can't find #list")
	}

	//	s.Find("dt").Each(func(i int, ss *goquery.Selection) {
	s.Find("dd a").Each(func(i int, ss *goquery.Selection) {
		v := ss.Text()
		href, _ := ss.Attr("href")
		if len(v) > 0 {
			vv, err := ChangeCodeType(v, "gbk", "utf-8")
			if err == nil {
				//				lst = append(lst, vv)
				var cc Chapter
				cc.url = str + strings.TrimSpace(href)
				cc.name = strings.TrimSpace(vv)
				key := strconv.Itoa(i)
				value, ok := lst[key]
				if ok {
					value = append(value, cc)
					lst[key] = value
				} else {
					value := make([]Chapter, 0)
					value = append(value, cc)
					lst[key] = value
				}

				//				log.Println(cc.name, cc.url)
			} else {
				log.Println(err)
			}
		} else {
			log.Println("v is null")
		}
	})

	//	log.Println(lst)
	return lst, nil
}

func GetContent(str string) (string, error) {
	doc, err := goquery.NewDocument(str)
	if err != nil {
		log.Println(err)
		return "", err
	}

	//先找出列表
	s := doc.Find("#content")
	if nil == s {
		log.Println("Find nil")
		return "", errors.New("Can't find #content")
	}

	text, _ := s.Html()
	//	log.Println("length content: ", len(text))

	//回车
	text = strings.Replace(text, "<br/>", "\r\n", -1)

	//特殊字符替换
	b, err := hex.DecodeString("c2a0")
	if nil == err {
		text = strings.Replace(text, string(b), " ", -1)
	}

	//
	//	vv, err := ChangeCodeType(text, "gbk", "utf-8")
	//	if err != nil {
	//		return "", err
	//	}

	return text, nil
}

//获取章节
func GetChapters(strUrl string) (map[string][]Chapter, error) {
	lst := make(map[string][]Chapter)

	resp, err := http.Get(strUrl)
	if nil != err {
		log.Println(err)
		return lst, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Println(err)
		return lst, err
	}
	//	log.Println(string(b))
	defer resp.Body.Close()

	begin := "<dt>"
	end := "</dt>"
	vv, err := ChangeCodeType(string(b), "gbk", "utf-8")
	if err == nil {
		//找出块
		GetStringItems(vv, begin, end)
	}

	//temp := string(b)
	//	log.Println(temp)
	//	for {
	//		b := strings.Index(temp, begin)
	//		if b == -1 {
	//			break
	//		}

	//		e := strings.Index(temp, end)
	//		if e == -1 {
	//			break
	//		}

	//		//如果还有下一个BEGIN
	//		temp = temp[e+len(end):]
	//		tempNew := temp
	//		part := strings.TrimSpace(temp[b+len(begin) : e])

	//		n := strings.Index(temp, begin)
	//		if n != -1 {
	//			tempNew =
	//		}

	//	}

	log.Println(lst)
	return lst, nil
}
