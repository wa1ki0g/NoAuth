package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"noauth/poc"
	"strings"
	"sync"
)

func GetStart(url, noauth, auth string, thread int, debug int) {

	fmt.Println(Blue("[+] GET poc start "))

	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}

	resp, err := http.Get(url + auth)
	if err != nil {
		//fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("Error reading body:", err)
		return
	}

	if strings.Contains(string(body), url+auth) {
		body = []byte(strings.Replace(string(body), url+auth, "", 1))
	}

	len1 := len(body)
	fmt.Printf(Green("[+] Length of the Original Authentication API %s: len=%d\n"), url+auth, len1)

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

			resp, err := http.Get(url + value)
			if err != nil {
				//fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				//fmt.Println("Error reading body:", err)
				return
			}

			if strings.Contains(string(body), url+value) {
				body = []byte(strings.Replace(string(body), url+value, "", 1))
			}

			len2 := len(body)

			mu.Lock()
			if len2 != len1 && resp.StatusCode != 404 && debug != 1 {
				fmt.Printf(Green("[+] GET: Length mismatch for %s len=%d code=%d\n"), url+value, len2, resp.StatusCode)
			}

			if debug == 1 {
				fmt.Printf(Green("[+] GET: Length mismatch for %s len=%d code=%d\n"), url+value, len2, resp.StatusCode)
			}

			mu.Unlock()
		}(value)
	}

	wg.Wait()
}
