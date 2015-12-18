package ziplib

import (
	"archive/zip"
	"errors"
	"github.com/pjh130/go/common/filelib"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
 * 输入参数： fileName-压缩包 zipDir-压缩的目录
 * 输出参数： 无
 * 返 回 值： error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151228           V1.0            panpan            创建
 ************************************************************************/
func PackZip(fileName string, dirName string) error {
	file, err := os.Create(fileName)
	if nil != err {
		return err
	}
	defer file.Close()

	if is, _ := filelib.IsDir(dirName); !is {
		return errors.New(dirName + " is not a dir")
	}

	files, _ := filelib.GetSubFilesAll(dirName)

	err = CreateZip(fileName, files)
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
		name = sanitizedName(name)

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
		}
	}
	return nil
}

func UnpackZippedFile(fileName string, zipFile *zip.File) error {
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

func CreateZip(filename string, files []string) error {
	file, err := os.Create(filename)
	if nil != err {
		return err
	}
	defer file.Close()

	zipper := zip.NewWriter(file)
	defer zipper.Close()

	for _, name := range files {
		err = writeFileToZip(zipper, name)
		if nil != err {
			return err
		}
	}

	return err
}

func writeFileToZip(zipper *zip.Writer, filename string) error {
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

	header.Name = sanitizedName(filename)
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
