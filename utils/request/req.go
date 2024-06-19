package request

import (
	"bytes"
	"crypto/tls"
	"echo/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var client *http.Client

func init() {
	//def := http.DefaultTransport
	//defPot, ok := def.(*http.Transport)
	defPot := new(http.Transport)
	//if !ok {
	//	panic("init transport出错")
	//}
	defPot.MaxIdleConns = 100
	defPot.MaxIdleConnsPerHost = 100
	defPot.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	//defPot.IdleConnTimeout = 60
	client = &http.Client{
		Timeout:   time.Second * time.Duration(30),
		Transport: defPot,
	}
}

func Get(url string, header map[string]string, params map[string]interface{}) ([]byte, error) {
	//client := &http.Client{
	//	Timeout: time.Second * time.Duration(20),
	//}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//logger.Info("[ERROR] new request 请求出错:", err.Error())
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	q := req.URL.Query()
	if params != nil {
		for Key, val := range params {
			v, _ := utils.ToString(val)
			q.Add(Key, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	r, err := client.Do(req)
	if err != nil {
		//logger.Info("do 发生错误:", err.Error())
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		//logger.Info(r.StatusCode)
		return nil, err
	}
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//logger.Info("ioutil reader:", err.Error())
		return nil, err
	}
	return bb, nil
}

func Post(url string, header map[string]string, param map[string]interface{}) ([]byte, error) {
	//client := &http.Client{
	//	Timeout: time.Second * time.Duration(20),
	//}
	dd, _ := json.Marshal(param)
	re := bytes.NewReader(dd)
	req, err := http.NewRequest("POST", url, re)
	if err != nil {
		//logger.Info("[ERROR] new request 请求出错:", err.Error())
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	r, err := client.Do(req)
	if err != nil {
		//logger.Info("do 发生错误:", err.Error())
		return nil, err
	}
	defer r.Body.Close()
	//if r.StatusCode != 200 {
	//	return nil,err
	//}
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//logger.Info("ioutil reader:", err.Error())
		return nil, err
	}
	return bb, nil
}

func PostMultipart(urlStr string, header map[string]string, param map[string]interface{}) ([]byte, error) {
	values := url.Values{}
	for k, v := range param {
		val, _ := utils.ToString(v)
		values[k] = []string{val}
	}
	formData := values
	req, err := http.NewRequest("POST", urlStr, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		//logger.Info("[ERROR] new request 请求出错:", err.Error())
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	r, err := client.Do(req)
	if err != nil {
		//logger.Info("do 发生错误:", err.Error())
		return nil, err
	}
	defer r.Body.Close()
	//if r.StatusCode != 200 {
	//	return nil,err
	//}
	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//logger.Info("ioutil reader:", err.Error())
		return nil, err
	}
	return bb, nil
}
