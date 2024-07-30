package lib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"noauth/poc"
	"strings"
	"sync"
)

func PostStart(url, noauth, auth string, thread int, debug int) {

	fmt.Println(Blue("[+] POST(Form-data and Json) poc start "))

	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}

	resp, err := http.Post(url+auth, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte{}))

	respjson, errjson := http.Post(url+auth, "application/json", bytes.NewBuffer([]byte("{}")))

	if errjson != nil {
		//fmt.Println("Error:", errjson)
		return
	}

	if err != nil {
		//fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	defer respjson.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyjson, errjson := ioutil.ReadAll(respjson.Body)
	if errjson != nil {
		//fmt.Println("Error reading body:", errjson)
		return
	}

	if err != nil {
		//fmt.Println("Error reading body:", err)
		return
	}

	if strings.Contains(string(body), url+auth) {
		body = []byte(strings.Replace(string(body), url+auth, "", 1))
	}
	if strings.Contains(string(bodyjson), url+auth) {
		body = []byte(strings.Replace(string(bodyjson), url+auth, "", 1))
	}

	len1 := len(body)
	fmt.Printf(Green("[+] Length of the Original Authentication API %s: len=%d\n"), url+auth, len1)

	lenjson := len(bodyjson)

	list := poc.Summary(noauth, auth)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, thread)
	mu := &sync.Mutex{}

	for _, value := range list {
		wg.Add(1)
		go func(value string) {
			defer wg.Done()

			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			resp, err := http.Post(url+value, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte{}))
			respjson, errjson := http.Post(url+value, "application/json", bytes.NewBuffer([]byte{}))

			if err != nil {
				//fmt.Println("Error:", err)
				return
			}
			if errjson != nil {
				//fmt.Println("Error:", errjson)
				return
			}
			defer respjson.Body.Close()
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			bodyjson, errjson := ioutil.ReadAll(respjson.Body)

			if err != nil {
				//fmt.Println("Error reading body:", err)
				return
			}

			if errjson != nil {
				//fmt.Println("Error reading body:", errjson)
				return
			}

			if strings.Contains(string(body), url+value) {
				body = []byte(strings.Replace(string(body), url+value, "", 1))
			}

			len2 := len(body)

			if strings.Contains(string(bodyjson), url+value) {
				body = []byte(strings.Replace(string(bodyjson), url+value, "", 1))
			}

			len2json := len(bodyjson)

			mu.Lock()
			if len2 != len1 && resp.StatusCode != 404 && debug != 1 {
				fmt.Printf(Green("[+] Post: Length mismatch for %s len=%d code=%d\n"), url+value, len2, resp.StatusCode)
			}

			if len2json != lenjson && len2json != len2 && respjson.StatusCode != 404 && debug != 1 {
				fmt.Printf(Green("[+] Post-Json: Length mismatch for %s len=%d code=%d\n"), url+value, len2json, respjson.StatusCode)
			}

			if debug == 1 {
				fmt.Printf(Green("[+] Post: Length mismatch for %s len=%d code=%d\n"), url+value, len2, resp.StatusCode)
			}
			if debug == 1 && len2json != len2 {
				fmt.Printf(Green("[+] Post-Json: Length mismatch for %s len=%d code=%d\n"), url+value, len2json, respjson.StatusCode)
			}

			mu.Unlock()
		}(value)
	}

	wg.Wait()
}
