package requests

import (
	"fmt"
	"kubeitcli/httpd"
)

func DeleteWorkflows(project, workflow string, c *httpd.RequestClient) (failed bool, err error) {

	path := fmt.Sprintf("/v1/delete?project=%v", project)

	if project == "" {
		path = fmt.Sprintf("/v1/delete?workflow=%v", workflow)
	}

	err, statuscode := c.SendRequest("GET", path, nil, nil)

	if statuscode != 200 {
		failed = true
	} else {
		failed = false
	}

	return failed, err
}
