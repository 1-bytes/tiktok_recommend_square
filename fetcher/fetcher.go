package fetcher

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Json 以 Json 格式获取网页内容.
func Json(method string, api string, body string, header *http.Header) (*http.Response, error) {
	u, err := url.Parse(api)
	if err != nil {
		return nil, err
	}
	// 加入对应的 header
	if header.Get("Host") != u.Host {
		header.Add("Host", u.Host)
		header.Add("Content-Length", strconv.Itoa(len(body)))
		header.Add("Content-type", "application/json")
	}
	// 构造请求
	req, err := http.NewRequest(method, api, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := sendRequest(req, header)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, unintended status code: %d", resp.StatusCode)
	}
	return resp, nil
}

// FormData 以 form-data 格式获取网页内容.
func FormData(method string, api string, body *map[string]string, header *http.Header) (*http.Response, error) {
	u, err := url.Parse(api)
	if err != nil {
		return nil, err
	}
	if header.Get("Host") != u.Host {
		header.Add("Host", u.Host)
	}
	// 构造 form-data
	data := new(bytes.Buffer)
	w := multipart.NewWriter(data)
	for key, value := range *body {
		err = w.WriteField(key, value)
		if err != nil {
			return nil, err
		}
	}
	_ = w.Close()
	// 加入对应的header
	header.Add("Content-Length", strconv.Itoa(len(w.Boundary())))
	header.Add("Content-Type", w.FormDataContentType())

	// 构造请求
	req, err := http.NewRequest(method, api, data)
	if err != nil {
		return nil, err
	}

	resp, err := sendRequest(req, header)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, unintended status code: %d", resp.StatusCode)
	}
	return resp, nil
}

// sendRequest 发出请求.
func sendRequest(req *http.Request, header *http.Header) (*http.Response, error) {
	if header.Get("User-Agent") == "" {
		header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36")
	}
	req.Header = *header
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return client.Do(req)
}
