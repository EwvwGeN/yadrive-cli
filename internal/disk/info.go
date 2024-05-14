package disk

import (
	"fmt"
	"io"

	v1 "github.com/EwvwGeN/yadrive-cli/pkg/yandex-api/v1"
)

func GetDiskInfo(writer io.Writer, oauthToken string) error {
	resp, err := v1.GetDiskInfo(oauthToken)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}