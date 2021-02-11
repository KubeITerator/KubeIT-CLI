package functions

import (
	"fmt"
	"kubeitcli/httpd"
	"kubeitcli/httpd/requests"
	"os"
)

func GetResults(name string, rClient *httpd.RequestClient) {

	status, _, _ := requests.GetStatus("", name, rClient)

	if status[0].Status != "Succeeded" {
		fmt.Println("[GET RESULTS] Error: Workflow not finished, status: " + status[0].Statusmessage)
		os.Exit(2)
	}

	resp, err := requests.GetResults(name, rClient)
	if err != nil {
		fmt.Println("[GET RESULTS] Abort: Failed to get results for workflow: " + name)
		os.Exit(2)
	}

	fmt.Println("[GET RESULTS] Workflow " + name + "finished, Results: ")
	for index, artifact := range resp {
		fmt.Println(fmt.Sprintf("[GET RESULTS] Nr. %v, Job: %v, URL: %v", index, artifact.Pod, artifact.URL))
	}

}
