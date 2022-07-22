package clean

import (
	"io/ioutil"
	"os"

	"github.com/maniizu3110/go-codegen/clean/usecase"
	"github.com/maniizu3110/go-codegen/util"
	"gopkg.in/yaml.v2"
)

func Create(filePath string) {
	setting := util.Setting{}
	//pathはcobraのオプションから受け取る

	b, _ := os.ReadFile(filePath)
	yaml.Unmarshal(b, &setting)
	files, err := ioutil.ReadDir(setting.DomainPath)
	if err != nil {
		panic(err)
	}
	usecase.Gen(setting.UsecaseSetting, files)
}
