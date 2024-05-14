package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/EwvwGeN/yadrive-cli/pkg/domain/models"
)

func GetResourceInfoFromTrash(oauthToken, path string, options ...models.Option) (models.Resource, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/trash/resources")
	if err != nil {
		return models.Resource{}, err
	}
	values := url.Values{}
	values.Set("path", path)
	for i := 0; i < len(options); i++ {
		values.Set(options[i].Key, options[i].Value)
	}
	parsedUrl.RawQuery = values.Encode()
	req, err := http.NewRequest(http.MethodGet, parsedUrl.String(), nil)
	if err != nil {
		return models.Resource{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth "+oauthToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.Resource{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respFormated := models.Resource{}
		err = json.NewDecoder(resp.Body).Decode(&respFormated)
		if err != nil {
			return models.Resource{}, err
		}
		return respFormated, nil
	}
	errResp := models.ErrorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return models.Resource{}, err
	}
	return models.Resource{}, fmt.Errorf(errResp.ErrDesc)
}

func CleanUpTresh(oauthToken string, options ...models.Option) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/trash/resources")
	if err != nil {
		return models.Link{}, err
	}
	if len(options) != 0 {
		values := url.Values{}
		for i := 0; i < len(options); i++ {
			values.Set(options[i].Key, options[i].Value)
		}
		parsedUrl.RawQuery = values.Encode()
	}
	req, err := http.NewRequest(http.MethodDelete, parsedUrl.String(), nil)
	if err != nil {
		return models.Link{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth " + oauthToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.Link{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		return models.Link{}, nil
	}
	if resp.StatusCode == 202 {
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

func RestoreFileFromTrash(oauthToken, path string, options ...models.Option) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/trash/resources/restore")
	if err != nil {
		return models.Link{}, err
	}
	values := url.Values{}
	values.Set("path", path)
	for i := 0; i < len(options); i++ {
		values.Set(options[i].Key, options[i].Value)
	}
	parsedUrl.RawQuery = values.Encode()

	req, err := http.NewRequest(http.MethodPut, parsedUrl.String(), nil)
	if err != nil {
		return models.Link{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.Link{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 201 || resp.StatusCode == 202 {
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