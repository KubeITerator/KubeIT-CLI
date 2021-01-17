package ConfigHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kubeitcli/httpd"
	"kubeitcli/httpd/requests"
	"os"
	"path/filepath"
)

type Scheme struct {
	LocalName  string            `json:"local_name"`
	RemoteName string            `json:"remote_name"`
	Parameters map[string]string `json:"parameters"`
}

type Config struct {
	URL     string   `json:"url"`
	Token   string   `json:"token"`
	Schemes []Scheme `json:"schemes"`
}

type ConfigHandler struct {
	File   *os.File
	Config Config
}

func (c *ConfigHandler) LoadConfig() error {
	fileSteam, err := ioutil.ReadAll(c.File)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileSteam, &c.Config)

	return err
}

func (c *ConfigHandler) SaveConfig() error {

	bytes, err := json.Marshal(c.Config)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.File.Name(), bytes, 0644)

	return err
}

func (c *ConfigHandler) SchemeExist(localName string) (exists bool) {
	for _, scheme := range c.Config.Schemes {
		if scheme.LocalName == localName {
			return true
		}
	}
	return false
}

func (c *ConfigHandler) DeleteLocalScheme(localName string) error {
	for index, scheme := range c.Config.Schemes {
		if scheme.LocalName == localName {
			c.Config.Schemes = append(c.Config.Schemes[:index], c.Config.Schemes[index+1:]...)
		}
	}
	err := c.SaveConfig()
	return err
}

func (c *ConfigHandler) ConfigureConDialogue() (err error) {
	var filePath string
	fmt.Println("This is the standard configuration interface for KubeIT.")
	fmt.Println("All specified parameters will be stored in ~/.kubeit/config.json if no other path is specified.")
	fmt.Println("Config entries marked as (optional) can be skipped by hitting enter, (required) parameters must be specified.")
	path := Ask("path", "(optional) Please enter a new config path (default: '~/.kubeit/config.json'): ", false)

	hdir, _ := os.UserHomeDir()
	if path == "" {
		filePath, err = filepath.Abs(hdir + "/.kubeit/config.json")
		if err != nil {
			return err
		}
	}

	_ = os.MkdirAll(filepath.Dir(filePath), 0744)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0744)
	if err != nil {
		return err
	}
	url := Ask("KubeIT URL", "(required) Please enter a valid kubeIT URL (example: https://kubeit.example.com): ", true)
	token := Ask("KubeIT Access Token", "(required) Please enter a valid kubeIT access token: ", true)

	testrClient := httpd.RequestClient{}
	testrClient.Init(url, token)
	_, failed, err := requests.GetSchemes(&testrClient)
	if err != nil {
		fmt.Println("Error in validating login from server: " + err.Error())
		os.Exit(2)
	} else if failed {
		fmt.Println("Error in validating login: Wrong AUTH token")
		os.Exit(2)
	}

	c.File = file

	c.Config = Config{
		URL:     url,
		Token:   token,
		Schemes: []Scheme{},
	}
	err = c.SaveConfig()

	fmt.Println("Success: New config created and validated")

	return err
}
