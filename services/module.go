package services

import (
	"log"
	"os"
	"path/filepath"

	"github.com/lucasnhso/sygnus-cli/helpers"
	"github.com/lucasnhso/sygnus-cli/utils"
)

func ModuleIsValid(module string) bool {
	modules := []string{"repository", "usecase", "controller", "resource"}

	for _, m := range modules {
		if m == module {
			return true
		}
	}

	return false
}

func GenerateModuleFile(module, name string) {
	outputFileName := getOutputFileName(module, name)

	// Create directories if they don't exist
	outputDir := filepath.Dir(outputFileName)
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	helpers.GenerateFileByTemplate(module, name, outputFile)

	log.Printf("Generated %s file: %s", module, outputFileName)
}

func getModuleFolder(module string) string {
	folders := map[string]string{
		"repository": "repositories",
		"usecase":    "useCases",
		"controller": "controllers",
	}
	return folders[module]
}

func getOutputFileName(module, name string) string {
	return "src/" + getModuleFolder(module) + "/" + utils.ToPascalCase(name) + utils.ToPascalCase(module) + ".ts"
}
