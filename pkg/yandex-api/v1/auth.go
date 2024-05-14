package v1

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/EwvwGeN/yadrive-cli/pkg/domain/models"
)

func GetAccessTokenForDevice(clientId string, options ...models.Option) (models.DeviceAccessCodeResponse, error) {
	values := url.Values{}
	values.Set("client_id", clientId)
	for i := 0; i < len(options); i++ {
		values.Set(options[i].Key, options[i].Value)
	}
	resp, err := http.PostForm("https://oauth.yandex.ru/device/code", values)
	if err != nil {
		return models.DeviceAccessCodeResponse{}, err
	}
	if resp.StatusCode == 200 {
		respFormated := models.DeviceAccessCodeResponse{}
		err = json.NewDecoder(resp.Body).Decode(&respFormated)
		if err != nil {
			return models.DeviceAccessCodeResponse{}, err
		}
		resp.Body.Close()
		return respFormated, nil
	}
	errResp := models.ErrorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return models.DeviceAccessCodeResponse{}, err
	}
	return models.DeviceAccessCodeResponse{}, fmt.Errorf(errResp.ErrDesc)
}

func GetOAuthTokenForDevice(deviceCode, clientId, clientSecret string) (models.OAuthTokenResponse, error) {
	values := url.Values{}
	values.Set("grant_type", "device_code")
	values.Set("code", deviceCode)
	req, err := http.NewRequest(http.MethodPost, "https://oauth.yandex.ru/", strings.NewReader(values.Encode()))
	if err != nil {
		return models.OAuthTokenResponse{}, err
	}
	clientDataUnion := []byte(clientId + ":" + clientSecret)
	encodedToken := base64.StdEncoding.EncodeToString(clientDataUnion)
	authToken := "Basic " + encodedToken
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", authToken)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.OAuthTokenResponse{}, err
	}
	if resp.StatusCode == 200 {
		respFormated := models.OAuthTokenResponse{}
		err = json.NewDecoder(resp.Body).Decode(&respFormated)
		if err != nil {
			return models.OAuthTokenResponse{}, err
		}
		resp.Body.Close()
		return respFormated, nil
	}
	errResp := models.ErrorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return models.OAuthTokenResponse{}, err
	}
	return models.OAuthTokenResponse{}, fmt.Errorf(errResp.ErrDesc)
}

func LoopGetOAuthTokenForDevice(deviceCode, clientId, clientSecret string, interval int, expires int64) (models.OAuthTokenResponse, error) {
	if int64(interval) > expires {
		return models.OAuthTokenResponse{}, fmt.Errorf("expire time is lower than interval")
	}
	var resp *http.Response
	values := url.Values{}
	values.Set("grant_type", "device_code")
	values.Set("code", deviceCode)

	clientDataUnion := []byte(clientId + ":" + clientSecret)
	encodedToken := base64.StdEncoding.EncodeToString(clientDataUnion)
	authToken := "Basic " + encodedToken

	client := http.Client{}
INNERLOOP:
	for {
		select {
		case <-time.After(time.Duration(interval) * time.Second):
			req, err := http.NewRequest(http.MethodPost, "https://oauth.yandex.ru/token", strings.NewReader(values.Encode()))
			if err != nil {
				return models.OAuthTokenResponse{}, err
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Authorization", authToken)
			resp, err = client.Do(req)
			if err != nil {
				return models.OAuthTokenResponse{}, err
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				respFormated := models.OAuthTokenResponse{}
				err = json.NewDecoder(resp.Body).Decode(&respFormated)
				if err != nil {
					return models.OAuthTokenResponse{}, err
				}
				return respFormated, nil
			}
		case <-time.After(time.Duration(expires) * time.Second):
			break INNERLOOP
		}
	}
	errResp := models.ErrorResponse{}
	err := json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return models.OAuthTokenResponse{}, err
	}
	return models.OAuthTokenResponse{}, fmt.Errorf(errResp.ErrDesc)
}