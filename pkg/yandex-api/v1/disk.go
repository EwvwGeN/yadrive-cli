package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/EwvwGeN/yadrive-cli/pkg/domain/models"
)

func GetDiskInfo(oauthToken string) (models.DiskInfo, error) {
	req, err := http.NewRequest(http.MethodGet, "https://cloud-api.yandex.net/v1/disk/", nil)
	if err != nil {
		return models.DiskInfo{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth "+oauthToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.DiskInfo{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respFormated := models.DiskInfo{}
		err = json.NewDecoder(resp.Body).Decode(&respFormated)
		if err != nil {
			return models.DiskInfo{}, err
		}
		return respFormated, nil
	}
	errResp := models.ErrorResponse{}
	err = json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return models.DiskInfo{}, err
	}
	return models.DiskInfo{}, fmt.Errorf(errResp.ErrDesc)
}