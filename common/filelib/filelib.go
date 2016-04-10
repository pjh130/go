package filelib

import (
	"errors"
	"path"
	"path/filepath"
	//"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
)

//获取程序绝对路径
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

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
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

/**********************************************************************
 * 功能描述： 判断文件夹是否存在
 * 输入参数： dir-文件夹路径
 * 输出参数： 无
 * 返 回 值： bool-是/否
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func IsDir(dir string) bool {
	f, e := os.Stat(dir)
	if e != nil {
		return false
	}
	return f.IsDir()
}

/**********************************************************************
 * 功能描述： 判断文是否是文件
 * 输入参数： filePath-文件路径
 * 输出参数： 无
 * 返 回 值： bool-是/否
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
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
 * 输入参数： dir-指定目录 absPath-true绝对路径/false相对路径
 * 输出参数： 无
 * 返 回 值： []string-目录列表  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func GetSubDirs(dir string, absPath bool) ([]string, error) {
	var names []string

	//绝对路径
	abs, err := filepath.Abs(dir)
	if nil != err {
		return names, err
	}

	dirs, err := ioutil.ReadDir(abs)
	if err != nil {
		return nil, err
	}

	for _, v := range dirs {
		if v.IsDir() {
			if absPath {
				names = append(names, abs+v.Name())
			} else {
				names = append(names, v.Name())
			}
		}
	}

	return names, err
}

/**********************************************************************
 * 功能描述： 获取所有子目录
 * 输入参数： dir-指定目录 absPath-true绝对路径/false相对路径
 * 输出参数： 无
 * 返 回 值： []string-目录列表  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func GetSubDirsAll(dir string, absPath bool) ([]string, error) {
	var dirs []string

	//绝对路径
	abs, err := filepath.Abs(dir)
	if nil != err {
		return dirs, err
	}

	err = filepath.Walk(abs, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 目录
			//过滤当前目录
			if abs != filename {
				if absPath {
					dirs = append(dirs, filename)
				} else {
					dirs = append(dirs, strings.Replace(filename, abs, "", -1))
				}
			}
		}
		return nil
	})

	return dirs, err
}

/**********************************************************************
 * 功能描述： 遍历目录下的所有子文件(不包括文件夹)
 * 输入参数： dir-指定目录 absPath-true绝对路径/false相对路径
 * 输出参数： 无
 * 返 回 值： []string-目录列表  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func GetSubFiles(dir, suffix string, absPath bool) ([]string, error) {
	var files []string

	//绝对路径
	abs, err := filepath.Abs(dir)
	if nil != err {
		return files, err
	}

	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	err = filepath.Walk(abs, func(filename string, fi os.FileInfo, err error) error {
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
				if absPath {
					files = append(files, filename)
				} else {
					files = append(files, strings.Replace(filename, abs, "", -1))
				}
			}
		} else {
			if absPath {
				files = append(files, filename)
			} else {
				files = append(files, strings.Replace(filename, abs, "", -1))
			}
		}

		return nil
	})
	return files, err
}

/**********************************************************************
 * 功能描述： 遍历目录下的所有子文件和文件夹
 * 输入参数： dir-指定目录 absPath-true绝对路径/false相对路径
 * 输出参数： 无
 * 返 回 值： []string-目录列表  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func GetSubFilesAll(dir string, absPath bool) ([]string, error) {
	var files []string

	//绝对路径
	abs, err := filepath.Abs(dir)
	if nil != err {
		return files, err
	}

	err = filepath.Walk(abs, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//忽略错误
		if err != nil {
			return err
		}

		//过滤当前目录
		if abs != filename {
			if true == absPath {
				files = append(files, filename)
			} else {
				files = append(files, strings.Replace(filename, abs, "", -1))
			}
		}

		return nil
	})

	return files, err
}

/**********************************************************************
 * 功能描述： 复制文件夹
 * 输入参数： srcPath-原始目录 destPath-目标目录
 * 输出参数： 无
 * 返 回 值： error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func CopyDir(srcPath, destPath string) error {
	// Check if target directory exists.
	if IsExist(destPath) {
		return errors.New("file or directory alreay exists: " + destPath)
	}

	//创建目录
	err := os.MkdirAll(destPath, os.ModePerm)
	if err != nil {
		return err
	}

	//先获取原目录下所有的文件
	list, err := GetSubFilesAll(srcPath, false)
	if nil != err {
		return err
	}

	//获取绝对路径
	absPath, err := filepath.Abs(srcPath)
	for _, v := range list {
		//重新拼接新的文件路径
		curPath := path.Join(destPath, v)
		vv := path.Join(absPath, v)

		if IsDir(vv) {
			err = os.MkdirAll(curPath, os.ModePerm)
		} else {
			err = CopyFile(vv, curPath)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// Copy copies file from source to target path.
func CopyFile(src, dest string) error {
	// Gather file information to set back later.
	si, err := os.Lstat(src)
	if err != nil {
		return err
	}

	// Handle symbolic link.
	if si.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(src)
		if err != nil {
			return err
		}
		// NOTE: os.Chmod and os.Chtimes don't recoganize symbolic link,
		// which will lead "no such file or directory" error.
		return os.Symlink(target, dest)
	}

	sr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sr.Close()

	dw, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dw.Close()

	if _, err = io.Copy(dw, sr); err != nil {
		return err
	}

	// Set back file information.
	if err = os.Chtimes(dest, si.ModTime(), si.ModTime()); err != nil {
		return err
	}
	return os.Chmod(dest, si.Mode())
}
