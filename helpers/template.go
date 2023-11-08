package helpers

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/lucasnhso/sygnus-cli/utils"
)

func GetTemplateFileName(name string) string {
	return filepath.Join("templates", name+".tmpl")
}

func GenerateFileByTemplate(name string, outputFile *os.File) {
	templateFileName := GetTemplateFileName(name)

	data := struct {
		Name      string
		PCaseName string
	}{
		Name:      name,
		PCaseName: utils.ToPascalCase(name),
	}

	tmpl, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
	}

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
