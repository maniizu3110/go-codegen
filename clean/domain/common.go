package dommain

import (
	"bytes"
	_ "embed"
	"html/template"
	"io/ioutil"
)

//go:embed template/common.txt
var templateTxt string

func Gen() {
	var buf bytes.Buffer
	t := template.Must(template.New("meta-txt").Parse(templateTxt))
	t.Execute(&buf, nil)
	ioutil.WriteFile("codegen.yaml", buf.Bytes(), 0644)

}
