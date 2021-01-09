package requests

import (
	"fmt"
	"kubeitcli/httpd"
)

func GetStatus(project, workflow string, c *httpd.RequestClient) (wfstatus []httpd.WFStatus, failed bool, err error) {

	path := fmt.Sprintf("/v1/status?project=%v", project)

	if project == "" {
		path = fmt.Sprintf("/v1/status?workflow=%v", workflow)
	}

	err, statuscode := c.SendRequest("GET", path, nil, wfstatus)

	if statuscode != 200 {
		failed = true
	} else {
		failed = false
	}

	return wfstatus, failed, err
}
