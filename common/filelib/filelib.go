package filelib

import (
	"errors"
	"path/filepath"
	"runtime"
	//"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
)

/**********************************************************************
 * 功能描述： 判断文件或者是文件夹是否存在
 * 输入参数： path-文件或者文件夹路径
 * 输出参数： 无
 * 返 回 值： bool-是/否  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func IsPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**********************************************************************
 * 功能描述： 判断文件夹是否存在
 * 输入参数： dir-文件夹路径
 * 输出参数： 无
 * 返 回 值： bool-是/否  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func IsDir(dir string) (bool, error) {
	finfo, err := os.Stat(dir)

	if err != nil {
		return false, err
	}

	if finfo.IsDir() {
		return true, nil
	}

	return false, errors.New(dir + " not a dir!")
}

/**********************************************************************
 * 功能描述： 判断是否是隐藏文件(windows平台)
 * 输入参数： path-指定路径
 * 输出参数： 无
 * 返 回 值： bool-是/否  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func CheckIsHiddenWin(path string) (bool, error) {
	file, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	//"通过反射来获取Win32FileAttributeData的FileAttributes
	fa := reflect.ValueOf(file.Sys()).Elem().FieldByName("FileAttributes").Uint()
	bytefa := []byte(strconv.FormatUint(fa, 2))
	if bytefa[len(bytefa)-2] == '1' {
		//fmt.Println("隐藏目录:", file.Name())
		return true, nil
	}

	return false, errors.New("Not hidden")
}

/**********************************************************************
 * 功能描述： 获取程序当前路径
 * 输入参数： 无
 * 输出参数： 无
 * 返 回 值： string-路径  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func GetAppCurrPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if nil != err {
		return "", err
	}

	path, err := filepath.Abs(file)
	if nil != err {
		return "", err
	}

	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]

	return ret, nil
}

/**********************************************************************
 * 功能描述： 获取一级子目录
 * 输入参数： dir-指定目录
 * 输出参数： 无
 * 返 回 值： []string-目录列表  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func GetSubDirs(dir string) ([]string, error) {
	var names []string
	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			names = append(names, sanitizedName(dir.Name()))
		}

	}

	return names, err
}

/**********************************************************************
 * 功能描述： 获取所有子目录(忽略后缀匹配的大小写)
 * 输入参数： dir-指定目录
 * 输出参数： 无
 * 返 回 值： []string-目录列表  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func GetSubDirsAll(dir string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 目录
			//过滤当前目录
			if dir != filename {
				dirs = append(dirs, sanitizedName(filename))
			}
		}
		return nil
	})

	return dirs, err
}

/**********************************************************************
 * 功能描述： 遍历目录下的所有子文件(不包括文件夹)
 * 输入参数： dir-指定目录
 * 输出参数： 无
 * 返 回 值： []string-目录列表  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func GetSubFiles(dir, suffix string) ([]string, error) {
	var files []string
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	err := filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error {
		//忽略错误
		if err != nil {
			return err
		}

		// 忽略目录
		if fi.IsDir() {
			return nil
		}

		if len(suffix) > 0 {
			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
				files = append(files, sanitizedName(filename))
			}
		} else {
			files = append(files, sanitizedName(filename))
		}

		return nil
	})
	return files, err
}

/**********************************************************************
 * 功能描述： 遍历目录下的所有子文件和文件夹
 * 输入参数： dir-指定目录
 * 输出参数： 无
 * 返 回 值： []string-目录列表  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func GetSubFilesAll(dir string) ([]string, error) {
	var files []string                                                                 //忽略后缀匹配的大小写
	err := filepath.Walk(dir, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//忽略错误
		if err != nil {
			return err
		}

		//过滤当前目录
		if dir != filename {
			files = append(files, sanitizedName(filename))
		}

		return nil
	})

	return files, err
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
