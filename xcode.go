package sprout

import "os"

type Xcode struct {
	adc ADC
}

func NewXcode(adc ADC) Xcode {
	return Xcode{adc}
}
func (x Xcode) DownloadCliTools() error {
	dmgUrl := "https://developer.apple.com/downloads/download.action?path=Developer_Tools/command_line_tools_os_x_mavericks_for_xcode__march_2014/commandline_tools_os_x_mavericks_for_xcode__march_2014.dmg"
	dmgFile, err := os.Create("/tmp/xcode-cli-tools.dmg")
	defer dmgFile.Close()

	if err != nil {
		return err
	}

	return x.adc.Download(dmgUrl, dmgFile)
}
