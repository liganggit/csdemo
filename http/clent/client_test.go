package clent

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/liganggit/gotool/cipher"
	myrsa "github.com/liganggit/gotool/cipher/rsa"
)

func TestRequest(t *testing.T) {
	var (
		data      []byte
		err       error
		randomkey []byte
	)
	data = []byte("你好服务端")
	randomkey = []byte("1234567890123456")
	if data, err = cipher.RequestEncrypt(data, randomkey, myrsa.PublicKey); err != nil {
		t.Fatal(err)
	} else {
		client := &http.Client{}
		var req *http.Request
		var resp *http.Response
		var body []byte
		if req, err = http.NewRequest("POST", "http://127.0.0.1:8080/cipher", bytes.NewReader(data)); err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Content-Type", "application/octet-stream")
		if resp, err = client.Do(req); err != nil {
			t.Fatal(err)
		}

		defer resp.Body.Close()
		if body, err = ioutil.ReadAll(resp.Body); err != nil {
			t.Fatal(err)
		}

		if body, err = cipher.ResponseDecrypt(body, randomkey, myrsa.PublicKey); err != nil {
			t.Fatal(err)
		}
		t.Log(string(body))
	}
}

func BenchmarkRequest(b *testing.B) {
	var (
		data      []byte
		err       error
		randomkey []byte
	)

	data = []byte("你好服务端")
	randomkey = []byte("1234567890123456")

	for i := 0; i < b.N; i++ {
		if data, err = cipher.RequestEncrypt(data, randomkey, myrsa.PublicKey); err != nil {
			b.Fatal(err)
		} else {
			client := &http.Client{}
			var req *http.Request
			var resp *http.Response
			var body []byte
			if req, err = http.NewRequest("POST", "http://127.0.0.1:8080/cipher", bytes.NewReader(data)); err != nil {
				b.Fatal(err)
			}
			req.Header.Add("Content-Type", "application/octet-stream")
			if resp, err = client.Do(req); err != nil {
				b.Fatal(err)
			}

			defer resp.Body.Close()
			if body, err = ioutil.ReadAll(resp.Body); err != nil {
				b.Fatal(err)
			}

			if body, err = cipher.ResponseDecrypt(body, randomkey, myrsa.PublicKey); err != nil {
				b.Fatal(err)
			}
		}
	}

}

func BenchmarkNoCipherRequest(b *testing.B) {
	var (
		data []byte
		err  error
	)

	data = []byte("你好服务端")
	for i := 0; i < b.N; i++ {
		client := &http.Client{}
		var req *http.Request
		var resp *http.Response
		// var body []byte
		if req, err = http.NewRequest("POST", "http://127.0.0.1:8080/cipher", bytes.NewReader(data)); err != nil {
			b.Fatal(err)
		}
		req.Header.Add("Content-Type", "application/octet-stream")
		if resp, err = client.Do(req); err != nil {
			b.Fatal(err)
		}

		defer resp.Body.Close()
		if _, err = ioutil.ReadAll(resp.Body); err != nil {
			b.Fatal(err)
		}
	}
}
