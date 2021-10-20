package htmlParser

import (
	"os"
	"text/template"

	"github.com/google/uuid"
)

type htmlStruct struct {
	rootPath string
}

func New(rootPath string) HTMLParserInterface {
	return &htmlStruct{rootPath: rootPath}
}

func (a *htmlStruct) Create(templateName string, data interface{}) (string, error) {

	templeteGenerator, err := template.ParseFiles(templateName)
	if err != nil {
		return "", err
	}

	fileName := a.rootPath + "/" + uuid.New().String() + ".html"

	fileWriter, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	if err := templeteGenerator.Execute(fileWriter, data); err != nil {
		return "", err
	}

	return fileName, nil
}
