package main

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
)

type TCNSession struct {
	client       *http.Client
	jar          *cookiejar.Jar
	ResponseCode int
	sync.Mutex
}

func NewCNSession() *TCNSession {
	c := new(TCNSession)
	c.client = new(http.Client)
	jar, _ := cookiejar.New(nil)
	c.jar = jar
	c.client.Jar = jar
	return c
}

func (c *TCNSession) ClearCookie() {
	c.jar.SetCookies(nil, nil)
}

func (c *TCNSession) Get(urlstr string) ([]byte, error) {
	c.Lock()
	defer c.Unlock()
	req, err := http.NewRequest(http.MethodGet, urlstr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := c.client.Do(req)
	c.ResponseCode = resp.StatusCode
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

func (c *TCNSession) Post(urlstr string, param url.Values) ([]byte, error) {
	c.Lock()
	defer c.Unlock()
	req, err := http.NewRequest(http.MethodPost, urlstr, strings.NewReader(param.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req)
	c.ResponseCode = resp.StatusCode
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
