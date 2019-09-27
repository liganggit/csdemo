package clent

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/liganggit/gotool/cipher"
	myrsa "github.com/liganggit/gotool/cipher/rsa"
)

//Request 客户端请求
func Request() []byte {
	var (
		data      []byte
		err       error
		randomkey []byte
	)
	data = []byte("你好服务端")
	randomkey = []byte("1234567890123456")
	if data, err = cipher.RequestEncrypt(data, randomkey, myrsa.PublicKey); err != nil {
		panic(err)
	} else {
		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://127.0.0.1:8080/cipher", bytes.NewReader(data))
		req.Header.Add("Content-Type", "application/octet-stream")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		if body, err = cipher.ResponseDecrypt(body, randomkey, myrsa.PublicKey); err != nil {
			panic(err)
		} else {
			return body
		}
	}
}
