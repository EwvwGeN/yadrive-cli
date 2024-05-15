package templates

import (
	"bytes"
	"text/tabwriter"
	"text/template"

	"github.com/EwvwGeN/yadrive-cli/pkg/domain/models"
)

const diskInfoTmpl = `
Designation{{"\t"}}Bytes
Total space{{"\t"}}{{.TotalSpace}}
Used space{{"\t"}}{{.UsedSpace}}
Of which the trash occupies{{"\t"}}{{.TrashSize}}

Drive occupancy: {{ printf "%.2f" (percent .UsedSpace .TotalSpace) }} %

System folders{{"\t"}}Path
{{range $index, $elem := .SystemFolders}}{{$index}}{{"\t"}}{{$elem}}{{"\n"}}{{end}}
`

func GetDiskInfoTmpl(diskInfo models.DiskInfo) ([]byte, error) {

	t, err := template.New("disk_info").Funcs(functions).Parse(diskInfoTmpl)
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	writer := tabwriter.NewWriter(&buffer, 5, 0, 3, ' ', tabwriter.Debug)
	err = t.Execute(writer, diskInfo)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}