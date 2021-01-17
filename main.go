package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"kubeitcli/ConfigHandler"
	"kubeitcli/httpd"
	"os"
)

var rClient httpd.RequestClient
var cHandler ConfigHandler.ConfigHandler

func main() {
	parser := argparse.NewParser("kubeit", "Handler for creating kubeIT workflows and schemes. A workflow is a specified instance of a scheme ")

	// SubCommands
	// main
	cmdCreate := parser.NewCommand("create", "Create an API object")
	cmdConfigure := parser.NewCommand("configure", "Run on first startup to configure, KubeIT URL and access token, as well as new local schemes (-s)")
	cmdGet := parser.NewCommand("get", "Get an specific API object")
	cmdVersion := parser.NewCommand("version", "Prints the current version number for KubeIT")
	cmdDelete := parser.NewCommand("delete", "Deletes an KubeIT API object or local scheme")

	// create
	cmdCreateWorkflow := cmdCreate.NewCommand("workflow", "create an object of type workflow")
	cmdCreateScheme := cmdCreate.NewCommand("scheme", "create an object of type scheme")
	cmdCreateS3 := cmdCreate.NewCommand("S3", "Upload a new object to S3 and retrieve an URL")

	// get
	cmdGetWorkflow := cmdGet.NewCommand("workflow", "Gets the current status of a workflow")
	cmdGetScheme := cmdGet.NewCommand("scheme", "Gets the currently available schemes from the server, or detailed information about a specific scheme")
	cmdGetResults := cmdGet.NewCommand("result", "Gets the URL(s) to results of a specified workflow")

	// delete
	cmdDeleteWorkflow := cmdDelete.NewCommand("workflow", "Deletes a specified workflow")
	cmdDeleteScheme := cmdDelete.NewCommand("scheme", "Deletes a local scheme")

	// Parameters:
	// global
	configFile := parser.String("c", "config", &argparse.Options{Required: false, Help: "Path to config file, Default: '~/.kubeit/config.json'"})

	// create Workflow
	wfInputFiles := cmdCreateWorkflow.StringList("i", "inputfile", &argparse.Options{Required: false, Help: "Inputfile(s) are automatically uploaded to S3 and substituted as kubeit.input.inputdata1-x, the first inputfile is mapped to kubeit.input.inputdata"})
	wfWatchFlag := cmdCreateWorkflow.Flag("w", "watch", &argparse.Options{Required: false, Help: "Watch can be used to monitor a workflows execution, in conjuction with -o the results will automatically be downloaded"})
	wfOutputFile := cmdCreateWorkflow.String("o", "output", &argparse.Options{Required: false, Help: "(optional) outputfile can be used in conjuction with the -w/--watch flag to automatically download results"})
	wfParameter := cmdCreateWorkflow.StringList("p", "parameter", &argparse.Options{Required: false, Help: "Parameter are all parameters that dont have a remote or local default, they must be specified with the following syntax: parameterName='parameterValue' "})
	// create Scheme
	cmdCreateScheme.String("n", "name", &argparse.Options{Required: true, Help: "Remote name for the new scheme"})
	cmdCreateScheme.String("f", "File", &argparse.Options{Required: true, Help: "YAML kubeIT workflow scheme"})
	// create S3
	cmdCreateS3.File("f", "File", os.O_RDWR, 0644, &argparse.Options{Required: true, Help: "File that should be uploaded to S3"})

	// get Workflow(s)
	cmdGetWorkflow.String("n", "name", &argparse.Options{Required: false, Help: "Name of a workflow, gets the current status of the workflow"})
	cmdGetWorkflow.String("p", "project", &argparse.Options{Required: false, Help: "Name of a project, gets the current status of all workflows in a project"})
	// get Scheme(s)
	cmdGetScheme.String("n", "name", &argparse.Options{Required: true, Help: "Name of a scheme, gets the current description of the specified scheme"})
	cmdGetScheme.Flag("l", "local", &argparse.Options{Default: true, Help: "Get local schemes if true, else get remote schemes"})
	// get Results
	cmdGetResults.String("n", "name", &argparse.Options{Required: true, Help: "Name of a running workflow, gets the results of a finished workflow"})

	// delete Workflow
	cmdDeleteWorkflow.String("n", "name", &argparse.Options{Required: false, Help: "Name of a workflow, deletes the workflow"})
	cmdDeleteWorkflow.String("p", "project", &argparse.Options{Required: false, Help: "Name of a project, deletes all workflows in a project"})
	// delete Scheme
	cmdDeleteScheme.String("n", "name", &argparse.Options{Required: true, Help: "Name of a local scheme workflow, deletes the specified local scheme"})

	// configure
	cmdConfigureScheme := cmdConfigure.Flag("s", "scheme", &argparse.Options{Required: false, Help: "Configure schemes to substitute serverside schemes with local defaults"})

	// Parse OS args
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(2)
	}
	cHandler = ConfigHandler.ConfigHandler{}

	if cmdConfigure.Happened() {
		if !*cmdConfigureScheme {
			err = cHandler.ConfigureConDialogue()
			if err != nil {
				fmt.Print(parser.Usage(err))
				os.Exit(2)
			}
		}
	}
	fmt.Println(*configFile)
	if *configFile != "" {
		cHandler.File, err = os.Open(*configFile)
	} else {
		hdir, _ := os.UserHomeDir()
		cHandler.File, err = os.Open(hdir + "/.kubeit/config.json")
	}
	err = cHandler.LoadConfig()
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No config file found, please use 'kubeit configure' to create a new one, or use the -c flag to specify another config.")
			os.Exit(0)
		}
		fmt.Println(fmt.Sprintf("Error in loading configfile: %v", configFile))
		fmt.Println(fmt.Sprintf("Error: %v", err.Error()))
		os.Exit(2)
	}

	rClient = httpd.RequestClient{}
	rClient.Init(cHandler.Config.URL, cHandler.Config.Token)

	if *cmdConfigureScheme {
		ConfigHandler.ConfigureSchemeDialogue(&rClient, &cHandler)
	}

	if cmdVersion.Happened() {
		fmt.Println("Current version: v0.0.3-alpha-1")
		os.Exit(0)
	} else if cmdCreate.Happened() {
		// CreateHandling
		if cmdCreateWorkflow.Happened() {

			wfInputFiles
			wfWatchFlag
			wfOutputFile
			wfParameter

		} else if cmdCreateS3.Happened() {

		} else if cmdCreateScheme.Happened() {

		}

	} else if cmdGet.Happened() {
		// GetHandling
	} else if cmdDelete.Happened() {
		// DeleteHandling
	}

}
