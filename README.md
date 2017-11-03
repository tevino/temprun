# temprun
Lightweight file generator using environment variables, used for rendering configuration in docker containers.

## Example
```shell
# Environment variables
export t_go_template=https://golang.org/pkg/text/template/
export t_separators="newline space comma"
export t_test="you could
use ' ' and ','
without escaping!"

# Template
./temprun <<EOF
It's just go template: {{getv "t_go_template"}}

separators could be(highest priority first):
{{- range getvs "t_separators"}}
    {{. -}}
{{- end}}

For example:
{{- range getvs "t_test"}}
    {{. -}}
{{end -}}
EOF
```

**Output:**
```
It's just go template: https://golang.org/pkg/text/template/

separators could be(highest priority first):
    newline
    space
    comma

For example:
    you could
    use ' ' and ','
    without escaping!
```
