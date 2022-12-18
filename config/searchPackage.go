package config

import (
	"log"
	"os"
)

var filterLicenseWhitelist []string
var filterLicenseBlacklist []string

var giteaInstances []string

var searchInGitea bool = false
var searchInGithub bool = true

func InitConfig(){
  if _, err := os.Stat("fpkg.cfg"); err != nil{
    _, err := os.Create("fpkg.cfg")
    if err != nil{
      log.Fatal("Erro em criar fpkg.cfg: ", err)
    }
    log.Println("Arquivo de configuração criado com sucesso")
  }
}
