package functions

import (
	"fmt"
	"kubeitcli/httpd"
	"kubeitcli/httpd/requests"
	"os"
)

func DeleteWorkflows(wf, project string, rClient *httpd.RequestClient) {

	if wf == "" {
		fmt.Println("[DELETE WORKFLOW] Deleting all workflows from project: " + project)
		failed, err := requests.DeleteWorkflows(project, "", rClient)

		if err != nil {
			fmt.Println("[DELETE WORKFLOW] Failed with error: " + err.Error())
			os.Exit(2)
		}

		if failed {
			fmt.Println("[DELETE WORKFLOW] Failed with non 200 exit status")
			os.Exit(2)
		}

		fmt.Println("[DELETE WORKFLOW] Project successful deleted")

	} else {

		fmt.Println("[DELETE WORKFLOW] Deleting workflow with name: " + wf)

		failed, err := requests.DeleteWorkflows("", wf, rClient)

		if err != nil {
			fmt.Println("[DELETE WORKFLOW] Failed with error: " + err.Error())
			os.Exit(2)
		}

		if failed {
			fmt.Println("[DELETE WORKFLOW] Failed with non 200 exit status")
			os.Exit(2)
		}

		fmt.Println("[DELETE WORKFLOW] Workflow successful deleted")
	}
}
