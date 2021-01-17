package requests

import (
	"fmt"
	"kubeitcli/httpd"
)

func GetScheme(name string, c *httpd.RequestClient) (scheme httpd.SchemeInfo, failed bool, err error) {

	path := fmt.Sprintf("/v1/scheme?name=%v", name)

	err, statuscode := c.SendRequest("GET", path, nil, &scheme)

	if statuscode != 200 {
		failed = true
	} else {
		failed = false
	}

	return scheme, failed, err
}

func GetSchemes(c *httpd.RequestClient) (tempNames []string, failed bool, err error) {

	path := fmt.Sprintf("/v1/scheme")

	err, statuscode := c.SendRequest("GET", path, nil, &tempNames)

	if statuscode != 200 {
		failed = true
	} else {
		failed = false
	}

	return tempNames, failed, err
}
