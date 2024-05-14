package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/EwvwGeN/yadrive-cli/pkg/domain/models"
)

func GetFileDownloadLink(oauthToken, path string, options ...models.Option) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources/download")
	if err != nil {
		return models.Link{}, err
	}
	values := url.Values{}
	values.Set("path", path)
	if len(options) != 0 {
		for i := 0; i < len(options); i++ {
			values.Set(options[i].Key, options[i].Value)
		}
	}
	parsedUrl.RawQuery = values.Encode()

	req, err := http.NewRequest(http.MethodGet, parsedUrl.String(), nil)
	if err != nil {
		return models.Link{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth "+oauthToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.Link{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respFormated := models.Link{}
		err = json.NewDecoder(resp.Body).Decode(&respFormated)
		if err != nil {
			return models.Link{}, err
		}
		return respFormated, nil
	}
	errResp := models.ErrorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return models.Link{}, err
	}
	return models.Link{}, fmt.Errorf(errResp.ErrDesc)
}

func DownloadFileFromLink(oauthToken, link string, reciver io.Writer) error {
	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth "+oauthToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		_, err = io.Copy(reciver, resp.Body)
		if err != nil {
			return err
		}
		return nil
	}
	errResp := models.ErrorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return err
	}
	return fmt.Errorf(errResp.ErrDesc)
}