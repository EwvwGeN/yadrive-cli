package resource

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	v1 "github.com/EwvwGeN/yadrive-cli/pkg/yandex-api/v1"
)

func DownloadFileByPath(writer io.Writer, oauthToken, pathFrom, pathTo string) (error) {
	// should use os.PathSeparator
	fullPath := strings.Split(pathTo, "/")
	if fullPath[len(fullPath)-1] == "" {
		pathTo = filepath.Join(pathTo, filepath.Base(pathFrom))
	}
	resp, err := v1.GetFileDownloadLink(oauthToken, pathFrom)
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(pathTo), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(pathTo, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	err = v1.DownloadFileFromLink(oauthToken, resp.Href, file)
	if err != nil {
		file.Close()
		return err
	}
	file.Close()
	return nil
}