package requests

import (
	"fmt"
	"kubeitcli/httpd"
)

func FinishS3Upload(key string, c *httpd.RequestClient) (failed bool, err error) {

	data2 := httpd.URLResponse{}

	path := fmt.Sprintf("/s3/finish?key=%v", key)

	err, statuscode := c.SendRequest("GET", path, nil, &data2)

	if statuscode != 200 {
		failed = true
	} else {
		failed = false
	}
	return failed, err
}
