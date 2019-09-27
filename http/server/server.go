package main

import (
	"io/ioutil"
	"net/http"

	"github.com/liganggit/gotool/cipher"
	myrsa "github.com/liganggit/gotool/cipher/rsa"
)

func Cipher(w http.ResponseWriter, r *http.Request) {
	var (
		body      []byte
		randomkey []byte
		err       error
	)
	if body, err = ioutil.ReadAll(r.Body); err != nil {
		w.WriteHeader(501)
	} else {
		response := []byte("客户端你好")
		if body, randomkey, err = cipher.RequestDecrypt(body, myrsa.PrivateKey); err != nil {
			w.WriteHeader(501)
		} else {
			if response, err = cipher.ResponseEncrypt(response, randomkey, myrsa.PrivateKey); err != nil {
				w.WriteHeader(501)
			} else {
				w.Write(response)
			}
		}
	}
}

func NoCipher(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)
	if _, err = ioutil.ReadAll(r.Body); err != nil {
		w.WriteHeader(501)
	} else {
		response := []byte("客户端你好")
		w.Write(response)
	}
}

func main() {
	http.HandleFunc("/cipher", Cipher)
	http.HandleFunc("/nocipher", NoCipher)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
