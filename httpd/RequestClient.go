package httpd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type RequestClient struct {
	BaseURL    string
	httpClient *http.Client
	token      string
}

func (c *RequestClient) Init(token string) {
	c.httpClient = &http.Client{}
	c.token = token
}

func (c *RequestClient) SendRequest(method, path string, requestBody, returnBody interface{}) (err error, statuscode int) {
	req, err := c.newRequest(method, path, requestBody)
	if err != nil {
		return err, 0
	}

	// reponse, error
	resp, err := c.do(req, returnBody)

	if err != nil && resp == nil {
		fmt.Println("Error in Sending request, abort, please check your connection")
		os.Exit(2)
	}

	return err, resp.StatusCode
}

func (c *RequestClient) S3UploadRequest(path string, body []byte) (string, error) {

	req, err := http.NewRequest("PUT", path, bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error in creating S3-Request")
		fmt.Println(err.Error())
		return "", err
	}

	resp, err := c.do(req, nil)

	if err != nil {
		fmt.Println("Error in executing S3-Request")
		fmt.Println(err.Error())
		return "", err
	}

	retst, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	return string(retst), err
}

func (c *RequestClient) newRequest(method, path string, body interface{}) (*http.Request, error) {

	u := c.BaseURL + path
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			fmt.Println("Error Encoding new Request")
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		fmt.Println("Error NewRequest")
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Token", c.token)
	return req, nil
}
func (c *RequestClient) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			if resp.StatusCode != 504 {
				fmt.Println("Error while decoding response to json")
			}
		}
	}
	return resp, err
}
