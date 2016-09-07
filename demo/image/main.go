package main

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/pjh130/go/common/image"
)

func main() {
	fmt.Println("--->", GetImageSuffix("./1.jpg"))
	fmt.Println("--->", GetImageSuffix("./1.gif"))
	fmt.Println("--->", GetImageSuffix("./1.png"))

	image.Png2jpg("./1.png", "./png2jpg.jpg")
	image.Jpg2png("./1.jpg", "./jpg2png.png")
}

//返回图片后缀
func GetImageSuffix(fileName string) string {
	var suf string
	if len(fileName) <= 0 {
		return fileName
	}

	fileJpg, err := os.Open(fileName)
	if err == nil {
		defer fileJpg.Close()
		_, err = jpeg.Decode(fileJpg)
		if nil != err {
			//			fmt.Println("jpeg :", err)
		} else {
			suf = "jpg"
			return suf
		}
	} else {
		fmt.Println("err: :", err)
		return suf
	}

	fileGif, err := os.Open(fileName)
	if err == nil {
		defer fileGif.Close()
		_, err = gif.Decode(fileGif)
		if nil != err {
			//			fmt.Println("gif :", err)
		} else {
			suf = "gif"
			return suf
		}
	}

	filePng, err := os.Open(fileName)
	if err == nil {
		defer filePng.Close()
		_, err = png.Decode(filePng)
		if nil != err {
			//			fmt.Println("png :", err)
		} else {
			suf = "png"
			return suf
		}
	}
	//如果都不符合上边的格式，返回原来的文件名
	return suf
}

//如果文件没有后缀，返回图片后缀
func GetAddImageSuffix(fileName string) string {

	var newFile string
	if len(fileName) <= 0 {
		return fileName
	}

	//如果图片已经有后缀不做处理
	info, _ := os.Lstat(fileName)
	baseName := info.Name()
	index := strings.LastIndex(baseName, ".")
	if -1 != index {
		fmt.Println("GetAddImageSuffix have suffix")
		return fileName
	}

	fileJpg, err := os.Open(fileName)
	if err == nil {
		_, err = jpeg.Decode(fileJpg)
		// fmt.Println("jpeg :", err)
		fileJpg.Close()
		if nil == err {
			newFile = fileName + ".jpg"
			return newFile
		}
	} else {
		//如果文件打不开，返回原来的字串
		return fileName
	}

	fileGif, err := os.Open(fileName)
	if err == nil {
		_, err = gif.Decode(fileGif)
		// fmt.Println("gif :", err)
		fileGif.Close()
		if nil == err {
			newFile = fileName + ".gif"
			return newFile
		}
	}

	filePng, err := os.Open(fileName)
	if err == nil {
		_, err = png.Decode(filePng)
		// fmt.Println("png :", err)
		filePng.Close()
		if nil == err {
			newFile = fileName + ".png"
			return newFile
		}
	}

	fmt.Println("GetAddImageSuffix not image")
	//如果都不符合上边的格式，返回原来的文件名
	return fileName
}
