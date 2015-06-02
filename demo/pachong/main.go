package main
 
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
)
 
var (
    ptnIndexItem    = regexp.MustCompile(`<a target="_blank" href="(.+\.html)" title=".+" >(.+)</a>`)
    ptnContentRough = regexp.MustCompile(`(?s).*<div class="artcontent">(.*)<div id="zhanwei">.*`)
    ptnBrTag        = regexp.MustCompile(`<br>`)
    ptnHTMLTag      = regexp.MustCompile(`(?s)</?.*?>`)
    ptnSpace        = regexp.MustCompile(`(^\s+)|( )`)
)
 
func Get(url string) (content string, statusCode int) {
    resp, err1 := http.Get(url)
    if err1 != nil {
        statusCode = -100
        return
    }
    defer resp.Body.Close()
    data, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        statusCode = -200
        return
    }
    statusCode = resp.StatusCode
    content = string(data)
    return
}
 
type IndexItem struct {
    url   string
    title string
}
 
func findIndex(content string) (index []IndexItem, err error) {
    matches := ptnIndexItem.FindAllStringSubmatch(content, 10000)
    index = make([]IndexItem, len(matches))
    for i, item := range matches {
        index[i] = IndexItem{"http://www.yifan100.com" + item[1], item[2]}
    }
    return
}
 
func readContent(url string) (content string) {
    raw, statusCode := Get(url)
    if statusCode != 200 {
        fmt.Print("Fail to get the raw data from", url, "\n")
        return
    }
 
    match := ptnContentRough.FindStringSubmatch(raw)
    if match != nil {
        content = match[1]
    } else {
        return
    }
 
    content = ptnBrTag.ReplaceAllString(content, "\r\n")
    content = ptnHTMLTag.ReplaceAllString(content, "")
    content = ptnSpace.ReplaceAllString(content, "")
    return
}
 
func main() {
    fmt.Println(`Get index ...`)
    s, statusCode := Get("http://www.yifan100.com/dir/15136/")
    if statusCode != 200 {
        return
    }
    index, _ := findIndex(s)
 
    fmt.Println(`Get contents and write to file ...`)
    for _, item := range index {
        fmt.Printf("Get content %s from %s and write to file.\n", item.title, item.url)
        fileName := fmt.Sprintf("%s.txt", item.title)
		fmt.Println(fileName)
        content := readContent(item.url)
        ioutil.WriteFile(fileName, []byte(content), 0644)
        fmt.Printf("Finish writing to %s.\n", fileName)
    }
}