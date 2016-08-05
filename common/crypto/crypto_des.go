package crypto

import (
	"crypto/cipher"
	"crypto/des"
)

/*
DES（Data Encryption Standard）是对称加密算法，也就是加密和解密用相同的密钥。其入口参数有三个：
key、data、mode。key为加密解密使用的密钥，data为加密解密的数据，mode为其工作模式。当模式为加密
模式时，明文按照64位进行分组，形成明文组，key用于对数据加密，当模式为解密模式时，key用于对数据解密。
实际运用中，密钥只用到了64位中的56位，这样才具有高的安全性。DES 的常见变体是三重 DES，使用 168 位
的密钥对资料进行三次加密的一种机制；它通常（但非始终）提供极其强大的安全性。如果三个 56 位的子元素都
相同，则三重 DES 向后兼容 DES。
DES加密，涉及到加密模式和填充方式，所以，和其他语言加解密时，应该约定好加密模式和填充方式。（模式定
义了Cipher如何应用加密算法。改变模式可以容许一个块加密程序变为流加密程序。）
关于分组加密：分组密码每次加密一个数据分组，这个分组的位数可以是随意的，一般选择64或者128位。另一方
面，流加密程序每次可以加密或解密一个字节的数据，这就使它比流加密的应用程序更为有用。
在用DES加密解密时，经常会涉及到一个概念：块（block，也叫分组），模式（比如cbc），初始向量（iv），
填充方式（padding，包括none，用’\0′填充，pkcs5padding或pkcs7padding）。多语言加密解密交互时
，需要确定好这些。
*/

func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

// 3DES加密
func TripleDesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil

}

// 3DES解密
func TripleDesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}
