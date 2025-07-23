package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

// Request 发送 HTTP 请求
func Request(method string, uri string, options map[string]interface{}, throwException bool) (interface{}, error) {
	var timeout time.Duration
	if t, ok := options["timeout"].(int); ok {
		timeout = time.Duration(t) * time.Second
	} else {
		timeout = 30 * time.Second // 默认超时为 30 秒
	}

	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	// 设置请求头或其他选项
	if headers, ok := options["headers"].(map[string]string); ok {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	// 处理请求体
	if jsonData, ok := options["json"]; ok {
		jsonValue, err := json.Marshal(jsonData)
		if err == nil {
			req.Body = io.NopCloser(bytes.NewBuffer(jsonValue))
			req.Header.Set("Content-Type", "application/json")
		}
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("HTTP request error:", err)
		if throwException {
			return nil, err
		}
		return nil, nil
	}

	// 立即关闭资源，并捕获关闭时可能出现的错误
	if err := resp.Body.Close(); err != nil {
		// 处理关闭错误，例如记录日志
		log.Printf("Error closing response body: %v", err)
	}

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusForbidden {
		var responseData interface{}
		if len(body) > 0 {
			if err := json.Unmarshal(body, &responseData); err != nil {
				log.Println("Error parsing JSON response:", err)
				if throwException {
					return nil, err
				}
				return nil, nil
			}
			log.Printf("Guzzle请求异常 (ClientException): %s, Status Code: %d, Response Data: %+v\n", uri, resp.StatusCode, responseData)
			return responseData, nil
		}
		if throwException {
			return nil, errors.New("received 403 status code with no response body")
		}
		return nil, nil
	}

	if IsJSON(body) {
		var responseData interface{}
		if err := json.Unmarshal(body, &responseData); err != nil {
			log.Println("Error parsing JSON response:", err)
			if throwException {
				return nil, err
			}
			return nil, nil
		}
		return responseData, nil
	}

	return string(body), nil
}
