package crypto

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//计算文件的md5值
func GetFileMd5(filePath string) (string, error) {
	fd, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	h := md5.New()
	h.Write(fd)

	return hex.EncodeToString(h.Sum(nil)), err // 输出加密结果
}

//计算文件的sha1值
func GetFileSha1(filePath string) (string, error) {
	fd, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	h := sha1.New()
	h.Write(fd)

	return hex.EncodeToString(h.Sum(nil)), err // 输出加密结果
}

//计算文件的sha256值
func GetFileSha256(filePath string) (string, error) {
	fd, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	h := sha256.New()
	h.Write(fd)

	return hex.EncodeToString(h.Sum(nil)), err // 输出加密结果
}

//获取一串字符串的md5值
func String2Md5(origData string) string {
	h := md5.New()

	//	h.Write(bytes.NewBufferString(origData).Bytes())
	io.WriteString(h, origData)

	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

//获取一串字符串的sha1值
func String2Sha1(origData string) string {
	h := sha1.New()

	//	h.Write(bytes.NewBufferString(origData).Bytes())
	io.WriteString(h, origData)

	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

//获取一串字符串的sha256值
func String2Sha256(origData string) string {
	h := sha256.New()

	//	h.Write(bytes.NewBufferString(origData).Bytes())
	io.WriteString(h, origData)

	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

func Bytes2Md5(origData []byte) string {
	h := md5.New()

	h.Write(origData)

	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

func Bytes2Sha1(origData []byte) string {
	h := sha1.New()

	h.Write(origData)

	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

func Bytes2Sh256(origData []byte) string {
	h := sha256.New()

	h.Write(origData)

	return hex.EncodeToString(h.Sum(nil)) // 输出加密结果
}

func CreateHttpSid(r *http.Request, cryptoType string) string {
	var sid string

	bs := make([]byte, 24)
	if _, err := io.ReadFull(rand.Reader, bs); err != nil {
		return sid
	}

	sig := fmt.Sprintf("%s%d%s", r.RemoteAddr, time.Now().UnixNano(), bs)

	if "md5" == cryptoType {
		h := md5.New()
		h.Write([]byte(sig))
		sid = hex.EncodeToString(h.Sum(nil))
	} else if "sha1" == cryptoType {
		h := sha1.New()
		h.Write([]byte(sig))
		sid = hex.EncodeToString(h.Sum(nil))
	} else if "sha256" == cryptoType {
		h := sha256.New()
		h.Write([]byte(sig))
		sid = hex.EncodeToString(h.Sum(nil))
	} else {
		h := md5.New()
		h.Write([]byte(sig))
		sid = hex.EncodeToString(h.Sum(nil))
	}

	return sid
}
