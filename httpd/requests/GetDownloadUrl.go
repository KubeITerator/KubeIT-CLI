package requests

import (
	"fmt"
	"kubeitcli/httpd"
)

func GetDownload(key string, c *httpd.RequestClient) (url string, err error) {

	data2 := httpd.URLResponse{}

	path := fmt.Sprintf("/s3/download?key=%v", key)

	err, _ = c.SendRequest("GET", path, nil, &data2)

	return data2.URL, err
}
