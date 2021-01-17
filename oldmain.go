package main

//
//import (
//	"fmt"
//	"github.com/akamensky/argparse"
//	"k8sedgar/helpers"
//	"kubeIT/API/apistructs"
//	"kubeitcli/httpd"
//	"kubeitcli/httpd/requests"
//	"os"
//	"sync"
//	"time"
//)
//
////var rClient httpd.RequestClient
////var rURL string
//
//func main2() {
//
//	// Parse Config
//
//	rURL = os.Getenv("KUBEIT_URL")
//	if rURL == "" {
//		fmt.Println("Envvar Error: KUBEIT_URL must be set")
//		os.Exit(2)
//	}
//	rToken := os.Getenv("KUBEIT_TOKEN")
//	if rToken == "" {
//		fmt.Println("Envvar Error: KUBEIT_URL must be set")
//		os.Exit(2)
//	}
//
//	rClient = httpd.RequestClient{BaseURL: rURL}
//	rClient.Init(rToken)
//
//	commands := map[string]func(){
//		"create":  CreateHandler,
//		"get":  StatusHandler,
//		"version": VersionHandler,
//		"delete":  DeletionHandler,
//		"configure": ConfigurationHandler,
//		"help": HelpHandler,
//	}
//	commandFound := false
//
//	if len(os.Args) == 1 {
//		fmt.Println("Error: create or status Arguments (os.Args) must be specified.")
//		os.Exit(2)
//	}
//
//	for command, function := range commands {
//		if os.Args[1] == command {
//			function()
//			commandFound = true
//		}
//	}
//
//	if !commandFound {
//		fmt.Println("Unknown command (args[1]) function: use 'create' or 'get'")
//		os.Exit(2)
//	}
//}
//
//func VersionHandler() {
//	fmt.Println("Current version: v0.0.3-alpha-1")
//	os.Exit(0)
//}
//func CreateHandler() {
//	var resType *bool
//
//	if os.Args[2] == "workflow"{
//		*resType = true
//	}else if os.Args[2] == "scheme"{
//		*resType = false
//	}
//	if resType == nil {
//		fmt.Println("Unknown command (args[2]) function: use 'workflow' or 'scheme'")
//		os.Exit(2)
//	}
//
//	parser := argparse.NewParser("KubeIT create", "Handler for creating kubeIT workflows and schemes. A workflow is a specified instance of a scheme ")
//	configfile := parser.String("f", "File", &argparse.Options{Required: false, Help: "JSON-File with Job-Description"})
//	input := parser.String("i", "input", &argparse.Options{Required: false, Help: "Inputfile-name"})
//	output := parser.String("o", "output", &argparse.Options{Required: false, Help: "Output File or folder"})
//	watch := parser.Flag("w", "watch", &argparse.Options{Required: false, Help: "Detaches an independent daemon process, for processing results.", Default: false})
//
//	// Parse input
//	err := parser.Parse(append(os.Args[:1], os.Args[3:]...))
//	if err != nil {
//		fmt.Print(parser.Usage(err))
//		os.Exit(2)
//	}
//	parser.NewCommand(test)
//
//	iClient := httpd.RequestClient{BaseURL: rURL}
//
//	if iClient == nil {
//		fmt.Println("Error in initiating client")
//		os.Exit(2)
//	}
//
//	if *configfile != "" {
//		if *input != "" || *output != "" {
//			fmt.Println("input and output parameter are forbidden if -f is used")
//			fmt.Print(parser.Usage(err))
//			os.Exit(2)
//		}
//
//		*daemonize = true
//
//		jcfg := helpers.LoadJobConfig(*configfile)
//		jcfg.Parameters = append(jcfg.Parameters, "database=/database/edgartest")
//
//		for _, fileName := range jcfg.InputFiles {
//			SubmitForFile(jcfg.SplitSize, fileName, jcfg.OutputFolder+fileName+".result", jcfg.Parameters, iClient)
//		}
//	} else {
//		if *input == "" || *output == "" {
//			fmt.Println("if -f is not used -i AND -o must be specified")
//			fmt.Print(parser.Usage(err))
//			os.Exit(2)
//		}
//
//		BatchID = SubmitForFile(*splitsize, *input, *output, []string{"database=/database/edgartest"}, iClient)
//	}
//
//	fmt.Println("Waiting for results...(This may take a while)")
//
//	errcounter := 0
//	if !*daemonize {
//		for {
//			time.Sleep(10 * time.Second)
//			if BatchID == "" {
//				fmt.Println("Error in retrieving Batch, id: Abort")
//				os.Exit(2)
//			}
//			jobs, err := iClient.GetJobs()
//			if err != nil {
//				errcounter++
//				if errcounter == 10 {
//					fmt.Println("Error in communicating with local backend. Timeout, try again.")
//					os.Exit(2)
//				}
//				time.Sleep(30 * time.Second)
//				continue
//			}
//			for _, job := range jobs.Scheduled {
//				if job.BatchId == BatchID {
//					continue
//				}
//			}
//			for _, job := range jobs.Running {
//				if job.BatchId == BatchID {
//					continue
//				}
//			}
//			for _, job := range jobs.Finished {
//				if job.BatchId == BatchID {
//					fmt.Println("Batch finished, id: " + BatchID)
//					os.Exit(0)
//				}
//			}
//			for _, job := range jobs.Failed {
//				if job.BatchId == BatchID {
//					fmt.Println("Batch failed, id: " + BatchID)
//					os.Exit(2)
//				}
//			}
//
//		}
//	}
//}
