package adc

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Session struct {
	client http.Client
}

func NewSession() (Session, error) {
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return Session{}, err
	}

	client := http.Client{Jar: jar}
	return Session{client}, nil
}

func (s Session) Start() error {
	loginUrl := "https://idmsa.apple.com/IDMSWebAuth/login?&appIdKey=891bd3417a7776362562d2197f89480a8547b108fd934911bcbea0110d07f757&path=%2F%2Fmembercenter%2Flogin.action"
	resp, err := s.client.Get(loginUrl)
	defer resp.Body.Close()

	return err
}

func (s Session) Login(creds Credentials) error {
	formUrl := "https://idmsa.apple.com/IDMSWebAuth/authenticate"
	values := url.Values{}
	values.Add("appleId", creds.username)
	values.Add("accountPassword", creds.password)

	resp, err := s.client.PostForm(formUrl, values)
	defer resp.Body.Close()

	return err
}

func (s Session) DownloadXCodeCliTools(dmgFile io.Writer) error {
	dmgUrl := "https://developer.apple.com/downloads/download.action?path=Developer_Tools/command_line_tools_os_x_mavericks_for_xcode__march_2014/commandline_tools_os_x_mavericks_for_xcode__march_2014.dmg"
	resp, err := s.client.Get(dmgUrl)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	_, err = io.Copy(dmgFile, resp.Body)
	return err
}
