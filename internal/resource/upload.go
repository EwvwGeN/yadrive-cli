package resource

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/EwvwGeN/yadrive-cli/pkg/domain/models"
	v1 "github.com/EwvwGeN/yadrive-cli/pkg/yandex-api/v1"
)

func UploadFileByLink(writer io.Writer, oauthToken, link, path string, redirect *string) error {
	option := []models.Option{}
	if redirect != nil {
		option = append(option, models.Option{
			Key: "disable_redirects",
			Value: *redirect,
		})
	}
	// print to template?
	_, err := v1.UploadFileFromLinkToDisk(oauthToken, link, path, option...)
	if err != nil {
		return err
	}

	return nil
}

func UploadFileByPath(writer io.Writer, oauthToken, pathFrom, pathTo string, overwrite *string) error {
	option := []models.Option{}
	if overwrite != nil {
		option = append(option, models.Option{
			Key: "overwrite",
			Value: *overwrite,
		})
	}
	fullPath := strings.Split(pathTo, "/")
	if fullPath[len(fullPath)-1] == "" {
		fullPath[len(fullPath)-1] = filepath.Base(pathFrom)
		pathTo = strings.Join(fullPath, "/")
	}
	resp, err := v1.GetUploadLink(oauthToken, pathTo, option...)
	if err != nil {
		return err
	}
	file, err := os.Open(pathFrom)
	if err != nil {
		return err
	}
	err = v1.UploadFileToDisk(resp.Href, file)
	if err != nil {
		return err
	}
	return nil
}