package usecase

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/maniizu3110/go-codegen/util"
)

//go:embed template/template.txt
var templateTxt string

const commonUsecaseFileName = "common.go"

//modeasl はfilesに含まれるので削除する予定
func Gen(setting util.UsecaseSetting, files []fs.FileInfo) {
	if _, err := os.Stat(setting.OutputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(setting.OutputDir, 0777); err != nil {
			fmt.Println(err)
		}
	}
	if !IsCommonFileExist(setting.OutputDir, commonUsecaseFileName) {
		createCommonUsecaseFile(setting)
	}
	for i := range files {
		Create(setting, files[i].Name())
	}
}

func Create(setting util.UsecaseSetting, fileName string) {
	model := strings.Split(fileName, ".")[0]
	params := map[string]any{
		"Model": strcase.ToCamel(model),
		"model": strcase.ToLowerCamel(model),
	}
	// Generate phase: template.Execute
	var buf bytes.Buffer
	t := template.Must(template.New("meta-txt").Parse(templateTxt))
	t.Execute(&buf, params)

	// Output phase
	var out bytes.Buffer
	out.WriteString(`
package usecase
`)
	out.Write(buf.Bytes())
	body, err := format.Source(out.Bytes())
	if err != nil {
		// The user can compile the output to see the error.
		log.Fatalf("warning: internal error: invalid Go generated: %s", err)
	}

	setting.OutputDir = setting.OutputDir + "/" + model + ".go"
	if err := ioutil.WriteFile(setting.OutputDir, body, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
