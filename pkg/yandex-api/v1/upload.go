package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/EwvwGeN/yadrive-cli/pkg/domain/models"
)

func GetUploadLink(oauthToken, path string, options ...models.Option) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources/upload")
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

func UploadFileToDisk(link string, data io.Reader) error {
	req, err := http.NewRequest(http.MethodPut, link, data)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		return nil
	}
	errResp := models.ErrorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return err
	}
	return fmt.Errorf(errResp.ErrDesc)
}

func UploadFileFromLinkToDisk(oauthToken, link, path string, options ...models.Option) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources/upload")
	if err != nil {
		return models.Link{}, err
	}
	values := url.Values{}
	values.Set("url", link)
	values.Set("path", path)
	if len(options) != 0 {
		for i := 0; i < len(options); i++ {
			values.Set(options[i].Key, options[i].Value)
		}
	}
	parsedUrl.RawQuery = values.Encode()

	req, err := http.NewRequest(http.MethodPost, parsedUrl.String(), nil)
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
	if resp.StatusCode == 201 {
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