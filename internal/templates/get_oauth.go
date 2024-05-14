package templates

import (
	"bytes"
	"fmt"
)

var accessToDeviceStr string = `
Vist the: %s
And enter the code: %s
`

var oauthStr string = `
received oauth token: %s
`

func GetAccessToDeviceTmpl(verificationLink, userCode string) ([]byte, error) {
	var buffer bytes.Buffer
	_, err := buffer.WriteString(fmt.Sprintf(accessToDeviceStr, verificationLink, userCode))
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func GetOauthTokenTmpl(ouathToken string) ([]byte, error) {
	var buffer bytes.Buffer
	_, err := buffer.WriteString(fmt.Sprintf(oauthStr, ouathToken))
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil 
}