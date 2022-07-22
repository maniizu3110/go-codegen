package main

import (
	"io/ioutil"
	"os"

	"github.com/maniizu3110/go-codegen/clean/usecase"
	"github.com/maniizu3110/go-codegen/util"
	"gopkg.in/yaml.v2"
)

func main() {
	setting := util.Setting{}

	b, _ := os.ReadFile("codegen.yaml")
	yaml.Unmarshal(b, &setting)
	files, err := ioutil.ReadDir(setting.DomainPath)
	if err != nil {
		panic(err)
	}
	usecase.Gen(setting.UsecaseSetting, files)
}
