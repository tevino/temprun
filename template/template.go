package template

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"text/template"
)

type EnvTemplate struct {
	// TBD
	prefix  string
	env     map[string]string
	funcMap template.FuncMap
	seps    []string
}

var defaultSeps = []string{"\n", " ", ","}

func NewEnvTemplate(osEnv []string, prefix string, seps []string) *EnvTemplate {
	if seps == nil {
		seps = defaultSeps
	}
	tpl := &EnvTemplate{
		prefix: prefix,
		seps:   seps,
		env:    makeEnv(osEnv),
	}
	tpl.generateFuncMap()
	return tpl
}

func (t *EnvTemplate) RenderToWriter(reader io.Reader, writer io.Writer) error {
	tpl, err := t.createTemplate(reader)
	if err != nil {
		return err
	}
	return tpl.Execute(writer, t.env)
}

func (t *EnvTemplate) getEnv(key string) (string, bool) {
	if value, exists := t.env[key]; exists {
		return value, true
	}
	if value, exists := t.env[t.prefix+key]; exists {
		return value, true
	}
	return "", false
}

// Template Function Map

func (t *EnvTemplate) fGetv(key string) (string, error) {
	if value, exists := t.getEnv(key); exists {
		return value, nil
	}
	return "", fmt.Errorf("Key  not found '%s'", key)
}

func (t *EnvTemplate) fGetvs(key string) ([]string, error) {
	value, err := t.fGetv(key)
	if err != nil {
		return nil, err
	}
	sep, ok := t.getSep(value)
	if !ok {
		return []string{value}, nil
	}
	return strings.Split(value, sep), nil
}

// End Template Function Map

func (t *EnvTemplate) getSep(str string) (string, bool) {
	for _, sep := range t.seps {
		if strings.Contains(str, sep) {
			return sep, true
		}
	}
	return "", false
}

func (t *EnvTemplate) generateFuncMap() {
	//t.funcMap = make(template.FuncMap, 2)
	t.funcMap = template.FuncMap{
		"getv":  t.fGetv,
		"getvs": t.fGetvs,
	}
	//	t.funcMap["getv"] = t.fGetv
	//	t.funcMap["getvs"] = t.fGetvs
}
func makeEnv(osEnv []string) map[string]string {
	envMap := make(map[string]string)
	for _, env := range osEnv {
		kv := strings.SplitN(env, "=", 2)
		if len(kv) < 2 { // is this possible?
			continue
		}
		key := kv[0]
		value := kv[1]
		envMap[strings.ToLower(key)] = value
	}
	return envMap
}

func (t *EnvTemplate) createTemplate(reader io.Reader) (*template.Template, error) {
	input, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	inputTpl, err := template.New("input").Funcs(t.funcMap).Parse(string(input))
	if err != nil {
		return nil, err
	}
	return inputTpl, nil
}
