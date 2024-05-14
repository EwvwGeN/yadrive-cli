package token

import (
	"io"

	"github.com/EwvwGeN/yadrive-cli/internal/constant"
	"github.com/EwvwGeN/yadrive-cli/internal/templates"
	v1 "github.com/EwvwGeN/yadrive-cli/pkg/yandex-api/v1"
	"github.com/spf13/viper"
)

func GetAccessTokenForDevice(writer io.Writer, clientId, clientSecret string) error {
	resp, err := v1.GetAccessTokenForDevice(clientId)
	if err != nil {
		return err
	}
	tmlpAccess, err := templates.GetAccessToDeviceTmpl(resp.VerificationUrl, resp.UserCode)
	if err != nil {
		return err
	}
	_, err = writer.Write(tmlpAccess)
	if err != nil {
		return err
	}
	oauthResp, err := v1.LoopGetOAuthTokenForDevice(resp.DeviceCode, clientId, clientSecret, resp.Interval, resp.ExpIn)
	if err != nil {
		return err
	}
	viper.Set(constant.OauthFlag, oauthResp.AccessToken)
	tmlpOauth, err := templates.GetOauthTokenTmpl(oauthResp.AccessToken)
	if err != nil {
		return err
	}
	_, err = writer.Write(tmlpOauth)
	if err != nil {
		return err
	}
	return nil
}