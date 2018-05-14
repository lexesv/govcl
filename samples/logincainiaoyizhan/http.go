package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
)

// httpGet http请求
func httpGet(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// httpGetJSON 获取json
func httpGetJSON(url string, result interface{}) error {
	data, err := httpGet(url)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, result); err != nil {
		return err
	}
	return nil
}
