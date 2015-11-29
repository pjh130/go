package imagelib

import (
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

/**********************************************************************
 * 功能描述： 获取图片后缀，如果不是图片返回错误
 * 输入参数： fileName-指定图片名
 * 输出参数： 无
 * 返 回 值： string-后缀名  error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
func GetImageSuffix(fileName string) (string, error) {

	if len(fileName) <= 0 {
		return "", errors.New("file name is empty")
	}

	//检查是否是jpg格式
	fileJpg, err := os.Open(fileName)
	if err == nil {
		defer fileJpg.Close()
		_, err = jpeg.Decode(fileJpg)
		if nil != err {
			//不做处理，留个下边分支去判断
		} else {
			return "jpg", nil
		}
	} else {
		return "", err
	}

	//检查是否是gif格式
	fileGif, err := os.Open(fileName)
	if err == nil {
		defer fileGif.Close()
		_, err = gif.Decode(fileGif)
		if nil != err {
			//不做处理，留个下边分支去判断
		} else {
			return "gif", nil
		}
	} else {
		return "", err
	}

	//检查是否是png格式
	filePng, err := os.Open(fileName)
	if err == nil {
		defer filePng.Close()
		_, err = png.Decode(filePng)
		if nil != err {
			//不做处理，留个下边分支去判断
		} else {
			return "png", nil
		}
	} else {
		return "", err
	}

	//如果都不符合上边的格式，返回
	return "", errors.New("Unknow image")
}

/**********************************************************************
 * 功能描述： png转换成jpg
 * 输入参数： from-原文件路径 to-目标文件路径
 * 输出参数： 无
 * 返 回 值： error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
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

/**********************************************************************
 * 功能描述： jpg转换成png
 * 输入参数： from-原文件路径 to-目标文件路径
 * 输出参数： 无
 * 返 回 值： error-错误信息
 * 其它说明： 无
 * 修改日期            版本号            修改人           修改内容
 * ----------------------------------------------------------------------
 *  20151129           V1.0            panpan            创建
 ************************************************************************/
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
