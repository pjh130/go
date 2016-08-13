package archiver

import (
	//	"log"
	"errors"
	"strings"
)

//压缩接口
func Compress(src string, dest []string) error {
	lowerFilename := strings.ToLower(src)
	for _, ff := range fileFormats {
		if !strings.HasSuffix(lowerFilename, ff.ext) {
			continue
		}
		return ff.create(src, dest)
	}
	return errors.New("Unsupport now!")
}

//解压接口
func Uncompress(src, dest string) error {
	lowerFilename := strings.ToLower(src)
	for _, ff := range fileFormats {
		if !strings.HasSuffix(lowerFilename, ff.ext) {
			continue
		}
		return ff.extract(src, dest)
	}
	return errors.New("Unsurport now!")
}

var fileFormats = []struct {
	ext     string
	create  CompressFunc
	extract UncompressFunc
}{
	{ext: ".zip", create: Zip, extract: Unzip},
	{ext: ".tar.gz", create: TarGz, extract: UntarGz},
	{ext: ".tgz", create: TarGz, extract: UntarGz},
	{ext: ".tar.bz2", create: TarBz2, extract: UntarBz2},
	{ext: ".rar", create: Rar, extract: Unrar},
}
