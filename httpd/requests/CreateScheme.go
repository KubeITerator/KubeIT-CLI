package requests

import (
	"fmt"
	"io/ioutil"
	"kubeitcli/httpd"
	"os"
)

func CreateScheme(filepath, name string, c *httpd.RequestClient) (returndata map[string]string, failed bool, err error) {

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error in reading S3-File")
		fmt.Println(err.Error())
		os.Exit(2)
	}

	template := httpd.Template{
		Yaml: string(data),
		Name: name,
	}

	data2 := map[string]string{}
	err, statuscode := c.SendRequest("POST", "/v1/createtemplate", template, &data2)

	if statuscode != 200 {
		failed = true
	} else {
		failed = false
	}

	return data2, failed, err
}
