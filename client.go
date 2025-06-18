package pocket

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Client struct {
	mu           sync.Mutex
	lastCallTime time.Time

	Interval time.Duration
	Token    string
	AppInfo  string
}

type AppInfo struct {
	Vendor     string `json:"vendor"`
	DeviceID   string `json:"deviceId"`
	AppVersion string `json:"appVersion"`
	AppBuild   string `json:"appBuild"`
	OSVersion  string `json:"osVersion"`
	OSType     string `json:"osType"`
	DeviceName string `json:"deviceName"`
	OS         string `json:"os"`
}

func NewRandomAppInfo() *AppInfo {
	return &AppInfo{
		Vendor:     "apple",
		DeviceID:   strings.ToUpper(uuid.NewString()),
		AppVersion: "7.1.22",
		AppBuild:   "25012101",
		OSVersion:  "17.1.1",
		OSType:     "ios",
		DeviceName: "iPhone 12 mini",
		OS:         "ios",
	}
}

func NewClient(token string, interval int, appInfo *AppInfo) *Client {
	if appInfo == nil {
		appInfo = NewRandomAppInfo()
	}
	c := Client{
		mu:           sync.Mutex{},
		lastCallTime: time.Time{},
		Interval:     time.Duration(interval) * time.Millisecond,
		Token:        token,
		AppInfo:      JsonMarshal(*appInfo),
	}
	return &c
}

func (c *Client) doPocketRequest(req *http.Request) (*http.Response, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(c.lastCallTime)
	if elapsed < c.Interval {
		waitTime := c.Interval - elapsed
		time.Sleep(waitTime)
	}
	c.lastCallTime = time.Now()

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", c.Token)
	req.Header.Set("appInfo", c.AppInfo)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	return resp, err
}

type responseFormatter[T any] struct {
	c *Client
}

// req should be able to marshaled to json
func (f responseFormatter[T]) doRestWithResult(method, url string, req any) (t T, err error) {
	if method != "GET" {
		method = "POST"
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return
	}

	r, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return
	}

	httpResp, err := f.c.doPocketRequest(r)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	// レスポンスボディを読み込む
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	resp := &Resp[T]{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if !resp.Success {
		err = fmt.Errorf("request failed with code %d, message: %s", resp.Status, resp.Message)
		return
	}
	return resp.Content, nil
}
