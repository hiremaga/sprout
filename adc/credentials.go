package adc

import (
	"fmt"

	"github.com/howeyc/gopass"
)

type Credentials struct {
	username string
	password string
}

func (c *Credentials) Ask() {
	fmt.Print("ADC username: ")
	fmt.Scanf("%s\n", &c.username)

	fmt.Print("ADC password: ")
	c.password = string(gopass.GetPasswdMasked())
}
