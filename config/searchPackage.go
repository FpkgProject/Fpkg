package config

import (
	"fmt"
	"os"

	"github.com/FpkgProject/Fpkg/utils"
)

var filterLicenseWhitelist []string
var filterLicenseBlacklist []string

var giteaInstances []string

var searchInGitea bool = false
var searchInGithub bool = true

func readFile() {
	fmt.Println("Lendo o arquivo fpkg.cfg")
}

// Warning (the fpkg.cfg need exists in /etc/fpkg.cfg)
func InitConfig() {
	if utils.VerifyIfFileExists("fpkg.cfg") == false {
		os.Create("fpkg.cfg")
	} else {
		readFile()
	}
}
