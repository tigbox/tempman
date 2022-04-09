package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var list []string = []string{`{"userID":56740525873831}`, `{"userID":56920281549018}`, `{"userID":51291733832514}`, `{"userID":55669227494966}`, `{"userID":56288281680374}`, `{"userID":56902507067896}`, `{"userID":57758498932824}`, `{"userID":57100159884284}`, `{"userID":48045586718322}`, `{"userID":57194970241265}`, `{"userID":57630105682220}`, `{"userID":57100212313216}`}

func main() {
	// list = []string{`{"userID":31356051677834}`}
	url := "http://172.20.4.14:8080/reload_cluster?namespace=default"
	method := "POST"

	for index, item := range list {
		time.Sleep(time.Second)
		fmt.Printf("index=%v, [item]%v\n", index, item)
		payload := strings.NewReader(item)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
	}

}
