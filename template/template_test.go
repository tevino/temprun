package template

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var OKEnvs = []string{
	"name=temprun",
	"seps=newline space comma",
}

const OKTemplate = `Testing: {{getv "name"}}
Testing range:
{{- range $i, $v := getvs "seps"}}
    {{$i}}: {{$v -}}
{{- end}}
`

const OKOutput = `Testing: temprun
Testing range:
    0: newline
    1: space
    2: comma
`

func TestRenderToWriter(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	tpl := NewEnvTemplate(OKEnvs, "", nil)
	err := tpl.RenderToWriter(strings.NewReader(OKTemplate), buf)
	assert.NoError(t, err)
	assert.Equal(t, OKOutput, buf.String())
}
