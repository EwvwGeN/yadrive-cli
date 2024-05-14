package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/EwvwGeN/yadrive-cli/pkg/domain/models"
)

func GetResourceInfoFromDisk(oauthToken, path string, options ...models.Option) (models.Resource, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources")
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

func GetFileList(oauthToken string, options ...models.Option) (models.FileResourceList, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources/files")
	if err != nil {
		return models.FileResourceList{}, err
	}
	if len(options) != 0 {
		values := url.Values{}
		for i := 0; i < len(options); i++ {
			values.Set(options[i].Key, options[i].Value)
		}
		parsedUrl.RawQuery = values.Encode()
	}
	req, err := http.NewRequest(http.MethodGet, parsedUrl.String(), nil)
	if err != nil {
		return models.FileResourceList{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth " + oauthToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.FileResourceList{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respFormated := models.FileResourceList{}
		err = json.NewDecoder(resp.Body).Decode(&respFormated)
		if err != nil {
			return models.FileResourceList{}, err
		}
		return respFormated, nil
	}
	errResp := models.ErrorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return models.FileResourceList{}, err
	}
	return models.FileResourceList{}, fmt.Errorf(errResp.ErrDesc)
}

func GetLastUploadedFiles(oauthToken string, options ...models.Option) (models.LastUploadedResourceList, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources/last-uploaded")
	if err != nil {
		return models.LastUploadedResourceList{}, err
	}
	if len(options) != 0 {
		values := url.Values{}
		for i := 0; i < len(options); i++ {
			values.Set(options[i].Key, options[i].Value)
		}
		parsedUrl.RawQuery = values.Encode()
	}
	req, err := http.NewRequest(http.MethodGet, parsedUrl.String(), nil)
	if err != nil {
		return models.LastUploadedResourceList{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth " + oauthToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.LastUploadedResourceList{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respFormated := models.LastUploadedResourceList{}
		err = json.NewDecoder(resp.Body).Decode(&respFormated)
		if err != nil {
			return models.LastUploadedResourceList{}, err
		}
		return respFormated, nil
	}
	errResp := models.ErrorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return models.LastUploadedResourceList{}, err
	}
	return models.LastUploadedResourceList{}, fmt.Errorf(errResp.ErrDesc)
}

func UpdateResourceMeta(oauthToken, path string, newMeta models.CustomProperties, options ...models.Option) (models.Resource, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources/")
	if err != nil {
		return models.Resource{}, err
	}
	values := url.Values{}
	values.Set("path", path)
	if len(options) != 0 {
		for i := 0; i < len(options); i++ {
			values.Set(options[i].Key, options[i].Value)
		}
	}
	parsedUrl.RawQuery = values.Encode()

	var jsonBody bytes.Buffer
	jsonBody.WriteString("\"custom_properties\":")
	jsonBody.Write(newMeta.Bytes())

	req, err := http.NewRequest(http.MethodPatch, parsedUrl.String(), &jsonBody)
	if err != nil {
		return models.Resource{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth " + oauthToken)
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

func CopyResource(oauthToken, fromPath, toPath string, options ...models.Option) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources/copy")
	if err != nil {
		return models.Link{}, err
	}
	values := url.Values{}
	values.Set("from", fromPath)
	values.Set("path", toPath)
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
	req.Header.Set("Authorization", "OAuth " + oauthToken)
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

func MoveResource(oauthToken, fromPath, toPath string, options ...models.Option) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources/move")
	if err != nil {
		return models.Link{}, err
	}
	values := url.Values{}
	values.Set("from", fromPath)
	values.Set("path", toPath)
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
	req.Header.Set("Authorization", "OAuth " + oauthToken)
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

func DeleteResource(oauthToken, path string, options ...models.Option) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources")
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

func CreateFolder(oauthToken, path string, options ...models.Option) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources")
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

	req, err := http.NewRequest(http.MethodPut, parsedUrl.String(), nil)
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

func PublishResource(oauthToken, path string) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources/publish")
	if err != nil {
		return models.Link{}, err
	}
	values := url.Values{}
	values.Set("path", path)
	parsedUrl.RawQuery = values.Encode()

	req, err := http.NewRequest(http.MethodPut, parsedUrl.String(), nil)
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

func UnpublishResource(oauthToken, path string) (models.Link, error) {
	parsedUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/resources/unpublish")
	if err != nil {
		return models.Link{}, err
	}
	values := url.Values{}
	values.Set("path", path)
	parsedUrl.RawQuery = values.Encode()

	req, err := http.NewRequest(http.MethodPut, parsedUrl.String(), nil)
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