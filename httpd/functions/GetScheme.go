package functions

import (
	"fmt"
	"kubeitcli/ConfigHandler"
	"kubeitcli/httpd"
	"kubeitcli/httpd/requests"
	"os"
)

func GetScheme(name string, local bool, cHandler *ConfigHandler.ConfigHandler, rClient *httpd.RequestClient) {

	if local {
		fmt.Println("[GET SCHEME] Requesting local scheme(s)")
		if name != "" {
			for _, scheme := range cHandler.Config.Schemes {
				if scheme.LocalName == name {

					fmt.Println(fmt.Sprintf("[GET SCHEME] Local name: '%v', remote name: '%v'", scheme.LocalName, scheme.RemoteName))
					fmt.Println(fmt.Sprintf("[GET SCHEME] Parameters/defaults"))
					for k, v := range scheme.Parameters {
						fmt.Println(fmt.Sprintf("[GET SCHEME] Parameter: '%v', default: '%v'", k, v))
					}

				}
			}
		} else {
			fmt.Println("[GET SCHEME] All local scheme(s) with remoteNames")
			for _, scheme := range cHandler.Config.Schemes {
				fmt.Println(fmt.Sprintf("[GET SCHEME] Local name: '%v', remote name: '%v'", scheme.LocalName, scheme.RemoteName))
			}

		}

	} else {
		fmt.Println("[GET SCHEME] Requesting server schemes, use -l for local schemes")

		if name != "" {

			scheme, failed, err := requests.GetScheme(name, rClient)

			if err != nil {
				fmt.Println("[GET SCHEME] Failed with error: " + err.Error())
				os.Exit(2)
			}

			if failed {
				fmt.Println("[GET SCHEME] Failed with non 200 exit status")
				os.Exit(2)
			}

			fmt.Println(fmt.Sprintf("[GET SCHEME] Name: '%v'", scheme.Name))
			fmt.Println(fmt.Sprintf("[GET SCHEME] Parameters/defaults"))
			for k, v := range scheme.Parameters {
				fmt.Println(fmt.Sprintf("[GET SCHEME] Parameter: '%v', default: '%v'", k, v))
			}

		} else {

			schemes, failed, err := requests.GetSchemes(rClient)

			if err != nil {
				fmt.Println("[GET SCHEME] Failed with error: " + err.Error())
				os.Exit(2)
			}

			if failed {
				fmt.Println("[GET SCHEME] Failed with non 200 exit status")
				os.Exit(2)
			}

			fmt.Println("[GET SCHEME] All remote scheme(s):")
			for _, scheme := range schemes {
				fmt.Println(fmt.Sprintf("[GET SCHEME] Name: '%v'", scheme))
			}

		}
	}

}
