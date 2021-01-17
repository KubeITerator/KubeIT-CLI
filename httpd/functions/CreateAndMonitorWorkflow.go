package functions

import (
	"fmt"
	"kubeitcli/ConfigHandler"
	"kubeitcli/httpd"
	"kubeitcli/httpd/requests"
	"os"
	"strconv"
	"strings"
	"time"
)

func ValidateParams(scheme ConfigHandler.Scheme, params []string) (missing []string, infilecount int) {

	for _, param := range params {
		pParam := strings.Split(param, "=")
		scheme.Parameters[pParam[0]] = pParam[1]
	}

	for key, value := range scheme.Parameters {
		if value == "" {
			if strings.Contains(key, "inputdata") {
				infilecount++
			} else {
				missing = append(missing, key)
			}
		}
	}
	return missing, infilecount
}

func CreateAndMonitorWorkflow(rClient *httpd.RequestClient, scheme ConfigHandler.Scheme, wfParameter, wfInputFiles []string, watch bool, output []string) {
	missing, infilecount := ValidateParams(scheme, wfParameter)

	if len(missing) != 0 {
		fmt.Println("[CREATE WORKFLOW] Missing parameters:")
		fmt.Println(fmt.Sprintf("[CREATE WORKFLOW] %v", missing))
		fmt.Println("[CREATE WORKFLOW] Abort: please specify missing parameters with -p \"paramName=value\"")
		os.Exit(2)
	}

	if infilecount != len(wfInputFiles) {
		fmt.Println(fmt.Sprintf("[CREATE WORKFLOW] Missing inputfiles Needed: %v/%v", len(wfInputFiles), infilecount))
		os.Exit(2)
	}

	if len(wfInputFiles) != 0 {
		fmt.Println("[CREATE WORKFLOW] Inputfile(s) detected, uploading to S3 and substituting (inputdata)")
		for index, fname := range wfInputFiles {
			url, err := UploadToS3(fname, rClient)
			if err != nil {
				fmt.Println("[CREATE WORKFLOW] Error in uploading file to S3")
				os.Exit(2)
			}
			if index == 0 {
				scheme.Parameters["inputdata"] = url
			} else {
				scheme.Parameters["inputdata"+strconv.Itoa(index)] = url
			}
		}
	}

	scheme.Parameters["scheme"] = scheme.RemoteName

	fmt.Println("[CREATE WORKFLOW] Creating workflow")

	failed, data, err := requests.ApplyWorkflow(scheme.Parameters, rClient)

	if err != nil {
		fmt.Println("[CREATE WORKFLOW] Failed to create new workflow")
		fmt.Println("[CREATE WORKFLOW] Error: " + err.Error())
		os.Exit(2)
	}

	if failed == true {
		fmt.Println("[CREATE WORKFLOW] Failed to create new workflow")
		fmt.Println("[CREATE WORKFLOW] Error: Non 200 http response")
		os.Exit(2)
	}

	fmt.Println("[CREATE WORKFLOW] Workflow creation successful, ID: " + data.WfName)

	if watch {
		fmt.Println("[CREATE WORKFLOW] Watch initiated:")

		backoff := 0
		for {
			time.Sleep(30 * time.Second)
			status, _, err := requests.GetStatus("", data.WfName, rClient)
			if err != nil {
				backoff++
			}
			if backoff > 10 {
				fmt.Println("[CREATE WORKFLOW] Failed to get status, after BackOffLimit (10/10)")
				os.Exit(2)
			}
			fmt.Println(fmt.Sprintf("[CREATE WORKFLOW] Name: %v, Status: %v, %v/%v Pods finished", data.WfName, status[0].Status, status[0].Status, status[0].Finished))
			if status[0].Status == "Succeeded" {
				break
			}
		}

		resp, err := requests.GetResults(data.WfName, rClient)
		if err != nil {
			fmt.Println("[CREATE WORKFLOW] Abort: Failed to get results for workflow: " + data.WfName)
			os.Exit(2)
		}

		if len(output) > 0 {
			if len(resp) != len(output) {
				fmt.Println(fmt.Sprintf("[CREATE WORKFLOW] Output file mapping failed: there must be exactly the same number of outputfiles specified as output artifacts created. Created:  %v, Specified: %v", len(resp), len(output)))
				os.Exit(2)
			}
			for index, artifact := range resp {
				fmt.Println("[CREATE WORKFLOW] Downloading artifact from: " + artifact.Pod)
				err = httpd.UntarUrlToSingleFile(artifact.URL, output[index])
				fmt.Println(fmt.Sprintf("[CREATE WORKFLOW] Download successful:  %v", output[index]))
			}

		} else {
			fmt.Println("[CREATE WORKFLOW] Workflow finished, Results: ")
			for index, artifact := range resp {
				fmt.Println(fmt.Sprintf("[CREATE WORKFLOW] Nr. %v, Job: %v, URL: %v", index, artifact.Pod, artifact.URL))
			}

		}

	}

}
