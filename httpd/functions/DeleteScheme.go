package functions

import (
	"fmt"
	"kubeitcli/ConfigHandler"
	"os"
)

func DeleteScheme(cHandler *ConfigHandler.ConfigHandler, schemeName string) {

	fmt.Println("[DELETE SCHEME] Deleting local scheme, name: " + schemeName)
	fmt.Println("[DELETE SCHEME] Only local schemes can be deleted, for remote deletion: contact your admin")
	err := cHandler.DeleteLocalScheme(schemeName)
	if err != nil {
		fmt.Println("[DELETE SCHEME] Error in deleting local scheme")
		fmt.Println("[DELETE SCHEME] Error: " + err.Error())
		os.Exit(2)
	}
	fmt.Println("[DELETE SCHEME] Deletion successful")
}
