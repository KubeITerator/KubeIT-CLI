package requests

import (
	"fmt"
	"kubeitcli/httpd"
)

func PutDataS3(url string, c *httpd.RequestClient, data []byte) (string, error) {

	retst, err := c.S3UploadRequest(url, data)

	if err != nil {
		fmt.Println("Error in S3 Access")
		fmt.Println(err.Error())
	}

	return retst, err
}
