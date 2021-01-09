package requests

import (
	"kubeitcli/httpd"
)

func ApplyWorkflow(parameters map[string]string, c *httpd.RequestClient) (failed bool, data2 httpd.ApplyReturn, err error) {

	err, statuscode := c.SendRequest("POST", "/v1/apply", parameters, &data2)

	if statuscode != 200 {
		failed = true
	} else {
		failed = false
	}

	return failed, data2, err
}
