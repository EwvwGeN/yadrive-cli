package models

import (
	"bytes"
	"fmt"
)

type CustomProperties map[string]string

func (cp CustomProperties) String() string {
	if len(cp) == 0 {
		return ""
	}
	builder := bytes.Buffer{}
	builder.WriteByte(0x7B)
	for k, v := range cp {
		builder.WriteString(fmt.Sprintf("\"%s\":\"%s\",", k, v))
	}
	out := builder.Bytes()
	out[len(out)-1] = 0x7D
	return string(out)
}

func (cp CustomProperties) Bytes() []byte {
	if len(cp) == 0 {
		return []byte{}
	}
	builder := bytes.Buffer{}
	builder.WriteByte(0x7B)
	for k, v := range cp {
		builder.WriteString(fmt.Sprintf("\"%s\":\"%s\",", k, v))
	}
	out := builder.Bytes()
	out[len(out)-1] = 0x7D
	return out
}