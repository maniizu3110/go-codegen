package util

type Setting struct {
	Architecture      string `yaml:"architecture" default:"clean"`
	DomainPath        string `yaml:"domainPath" default:"./model"`
	DeliverySetting   `yaml:"delivery"`
	UsecaseSetting    `yaml:"usecase"`
	RepositorySetting `yaml:"repository"`
}

type DeliverySetting struct {
	OutputDir     string `yaml:"outputDir" default:"./delivery"`
	DomainPackage string `yaml:"domainPackage" default:"domain"`
	ErrorType     string `yaml:"errorType" default:"error"`
}

type UsecaseSetting struct {
	OutputDir         string `yaml:"outputDir" default:"./usecase"`
	DomainPackage     string `yaml:"domainPackage" default:"domain"`
	ErrorType         string `yaml:"errorType" default:"error"`
	RepositoryPackage string `yaml:"repositoryPackage" default:"repository"`
}

type RepositorySetting struct {
	OutputDir         string `yaml:"outputDir" default:"./usecase"`
	DomainPackage     string `yaml:"domainPackage" default:"domain"`
	ErrorType         string `yaml:"errorType" default:"error"`
	RepositoryPackage string `yaml:"repositoryPackage" default:"repository"`
}
