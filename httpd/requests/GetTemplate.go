package requests

import (
	"fmt"
	"kubeitcli/httpd"
)

func GetTemplates(name string, c *httpd.RequestClient) (wfstatus []httpd.WFStatus, failed bool, err error) {

	path := fmt.Sprintf("/v1/template?name=%v", name)

	if name == "" {
		path = "/v1/template"
	}

	err, statuscode := c.SendRequest("GET", path, nil, wfstatus)

	if statuscode != 200 {
		failed = true
	} else {
		failed = false
	}

	return wfstatus, failed, err
}
