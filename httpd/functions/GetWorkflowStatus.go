package functions

import (
	"fmt"
	"kubeitcli/httpd"
	"kubeitcli/httpd/requests"
	"os"
)

func GetWorkflowStatus(project, wf string, rClient *httpd.RequestClient) {
	if wf != "" {
		project = ""
		fmt.Println("[GET WORKFLOW] Status for workflow: " + wf)
		status, failed, err := requests.GetStatus(project, wf, rClient)
		if err != nil {
			fmt.Println("[GET WORKFLOW] Failed to get workflow status")
			fmt.Println("[GET WORKFLOW] Error: " + err.Error())
			os.Exit(2)
		}

		if failed == true {
			fmt.Println("[GET WORKFLOW] Failed to get workflow status")
			fmt.Println("[GET WORKFLOW] Error: Non 200 http response")
			os.Exit(2)
		}
		fmt.Println(fmt.Sprintf("[GET WORKFLOW] Name: %v, Status: %v, %v/%v Pods finished", status[0].Workflow, status[0].Status, status[0].Finished, status[0].Running))

	} else if project != "" {
		status, failed, err := requests.GetStatus(project, wf, rClient)
		if err != nil {
			fmt.Println("[GET WORKFLOW] Failed to get workflow status")
			fmt.Println("[GET WORKFLOW] Error: " + err.Error())
			os.Exit(2)
		}

		if failed == true {
			fmt.Println("[GET WORKFLOW] Failed to get workflow status")
			fmt.Println("[GET WORKFLOW] Error: Non 200 http response")
			os.Exit(2)
		}

		fmt.Println("[GET WORKFLOW] Status for project: " + project)
		for _, stat := range status {
			fmt.Println(fmt.Sprintf("[GET WORKFLOW] Name: %v, Status: %v, %v/%v Pods finished", stat.Workflow, stat.Status, stat.Finished, stat.Running))
		}
	}
}
