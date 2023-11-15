package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"bytes"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func GetIP(c *beego.Controller) string {
	//utils.GetIP(&c.Controller)
	//也可以直接用 c.Ctx.Input.IP() 取真实IP
	ip := c.Ctx.Request.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip = c.Ctx.Request.Header.Get("Remote_addr")
	if ip == "" {
		ip = c.Ctx.Request.RemoteAddr
	}
	return ip
}

func HttpGet(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		logs.Error("HttpGet error: ", err)
		return "", err
	}

	if resp == nil {
		return "", errors.New("返回对象为空")
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		logs.Debug("HttpGet result: ", result)
	} else {
		logs.Error("HttpGet error: ", err)
	}

	return result, nil
}

func HttpSimpleGet(reqUrl string) (response string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(reqUrl)

	defer resp.Body.Close()

	if err != nil {
		logs.Error("HttpGet error: ", err)
		return ""
	}

	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		//logs.Debug("HttpGet result: ", result)
	} else {
		logs.Error("HttpGet error: ", err)
	}

	return result
}

func HttpGet2(reqUrl string, headerParam map[string]string) (response string, err error) {
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		logs.Error("HttpGet error: ", err)
		return "", err
	}

	for k, v := range headerParam {
		req.Header.Set(k, v)
	}

	//resp, err := http.DefaultClient.Do(req)
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		logs.Error("HttpGet error: ", err)
		return "", err
	}

	defer resp.Body.Close()

	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		//logs.Debug("HttpGet result: ", result)
	} else {
		logs.Error("HttpGet error: ", err)
	}

	return result, nil
}

func HttpPostJson(reqUrl string, json string) (string, error) {
	resp, err := http.Post(reqUrl, "application/json", strings.NewReader(json))
	if err != nil {
		logs.Error("HttpPostJson error: ", err)
		return "", err
	}

	if resp == nil {
		return "", errors.New("返回对象为空")
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		//logs.Debug("HttpPostJson result: ", result)
	} else {
		logs.Error("HttpPostJson error: ", err)
	}

	return result, nil
}

func HttpPostJsonReturnByte(reqUrl string, json string) ([]byte, error) {
	resp, err := http.Post(reqUrl, "application/json", strings.NewReader(json))
	if err != nil {
		logs.Error("HttpPostJson error: ", err)
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("返回对象为空")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return body, err
		//logs.Debug("HttpPostJson result: ", result)
	} else {
		logs.Error("HttpPostJson error: ", err)
		return nil, err
	}

}

func HttpPost(url string, param map[string]string) (string, error) {

	var paramBuf bytes.Buffer
	paramBuf.WriteString("curTime=" + GetCurrentTime())
	for k, v := range param {
		paramBuf.WriteString("&" + k + "=" + v)
	}

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(paramBuf.String()))
	if err != nil {
		logs.Error("HttpPost error: ", err)
		return "", err
	}

	if resp == nil {
		return "", errors.New("返回对象为空")
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		logs.Debug("HttpPost result: ", result)
	} else {
		logs.Error("HttpPost error: ", err)
	}

	return result, nil
}

func HttpPost2(reqUrl string, param url.Values, headerParam map[string]string) (string, error) {

	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(param.Encode()))

	if err != nil {
		logs.Error("HttpPost error: ", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for k, v := range headerParam {
		req.Header.Set(k, v)
	}

	//resp, err := http.DefaultClient.Do(req)
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		logs.Error("HttpPost error: ", err)
		return "", err
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		logs.Debug("HttpPost result: ", result)
	} else {
		logs.Error("HttpPost error: ", err)
	}

	return result, nil

}

func HttpSimplePost(reqUrl string, param map[string]string) (string, error) {

	values := url.Values{}
	for k, v := range param {
		values.Set(k, v)
	}

	resp, err := http.Post(reqUrl, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		logs.Error("HttpSimplePost error: ", err)
		return "", err
	}

	if resp == nil {
		return "", errors.New("返回对象为空")
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		logs.Debug("HttpSimplePost result: ", result)
	} else {
		logs.Error("HttpSimplePost error: ", err)
	}

	return result, nil
}

