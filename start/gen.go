package start

import (
	"bytes"
	_ "embed"
	"html/template"
	"io/ioutil"
)

//go:embed codegen.txt
var templateTxt string

func CreateDefatultYaml() {
	var buf bytes.Buffer
	t := template.Must(template.New("meta-txt").Parse(templateTxt))
	t.Execute(&buf, nil)
	ioutil.WriteFile("codegen.yaml", buf.Bytes(), 0644)

}
