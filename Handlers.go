
/*
当时在北京一家公司实习，花了一天的时间查博客看教程，最终实现了产品所希望的功能，这段代码是当时提交工作代码之前在本地写的demo，测试可以跑通，非常开心，算是code最有成就感的一次。那份实习也是我第一份写代码的工作，当时的golang也是重头学起的语言，边学边干，所以印象最深。
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>drill</title>
<body><h3>drill</h3>`
	form = `<form action="/todos/sub" method="POST">
<label for="numbers">query</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="submit">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func PostTodoIndex(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // Must be called before writing response
	fmt.Fprint(w, pageTop, form)
	if err != nil {
		fmt.Fprintf(w, anError, err)
	}
	slice, found := r.Form["numbers"]
	glog.Info("s", slice)
	glog.Info("\n")

	if !found || len(slice) < 0 {
		fmt.Println("please input some sqls")
	}
	headers := map[string]string{"Content-Type": "application/json; charset=UTF-8"}
	body := map[string]interface{}{
		"queryType": "SQL",
		"query":     " ",
		// "query": " SELECT url FROM dfs.`/tmp/a2/ass/userlog/v1/2016-06/23/ass-userlog-2016-06-23-120837.588-b96fbd0be7fb.json` limit 10",
	}
	body["query"] = slice[0]
	// var response interface{}
	response, err := PostJSONBody("http://localhost:8047/query.json", headers, body)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	jsonResponse, _ := json.Marshal(response)
	fmt.Fprint(w, string(jsonResponse))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // Must be called before writing response
	fmt.Fprint(w, pageTop, form)
	if err != nil {
		fmt.Fprintf(w, anError, err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, pageBottom)
}

// func PostJSONBody(urlString string, headers map[string]string, body map[string]interface{}, response interface{}) error {
func PostJSONBody(urlString string, headers map[string]string, body map[string]interface{}) (interface{}, error) {

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	// glog.Info("body", string(bodyBytes))
	req, err := http.NewRequest("POST", urlString, bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Printf("err: %v", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v", err)
		return nil, err
	}
	// glog.Info("res", string(respBody))
	var response interface{}
	if err = json.Unmarshal(respBody, &response); err != nil {
		fmt.Printf("err: %v", err)
		return nil, err
	}

	fmt.Printf("done")
	return response, nil
}
