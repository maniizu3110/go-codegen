package clean

import (
	"io/ioutil"
	"os"

	"github.com/maniizu3110/go-codegen/clean/repository"
	"github.com/maniizu3110/go-codegen/clean/usecase"
	"github.com/maniizu3110/go-codegen/util"
	"gopkg.in/yaml.v2"
)

func CreateAll(filePath string) {
	setting := util.Setting{}

	b, _ := os.ReadFile(filePath)
	yaml.Unmarshal(b, &setting)
	files, err := ioutil.ReadDir(setting.DomainPath)
	if err != nil {
		panic(err)
	}
	usecase.Gen(setting.UsecaseSetting, files)
	// TODO: domainの処理追加
	// TODO: repositoryの処理追加
	// TODO: delivaryの処理追加
}

func CreateUsecase(filePath string) {
	setting := util.Setting{}

	b, _ := os.ReadFile(filePath)
	yaml.Unmarshal(b, &setting)
	files, err := ioutil.ReadDir(setting.DomainPath)
	if err != nil {
		panic(err)
	}
	usecase.Gen(setting.UsecaseSetting, files)
}
func CreateRepository(filePath string) {
	setting := util.Setting{}

	b, _ := os.ReadFile(filePath)
	yaml.Unmarshal(b, &setting)
	files, err := ioutil.ReadDir(setting.DomainPath)
	if err != nil {
		panic(err)
	}
	repository.Gen(setting.RepositorySetting, files)
}
