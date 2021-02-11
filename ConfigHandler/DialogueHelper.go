package ConfigHandler

import (
	"bufio"
	"fmt"
	"kubeitcli/httpd"
	"kubeitcli/httpd/requests"
	"os"
	"strings"
)

func Ask(pname, question string, required bool) (answer string) {

	fmt.Println(question)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer = scanner.Text()
	if answer == "" {
		if required {
			fmt.Println("ConfigEntry: " + pname + " must be specified. Try again.")
			answer = Ask(pname, question, required)
		}
	}
	return answer

}

func AskYN(question string) (yesorno bool) {

	answer := ""
	fmt.Println(question)
	fmt.Scanln(&answer)
	if answer == "" {
		return true
	} else {
		switch strings.ToLower(answer) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Unknown input: Valid inputs are Y/y/yes or N/n/no")
			yesorno = AskYN(question)
		}
	}

	return yesorno
}

func ConfigureSchemeDialogue(rClient *httpd.RequestClient, configHandler *ConfigHandler) {
	fmt.Println("[CONFIG SCHEME] This dialogue is used to preconfigure remote schemes with local defaults and versions.")

	schemes, failed, err := requests.GetSchemes(rClient)
	if err != nil {
		fmt.Println("[CONFIG SCHEME] Error in requesting schemes from Server: " + err.Error())
		os.Exit(2)
	} else if failed {
		fmt.Println("[CONFIG SCHEME] Error in requesting schemes from Server: Non 200 Status")
		os.Exit(2)
	}
	fmt.Println("[CONFIG SCHEME] The following schemes are currently installed on the server: ")
	for index, scheme := range schemes {
		fmt.Println(fmt.Sprintf("[CONFIG SCHEME] Number: %v, Name: %v", index, scheme))
	}

	qdef := AskYN("[CONFIG SCHEME] Do you want to configure one of these schemes with local defaults ? [Y/n]")
	if !qdef {
		fmt.Println("[CONFIG SCHEME] No configuration requested. Exit")
		os.Exit(0)
	}

	schemeName := Ask("scheme", "[CONFIG SCHEME] (required) Please enter a valid remote scheme name: ", true)

	for _, scheme := range schemes {
		if schemeName == scheme {
			goto validation
		}
	}
	fmt.Println("[CONFIG SCHEME] Unknown schemename: " + schemeName + " -> abort.")
	os.Exit(2)

validation:
	info, failed, err := requests.GetScheme(schemeName, rClient)
	if err != nil {
		fmt.Println("[CONFIG SCHEME] Error in requesting schemes from Server: " + err.Error())
		os.Exit(2)
	} else if failed {
		fmt.Println("[CONFIG SCHEME] Error in requesting schemes from Server: Non 200 Status")
		os.Exit(2)
	}

	fmt.Println("[CONFIG SCHEME] You can specify a different schemeName locally. This can be used to create multiple local presets for the same remote scheme, by default the remote scheme name will be used")

	altName := Ask("local scheme name", "[CONFIG SCHEME] (optional) Specify an alternative local scheme name: ", false)
	if configHandler.SchemeExist(altName) {
		fmt.Println(fmt.Sprintf("[CONFIG SCHEME] Scheme already exists. Use a different Name or delete it with: 'kubeit delete scheme -l -n %v'", altName))
	}
	if altName == "" {
		altName = schemeName
	}

	fmt.Println("[CONFIG SCHEME] The scheme has the following parameters, please set local defaults.")
	fmt.Println("[CONFIG SCHEME] If local (required) defaults are left empty, they need to be specified on workflow creation with -p name=value")
	fmt.Println("[CONFIG SCHEME] Except: inputdata, this special parameter can be automatically substituted by -i FileName")

	saveparams := make(map[string]string)
	for name, defaultName := range info.Parameters {
		if defaultName == "" {
			saveparams[name] = Ask(name, fmt.Sprintf("[CONFIG SCHEME] (required) Value for ParameterName: %v", name), false)
		} else {
			answer := Ask(name, fmt.Sprintf("[CONFIG SCHEME] (optional) Value for ParameterName: %v, Remote Default: %v", name, defaultName), false)
			if answer != "" {
				saveparams[name] = answer
			} else {
				saveparams[name] = defaultName
			}
		}
	}

	configHandler.Config.Schemes = append(configHandler.Config.Schemes, Scheme{
		LocalName:  altName,
		RemoteName: schemeName,
		Parameters: saveparams,
	})

	err = configHandler.SaveConfig()

	if err == nil {
		fmt.Println(fmt.Sprintf("[CONFIG SCHEME] Local scheme: %v successfully configured", altName))
	}
}
