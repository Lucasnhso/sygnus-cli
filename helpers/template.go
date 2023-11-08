package helpers

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/lucasnhso/sygnus-cli/templates"
	"github.com/lucasnhso/sygnus-cli/utils"
)

func getTemplate(name string) string {
	templates := map[string]string{
		"repository": templates.Repository(),
		"usecase":    templates.UseCase(),
		"controller": templates.Controller(),
	}
	return templates[name]
}

func GetTemplateFileName(name string) string {
	path := filepath.Join("templates", name+".tmpl")
	res, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("Error getting template file path: %v", err)
	}

	return res
}

func GenerateFileByTemplate(module string, name string, outputFile *os.File) {
	templateText := getTemplate(module)

	tmpl, err := template.New("file").Parse(templateText)
	if err != nil {
		log.Fatalf("Error loading template: %v", err)
	}

	data := struct {
		Name      string
		PCaseName string
	}{
		Name:      name,
		PCaseName: utils.ToPascalCase(name),
	}

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
