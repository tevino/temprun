#!/usr/bin/env bash

# Environment variables
export t_go_template=https://golang.org/pkg/text/template/
export t_separators="newline space comma"
export t_test="you could
use ' ' and ','
without escaping!"
export x_test="this should not be printed"

# Build temprun
make clean
make

# Template
./temprun -prefix t_ <<EOF
It's just go template: {{getv "go_template"}}

separators could be(highest priority first):
{{- range getvs "separators"}}
    {{. -}}
{{- end}}

For example:
{{- range getvs "test"}}
    {{. -}}
{{end -}}
EOF
