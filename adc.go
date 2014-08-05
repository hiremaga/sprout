package sprout

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/howeyc/gopass"
)

type ADC struct {
	client   http.Client
	username string
	password string
}

func NewADC() (ADC, error) {
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return ADC{}, err
	}

	client := http.Client{Jar: jar}
	return ADC{client: client}, nil
}

func (adc *ADC) AskCredentials() {
	fmt.Print("ADC username: ")
	fmt.Scanf("%s\n", &adc.username)

	fmt.Print("ADC password: ")
	adc.password = string(gopass.GetPasswdMasked())
}

func (adc ADC) StartSession() error {
	loginUrl := "https://idmsa.apple.com/IDMSWebAuth/login?&appIdKey=891bd3417a7776362562d2197f89480a8547b108fd934911bcbea0110d07f757&path=%2F%2Fmembercenter%2Flogin.action"
	resp, err := adc.client.Get(loginUrl)
	defer resp.Body.Close()

	return err
}

func (adc ADC) Login() error {
	formUrl := "https://idmsa.apple.com/IDMSWebAuth/authenticate"
	values := url.Values{}
	values.Add("appleId", adc.username)
	values.Add("accountPassword", adc.password)

	resp, err := adc.client.PostForm(formUrl, values)
	defer resp.Body.Close()

	return err
}

func (adc ADC) DownloadXCodeCliTools(dmgFile io.Writer) error {
	dmgUrl := "https://developer.apple.com/downloads/download.action?path=Developer_Tools/command_line_tools_os_x_mavericks_for_xcode__march_2014/commandline_tools_os_x_mavericks_for_xcode__march_2014.dmg"
	resp, err := adc.client.Get(dmgUrl)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	_, err = io.Copy(dmgFile, resp.Body)
	return err
}
