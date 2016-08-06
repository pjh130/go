package zip

import (
	"archive/zip"
	"errors"
	MyFile "github.com/pjh130/go/common/file"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

/**********************************************************************
 * 功能描述： 获取ZIP压缩包里边的文件名
 * 输入参数： fileName-压缩包
 * 输出参数： 无
 * 返 回 值： []string-目录列表  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151228           V1.0            panpan            创建
 ************************************************************************/
func ItemsZip(fileName string) ([]string, error) {
	var lst []string
	// Open a zip archive for reading.
	r, err := zip.OpenReader(fileName)
	if err != nil {
		return lst, err
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.Reader.File {
		lst = append(lst, f.Name)
	}

	return lst, err
}

/**********************************************************************
 * 功能描述： 打包ZIP文件
 * 输入参数： fileName-压缩包 srcName-压缩的东西
 * 输出参数： 无
 * 返 回 值： error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151228           V1.0            panpan            创建
 ************************************************************************/
func PackZip(fileName string, srcName string) error {
	v := time.Now().UnixNano()
	//判断要压缩的文件是否存在
	if false == MyFile.IsExist(srcName) {
		return 	errors.New(srcName + "  not exist!")
	}
	
	//绝对路径
	abs, err := filepath.Abs(srcName)
	if nil != err {
		return err
	}

	file, err := os.Create(fileName)
	if nil != err {
		return err
	}
	defer file.Close()

	var files []string
	//如果是目录就遍历全部列表
	if true == MyFile.IsDir(abs) {
		files , _ = MyFile.GetSubFilesAll(srcName, true)
	} else if MyFile.IsFile(abs){
		//如果是单一的文件，直接添加
		files = append(files, srcName)
	}
	log.Println(files)
	//没有找到要压缩的东西
	if len(files) <= 0 {
		return 	errors.New(srcName + "  is empty!")
	}

	err = CreateZip(fileName, files, abs)
	
	log.Println("Pack zip", time.Now().UnixNano()-v, "nano")
	return err
}

/**********************************************************************
 * 功能描述： 解压ZIP文件
 * 输入参数： fileName-压缩包 zipDir-解压的目录
 * 输出参数： 无
 * 返 回 值： error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151228           V1.0            panpan            创建
 ************************************************************************/
func UnpackZip(fileName string, zipDir string) error {
	v := time.Now().UnixNano()
	reader, err := zip.OpenReader((fileName))
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
		name = zipDir + string(filepath.Separator) + name

		mode := zipFile.Mode()
		//如果是目录直接创建
		if mode.IsDir() {
			_, err := os.Stat(name)
			if nil != err {
				err = os.MkdirAll(name, 0755)
				if nil != err {
					return err
				}
			}
		} else {
			err = UnpackZippedFile(name, zipFile)
			if nil != err {
				return err
			}
		}
	}

	log.Println("Unpack zip", time.Now().UnixNano()-v, "nano")
	return nil
}

func UnpackZippedFile(fileName string, zipFile *zip.File) error {
	dir := filepath.Dir(fileName)

	err := os.MkdirAll(dir, 0755)
	if nil != err {
		return err
	}
	writer, err := os.Create(fileName)
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

func sanitizedName(filename string) string {
	if len(filename) > 1 && filename[1] == ':' &&
		runtime.GOOS == "windows" {
		filename = filename[2:]
	}

	filename = filepath.ToSlash(filename)
	filename = strings.TrimLeft(filename, "/.")
	return strings.Replace(filename, "../", "", -1)
}

func CreateZip(filename string, files []string, absPath string) error {
	file, err := os.Create(filename)
	if nil != err {
		return err
	}
	defer file.Close()

	zipper := zip.NewWriter(file)
	defer zipper.Close()

	for _, name := range files {
		err = writeFileToZip(zipper, name, absPath)
		if nil != err {
			return err
		}
	}

	return err
}

func writeFileToZip(zipper *zip.Writer, filename string, absPath string) error {
	file, err := os.Open(filename)
	if nil != err {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if nil != err {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if nil != err {
		return err
	}
	
	if filename != absPath {
		header.Name = sanitizedName(strings.Replace(filename, absPath, "", -1))
	}

	//文件夹要加处理，要不解压时候会报错
	if info.IsDir() {
		header.Name += "/"
	}

	writer, err := zipper.CreateHeader(header)
	if nil != err {
		return err
	}

	if info.IsDir() {
		return nil
	} else {
		_, err = io.Copy(writer, file)
	}

	return err
}