func HttpRequestRaw(filePath string, https bool) (string, error) {
	client := &http.Client{}

	rawPayload, err := ioutil.ReadFile(filePath)
	if err != nil {
		//logs.Error("", err)
		return "", err
	}

	rawReq, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(rawPayload)))
	if err != nil && err != io.ErrUnexpectedEOF {
		return "", err
	}

	var reqUrl string
	if https {
		reqUrl = "https://" + rawReq.Host + rawReq.RequestURI
	} else {
		reqUrl = "http://" + rawReq.Host + rawReq.RequestURI
	}

	httpReq, err := http.NewRequest(rawReq.Method, reqUrl, rawReq.Body)
	httpReq.Header = rawReq.Header
	response, err := client.Do(httpReq)
	fmt.Println(err)
	bodyBytes, err := ioutil.ReadAll(response.Body)

	return string(bodyBytes), nil

}

func UrlEncode(input string) string {
	if IsEmpty(input) {
		return ""
	}
	return url.QueryEscape(input)
}

func UrlDecode(input string) string {
	if IsEmpty(input) {
		return ""
	}
	result, err := url.QueryUnescape(input)
	if err != nil {
		return input
	} else {
		return result
	}
}

func HttpGet3(reqUrl string, headerParam map[string]string) (response string, err error) {
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		logs.Error("HttpGet error: ", err)
		return "", err
	}

	for k, v := range headerParam {
		req.Header.Set(k, v)
	}

	//resp, err := http.DefaultClient.Do(req)
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		logs.Error("HttpGet error: ", err)
		return "", err
	}

	defer resp.Body.Close()

	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		//logs.Debug("HttpGet result: ", result)
	} else {
		logs.Error("HttpGet error: ", err)
	}

	return result, nil
}

func HttpDownload(reqUrl string, toPath string) bool {
	res, err := http.Get(reqUrl)
	if err != nil {
		logs.Error("下载资源出错：", err)
		return false
	}
	fmt.Println("res.StatusCode", res.StatusCode)

	workPath, _ := os.Getwd()
	fullPath := workPath + "/" + toPath
	os.MkdirAll(path.Dir(fullPath), os.ModePerm)
	file, err2 := os.Create(fullPath)
	if err2 != nil {
		logs.Error("下载资源出错：", err2)
		return false
	}

	_, err3 := io.Copy(file, res.Body)
	if err3 != nil {
		logs.Error("下载资源出错：", err3)
		return false
	}

	logs.Debug("下载完成：", fullPath)
	return true
}

func HttpSimplePostFile(reqUrl string, param map[string]string, fieldName, filePath string) (string, error) {

	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	//上传单个文件，如果上传多个文件可重复该段逻辑
	fileWriter, _ := bodyWriter.CreateFormFile(fieldName, FileFullName(filePath))
	file, _ := os.Open(filePath)
	defer file.Close()
	io.Copy(fileWriter, file)

	//如果同时含有表单参数
	for k, v := range param {
		_ = bodyWriter.WriteField(k, v)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(reqUrl, contentType, bodyBuffer)
	if err != nil {
		logs.Error("HttpPost error: ", err)
		return "", err
	}

	if resp == nil {
		return "", errors.New("返回对象为空")
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		logs.Debug("HttpPost result: ", result)
	} else {
		logs.Error("HttpPost error: ", err)
	}

	return result, nil

}

func HttpPostFile(reqUrl string, param url.Values, headerParam map[string]string, fieldName, filePath string) (string, error) {

	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	//上传单个文件，如果上传多个文件可重复该段逻辑
	fileWriter, _ := bodyWriter.CreateFormFile(fieldName, FileFullName(filePath))
	file, _ := os.Open(filePath)
	defer file.Close()
	io.Copy(fileWriter, file)

	//如果同时含有表单参数
	for k, v := range param {
		for _, _v := range v {
			_ = bodyWriter.WriteField(k, _v)
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest("POST", reqUrl, bodyBuffer)

	if err != nil {
		logs.Error("HttpPostFile error: ", err)
		return "", err
	}

	req.Header.Set("Content-Type", contentType)

	for k, v := range headerParam {
		req.Header.Set(k, v)
	}

	//resp, err := http.DefaultClient.Do(req)
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("HttpPostFile error: ", err)
		return "", err
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		logs.Debug("HttpPostFile result: ", result)
	} else {
		logs.Error("HttpPostFile error: ", err)
	}

	return result, nil

}
