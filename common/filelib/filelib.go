package filelib

import (
	"errors"
	"path/filepath"
	//"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
)

func IsDir(dirpath string) (bool, error) {
	finfo, err := os.Stat(dirpath)

	if err != nil {
		return false, err
	}

	if finfo.IsDir() {
		return true, nil
	}

	return false, errors.New(dirpath + " not a dir!")
}

func GetDirNames(dirpath string) ([]string, error) {
	var names []string
	dirs, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			names = append(names, dir.Name())
		}

	}

	return names, err
}

func CheckIsHiddenWin(file os.FileInfo) bool {
	//"通过反射来获取Win32FileAttributeData的FileAttributes
	fa := reflect.ValueOf(file.Sys()).Elem().FieldByName("FileAttributes").Uint()
	bytefa := []byte(strconv.FormatUint(fa, 2))
	if bytefa[len(bytefa)-2] == '1' {
		//fmt.Println("隐藏目录:", file.Name())
		return true
	}
	return false
}

//获取程序绝对路径
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

//获取指定目录及所有子目录下的所有文件()，可以匹配后缀过滤。
func WalkDirFiles(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if len(suffix) > 0 {
			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
				files = append(files, filename)
			}
		} else {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

//获取指定目录及所有子目录下的所有文件和文件夹
func WalkDirAll(dirPth string) (files []string, err error) {
	files = make([]string, 0, 30)                                                        //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if dirPth != filename {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

//获取指定目录及所有子目录下的所有文件和文件夹
func WalkSubDirs(dirPth string) (dirs []string, err error) {
	dirs = make([]string, 0, 30)                                                         //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 目录
			if dirPth != filename {
				dirs = append(dirs, filename)
			}
		}
		return nil
	})
	return dirs, err
}
