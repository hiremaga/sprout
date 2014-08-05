package main

import (
	"fmt"
	"os"

	"github.com/hiremaga/sprout/adc"
)

func main() {
	creds := adc.Credentials{}
	creds.Ask()

	session, err := adc.NewSession()
	if err != nil {
		fmt.Printf("Could not create ADC session: %v", err)
		os.Exit(1)
	}

	err = session.Start()
	if err != nil {
		fmt.Printf("Error loading ADC login page to start session: %v", err)
		os.Exit(1)
	}

	err = session.Login(creds)
	if err != nil {
		fmt.Printf("Error logging in to ADC: %v", err)
		os.Exit(1)
	}

	dmgFile, err := os.Create("/tmp/xcode-cli-tools.dmg")
	defer dmgFile.Close()

	if err != nil {
		fmt.Printf("Error creating temporary dmg: %v", err)
		os.Exit(1)
	}

	err = session.DownloadXCodeCliTools(dmgFile)
	if err != nil {
		fmt.Printf("Error download XCode CLI tools dmg: %v", err)
		os.Exit(1)
	}
}
