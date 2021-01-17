package functions

import (
	"fmt"
	"kubeitcli/httpd"
	"kubeitcli/httpd/requests"
	"os"
)

func CreateRemoteScheme(name, yamlFile string, rclient *httpd.RequestClient) {

	fmt.Println("[CREATE SCHEME] Creating remote scheme with name: " + name)
	_, failed, err := requests.CreateScheme(yamlFile, name, rclient)
	if err != nil {
		fmt.Println("[CREATE SCHEME] Failed with error: " + err.Error())
		os.Exit(2)
	}

	if failed {
		fmt.Println("[CREATE SCHEME] Failed with non 200 exit status")
		os.Exit(2)
	}

	fmt.Println("[CREATE SCHEME] Remote scheme created: " + name)

}
