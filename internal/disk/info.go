package disk

import (
	"io"

	"github.com/EwvwGeN/yadrive-cli/internal/templates"
	v1 "github.com/EwvwGeN/yadrive-cli/pkg/yandex-api/v1"
)

func GetDiskInfo(writer io.Writer, oauthToken string) error {
	resp, err := v1.GetDiskInfo(oauthToken)
	if err != nil {
		return err
	}
	tmlpInfo, err := templates.GetDiskInfoTmpl(resp)
	if err != nil {
		return err
	}
	_, err = writer.Write(tmlpInfo)
	if err != nil {
		return err
	}
	return nil
}