package requests

import (
	"kubeitcli/httpd"
)

func InitS3(filename string, multi bool, c *httpd.RequestClient) (passkey string, err error) {

	data2 := httpd.S3InitResponse{}

	s3struct := httpd.S3InitRequest{
		Filename: filename,
		Multi:    multi,
	}

	err, _ = c.SendRequest("POST", "/s3/init", s3struct, &data2)

	if err != nil {
		return "", err
	}

	return data2.Passkey, err
}
