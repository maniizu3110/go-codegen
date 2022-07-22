package usecase

import (
	"bytes"
	_ "embed"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"

	"github.com/maniizu3110/go-codegen/util"
)

//go:embed template/common.txt
var commonTxt string

func createCommonUsecaseFile(setting util.UsecaseSetting) {
	params := map[string]any{
		"domain":            setting.DomainPackage,
		"errorType":         setting.ErrorType,
		"repositoryPackage": setting.RepositoryPackage,
	}
	var buf bytes.Buffer
	t := template.Must(template.New("meta-txt").Parse(commonTxt))
	t.Execute(&buf, params)

	// Output phase
	var out bytes.Buffer
	out.WriteString(
		`// Code generated; DO NOT EDIT.
package usecase
`)
	out.Write(buf.Bytes())
	body, err := format.Source(out.Bytes())
	if err != nil {
		// The user can compile the output to see the error.
		log.Fatalf("warning: internal error: invalid Go generated: %s", err)
	}
	setting.OutputDir = setting.OutputDir + "/" + commonUsecaseFileName
	if err := ioutil.WriteFile(setting.OutputDir, body, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}

}
