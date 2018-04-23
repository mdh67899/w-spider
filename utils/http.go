package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	tr = &http.Transport{
		MaxIdleConns:          10,
		IdleConnTimeout:       3 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
		MaxIdleConnsPerHost:   1000,
		DisableCompression:    true,
	}

	client = &http.Client{Transport: tr}
)

func PostJSON(url string, v interface{}) (response []byte, err error) {
	var bs []byte
	bs, err = json.Marshal(v)
	if err != nil {
		return
	}

	bf := bytes.NewBuffer(bs)

	req, err := http.NewRequest("POST", url, bf)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.Body != nil {
		defer resp.Body.Close()
		response, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
	}

	if resp.StatusCode != 200 {
		err = errors.New("status code not equals 200")
	}

	return
}
