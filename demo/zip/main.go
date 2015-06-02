package main

import (
	"archive/zip"
	"bytes"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pjh130/common/filelib"
)

func main() {
	Test1()
	return
	//dirName := "./"
	//files1, _ := filelib.WalkDirAll(dirName)
	//fmt.Println(files1)
	//files2, _ := filelib.WalkDirFiles(dirName, "")
	//fmt.Println(files2)

	name := "./test.zip"
	createZipDir(name, "G:\\BaiduMusic")

	//var files []string
	//files = append(files, "./main.go")
	//files = append(files, "./test_dir")

	//createZip(name, files)
	unpackZip(name, "./unzip")
}

func Test1() {
	var input = []byte("data:image/png;base64,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")
	var buf bytes.Buffer
	compressor, err := zlib.NewWriterLevelDict(&buf, zlib.BestCompression, input)
	if err != nil {
		fmt.Println("压缩失败")
		return
	}
	compressor.Write(input)
	compressor.Close()
	//	fmt.Println(buf.Bytes())
	fmt.Println(len(buf.Bytes()))
	fmt.Println(len(input))

}

func createZipDir(filename string, dirName string) error {
	file, err := os.Create(filename)
	if nil != err {
		return err
	}
	defer file.Close()

	if is, _ := filelib.IsDir(dirName); !is {
		return errors.New(dirName + " is not a dir")
	}

	files, _ := filelib.WalkDirAll(dirName)
	//fmt.Println(files)

	err = createZip(filename, files)
	fmt.Println(err)
	return err
}

func createZip(filename string, files []string) error {
	file, err := os.Create(filename)
	if nil != err {
		fmt.Println("1")
		return err
	}
	defer file.Close()

	zipper := zip.NewWriter(file)
	defer zipper.Close()

	for _, name := range files {
		err = writeFileToZip(zipper, name)
		if nil != err {
			fmt.Println("2")
			return err
		}
	}

	return err
}

func writeFileToZip(zipper *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if nil != err {
		fmt.Println("7777")
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if nil != err {
		fmt.Println("111")
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if nil != err {
		fmt.Println("222")
		return err
	}

	header.Name = sanitizedName(filename)
	writer, err := zipper.CreateHeader(header)
	if nil != err {
		fmt.Println("333")
		return err
	}
	if info.IsDir() {
		return nil
	} else {
		_, err = io.Copy(writer, file)
		//fmt.Println("444")
	}

	return err
}

func sanitizedName(filename string) string {
	if len(filename) > 1 && filename[1] == ':' &&
		runtime.GOOS == "windows" {
		filename = filename[2:]

	}

	filename = filepath.ToSlash(filename)
	filename = strings.TrimLeft(filename, "/.")
	return strings.Replace(filename, "../", "", -1)
}

func unpackZip(filename string, zipDir string) error {
	reader, err := zip.OpenReader((filename))
	if nil != err {
		return err
	}
	defer reader.Close()

	if len(zipDir) > 0 {
		err = os.MkdirAll(zipDir, 0755)
		if nil != err {
			return err
		}
	}

	for _, zipFile := range reader.Reader.File {
		name := sanitizedName(zipFile.Name)
		name = zipDir + "/" + name
		mode := zipFile.Mode()
		if mode.IsDir() {
			err = os.MkdirAll(name, 0755)
			if nil != err {
				return err
			}
		} else {
			err = unpackZippedFile(name, zipFile)
		}
	}
	return nil
}

func unpackZippedFile(filename string, zipFile *zip.File) error {
	writer, err := os.Create(filename)
	if nil != err {
		return err
	}
	defer writer.Close()

	reader, err := zipFile.Open()
	if nil != err {
		return err
	}
	defer reader.Close()

	_, err = io.Copy(writer, reader)
	if nil != err {
		return err
	}

	return err
}
