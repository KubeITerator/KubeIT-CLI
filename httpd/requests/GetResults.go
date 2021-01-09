package requests

import (
	"fmt"
	"kubeitcli/httpd"
)

func GetResults(name string, c *httpd.RequestClient) (resp []httpd.ArtifactResponse, err error) {

	path := fmt.Sprintf("/v1/result?name=%v", name)

	err, _ = c.SendRequest("GET", path, nil, &resp)

	return resp, err
}
