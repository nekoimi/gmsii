package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Get(url string, params map[string]interface{}) ([]byte, error) {
	url = url + parseToQueryString(params)
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func Post(url string, params map[string]interface{}) ([]byte, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	result, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(result))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func parseToQueryString(params map[string]interface{}) string {
	if params == nil {
		return ""
	}
	result := "?"
	for s, _ := range params {
		result += fmt.Sprintf("%s=%s&", s, params[s])
	}
	if strings.Contains(result, "&") {
		result = strings.TrimRight(result, "&")
	}
	return result
}

func parseQueryString(query string) map[string]interface{} {
	params := make(map[string]interface{})
	if query == "" {
		return params
	}

	if strings.HasPrefix(query, "?") {
		query = query[1:]
	}

	splitList := strings.Split(query, "&")
	for _, split := range splitList {
		if strings.Contains(split, "=") {
			splits := strings.Split(split, "=")
			params[splits[0]] = splits[1]
		} else {
			params[split] = ""
		}
	}

	return params
}
