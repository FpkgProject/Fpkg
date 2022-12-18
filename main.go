package main

import (
	"flag"
	"fmt"

	"github.com/FpkgProject/Fpkg/config"
)

func contextSwitch(flagChoosed string, packageName string){
    if flagChoosed == "-install" || flagChoosed == "--install"{
        fmt.Printf("Install a %s package \n", packageName)
    }
    if flagChoosed == "-uninstall" || flagChoosed == "--uninstall"{
        fmt.Printf("Uninstall a %s package \n", packageName)
    }
}

func init_flags(){
    contextInstall := flag.String("install", "", "Install a new Package")
    contextUninstall := flag.String("uninstall", "", "Uninstall a new Package")
    flag.Parse()

    if *contextInstall != ""{
        contextSwitch("--install", *contextInstall)
    }

    if *contextUninstall != ""{
        contextSwitch("--uninstall", *contextUninstall)
    }
}

func main(){
  config.InitConfig()
  init_flags()
}
