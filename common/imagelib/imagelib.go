package imagelib

import (
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

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
	//如果都不符合上边的格式，返回
	return suf
}

//png转换成jpg
func Png2jpg(from, to string) error {
	filePng, err := os.Open(from)
	if err == nil {
		defer filePng.Close()
		m, err := png.Decode(filePng)
		if nil != err {
			return err
		} else {
			fileJpg, err := os.Create(to)
			if nil != err {
				return err
			} else {
				o := &jpeg.Options{Quality: 100}
				err = jpeg.Encode(fileJpg, m, o)
				return err
			}
		}
	} else {
		return err
	}
}

//jpg转换成png
func Jpg2png(from, to string) error {
	fileJpg, err := os.Open(from)
	if err == nil {
		defer fileJpg.Close()
		m, err := jpeg.Decode(fileJpg)
		if nil != err {
			return err
		} else {
			filePng, err := os.Create(to)
			if nil != err {
				return err
			} else {
				defer filePng.Close()
				err = png.Encode(filePng, m)
				return err
			}
		}
	} else {
		return err
	}
}
