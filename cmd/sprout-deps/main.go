package main

import (
	"fmt"
	"os"

	"github.com/hiremaga/sprout"
)

func main() {
	adc, err := sprout.NewADC()
	if err != nil {
		fmt.Printf("Could not create ADC adc: %v", err)
		os.Exit(1)
	}

	err = adc.StartSession()
	if err != nil {
		fmt.Printf("Error loading ADC login page to start adc: %v", err)
		os.Exit(1)
	}

	adc.AskCredentials()
	err = adc.Login()
	if err != nil {
		fmt.Printf("Error logging in to ADC: %v", err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error creating temporary dmg: %v", err)
		os.Exit(1)
	}

	xcode := sprout.NewXcode(adc)
	err = xcode.DownloadCliTools()
	if err != nil {
		fmt.Printf("Error download XCode CLI tools dmg: %v", err)
		os.Exit(1)
	}
}
